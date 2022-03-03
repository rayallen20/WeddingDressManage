package dress

import (
	"WeddingDressManage/business/v1/customer"
	"WeddingDressManage/lib/helper/sliceHelper"
	"WeddingDressManage/lib/helper/urlHelper"
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/model"
	requestParam "WeddingDressManage/param/request/v1/dress"
	"WeddingDressManage/param/resps/v1/pagination"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strings"
	"time"
)

// Dress 礼服类 即具体的每一件礼服
type Dress struct {
	Id              int
	CategoryId      int
	Category        *Category
	SerialNumber    int
	Size            string
	RentCounter     int
	LaundryCounter  int
	MaintainCounter int
	CoverImg        string
	SecondaryImg    []string
	Status          string
}

func (d *Dress) Add(param *requestParam.AddParam) ([]*Dress, error) {
	// step1. 查询品类是否存在
	categoryORM := &model.DressCategory{
		Id: param.Category.Id,
	}

	err := categoryORM.FindById()

	// 数据库错误
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, &sysError.DbError{RealError: err}
	}

	// 品类信息不存在错误
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, &sysError.CategoryNotExistError{Id: param.Category.Id}
	}

	// step2. 创建礼服ORM集合
	dressORM := &model.Dress{CategoryId: param.Category.Id}
	maxSerialNumber, err := dressORM.FindMaxSerialNumberByCategoryId()
	if err != nil {
		return nil, &sysError.DbError{RealError: err}
	}

	dressORMs := d.createDressORMForAdd(param, maxSerialNumber)
	categoryORM.Quantity += len(dressORMs)
	categoryORM.RentableQuantity += len(dressORMs)

	// step3. 使用事务创建礼服信息并更新礼服品类信息
	err = dressORM.AddDressesAndUpdateCategory(categoryORM, dressORMs)
	if err != nil {
		return nil, &sysError.DbError{RealError: err}
	}

	dresses := make([]*Dress, 0, param.Dress.Number)
	for _, completeDressORM := range dressORMs {
		dress := &Dress{}
		dress.fill(completeDressORM)
		dresses = append(dresses, dress)
	}

	return dresses, nil
}

func (d *Dress) createDressORMForAdd(param *requestParam.AddParam, maxSerialNumber int) []*model.Dress {
	dressORMs := make([]*model.Dress, 0, param.Dress.Number)
	for i := 1; i <= param.Dress.Number; i++ {
		dressORM := &model.Dress{
			CategoryId:      param.Category.Id,
			SerialNumber:    maxSerialNumber + i,
			Size:            param.Dress.Size,
			RentCounter:     0,
			LaundryCounter:  0,
			MaintainCounter: 0,
			CoverImg:        param.Dress.CoverImg,
			SecondaryImg:    sliceHelper.ImpactSliceToStr(param.Dress.SecondaryImg, "|"),
			Status:          model.DressStatus["onSale"],
		}

		dressORMs = append(dressORMs, dressORM)
	}

	return dressORMs
}

func (d *Dress) fill(orm *model.Dress) {
	d.Id = orm.Id
	d.CategoryId = orm.CategoryId
	if orm.Category != nil {
		d.Category = &Category{
			Id:               orm.Category.Id,
			SerialNumber:     orm.Category.SerialNumber,
			Quantity:         orm.Category.Quantity,
			RentableQuantity: orm.Category.RentableQuantity,
			CharterMoney:     orm.Category.CharterMoney,
			AvgCharterMoney:  orm.Category.AvgCharterMoney,
			CashPledge:       orm.Category.CashPledge,
			RentCounter:      orm.Category.RentCounter,
			LaundryCounter:   orm.Category.LaundryCounter,
			MaintainCounter:  orm.Category.MaintainCounter,
			CoverImg:         urlHelper.GenFullImgWebSite(orm.Category.CoverImg),
			SecondaryImg:     urlHelper.GenFullImgWebSites(strings.Split(orm.Category.SecondaryImg, "|")),
			Status:           orm.Category.Status,
		}

		if orm.Category.Kind != nil {
			d.Category.Kind = &Kind{
				Id:     orm.Category.Kind.Id,
				Name:   orm.Category.Kind.Name,
				Code:   orm.Category.Kind.Code,
				Status: orm.Category.Kind.Status,
			}
		}
	}
	d.SerialNumber = orm.SerialNumber
	d.Size = orm.Size
	d.RentCounter = orm.RentCounter
	d.LaundryCounter = orm.LaundryCounter
	d.MaintainCounter = orm.MaintainCounter
	d.CoverImg = urlHelper.GenFullImgWebSite(orm.CoverImg)
	d.SecondaryImg = urlHelper.GenFullImgWebSites(strings.Split(orm.SecondaryImg, "|"))
	d.Status = orm.Status
}

func (d *Dress) ShowUsable(param *requestParam.ShowUsableParam) (category *Category, usableDresses []*Dress, totalPage int, count int64, err error) {
	// step1. 查品类信息是否存在
	categoryOrm := &model.DressCategory{Id: param.Category.Id}
	err = categoryOrm.FindById()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil, 0, 0, &sysError.DbError{RealError: err}
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil, 0, 0, &sysError.CategoryNotExistError{Id: param.Category.Id}
	}

	// step2. 查询总页数
	dressOrm := &model.Dress{CategoryId: param.Category.Id}
	count, err = dressOrm.CountUsableByCategoryId()
	if err != nil {
		return nil, nil, 0, 0, &sysError.DbError{RealError: err}
	}
	totalPage = pagination.CalcTotalPage(count, param.Pagination.ItemPerPage)

	// step3. 根据品类信息分页查询礼服
	// Tips: 查询总页数时使用的orm由于已经被用作查询过 所以导致其内部有Id字段等信息 故此处需重新创建一个orm
	dressOrm = &model.Dress{CategoryId: param.Category.Id}
	usableDressOrms, err := dressOrm.FindUsableByCategoryId(param.Pagination.CurrentPage, param.Pagination.ItemPerPage)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil, 0, 0, &sysError.DbError{RealError: err}
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil, 0, 0, &sysError.DressNotExistError{}
	}

	category = &Category{}
	category.fill(categoryOrm)
	usableDresses = make([]*Dress, 0, len(usableDressOrms))
	for _, usableOrm := range usableDressOrms {
		usableDress := &Dress{}
		usableDress.fill(usableOrm)
		usableDresses = append(usableDresses, usableDress)
	}

	return category, usableDresses, totalPage, count, nil
}

func (d *Dress) ApplyDiscard(param *requestParam.ApplyDiscardParam) error {
	// step1. 查询礼服是否存在
	orm := &model.Dress{Id: param.Dress.Id}
	err := orm.FindById()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return &sysError.DbError{RealError: err}
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &sysError.DressNotExistError{Id: param.Dress.Id}
	}

	// step2. 确认礼服状态 当礼服状态为已赠与 或 已销库 时 不可提出销库申请
	d.fill(orm)
	if d.Status == model.DressStatus["gift"] {
		return &sysError.DressHasGiftedError{}
	}

	if d.Status == model.DressStatus["discard"] {
		return &sysError.DressHasDiscardedError{}
	}

	// step3. 写入销库申请
	discardAskBiz := &DiscardAsk{
		Dress: d,
		Note:  param.DiscardAsk.Note,
	}

	return discardAskBiz.Apply()
}

func (d *Dress) ApplyGift(param *requestParam.ApplyGiftParam) error {
	// step1. 查询礼服是否存在
	dressOrm := &model.Dress{Id: param.Dress.Id}
	err := dressOrm.FindById()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return &sysError.DbError{RealError: err}
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &sysError.DressNotExistError{Id: param.Dress.Id}
	}

	// step2. 确认礼服状态 当礼服状态为已赠与 或 已销库 时 不可提出销库申请
	d.fill(dressOrm)
	if d.Status == model.DressStatus["gift"] {
		return &sysError.DressHasGiftedError{}
	}

	if d.Status == model.DressStatus["discard"] {
		return &sysError.DressHasDiscardedError{}
	}

	// step3. 查询客户是否存在
	customerBiz := &customer.Customer{
		Name:   param.Customer.Name,
		Mobile: param.Customer.Mobile,
	}

	err = customerBiz.FindNormalByNameAndMobile()

	if err != nil {
		return err
	}

	// step4. 写入赠与申请
	giftAskBiz := &GiftAsk{
		Dress:    d,
		Customer: customerBiz,
		Note:     param.GiftAsk.Note,
	}

	return giftAskBiz.Apply()
}

func (d *Dress) Laundry(param *requestParam.LaundryParam) error {
	// step1. 查询礼服是否存在
	dressOrm := &model.Dress{Id: param.Dress.Id}
	err := dressOrm.FindById()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return &sysError.DbError{RealError: err}
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &sysError.DressNotExistError{Id: param.Dress.Id}
	}

	// step2. 确认礼服状态 当礼服状态不为 在售/预租赁/预上架 时 不允许送洗
	d.fill(dressOrm)
	if !d.canBeLaundry() {
		return &sysError.LaundryStatusError{DressNowStatus: d.Status}
	}

	// step3. 修改礼服ORM 设置礼服状态为送洗中 送洗次数+1
	dressOrm.Status = model.DressStatus["laundry"]
	dressOrm.LaundryCounter += 1

	// step4. 对品类ORM 送洗次数+1
	categoryOrm := &model.DressCategory{
		Id:             d.Category.Id,
		LaundryCounter: d.Category.LaundryCounter + 1,
	}

	// step5. 创建送洗记录
	laundryRecordBiz := &LaundryRecord{
		Dress:             d,
		DirtyPositionImg:  param.LaundryDetail.DirtyPositionImg,
		Note:              param.LaundryDetail.Note,
		StartLaundryDate:  time.Now(),
		DueEndLaundryDate: time.Now().Add(LaundryPlanDurationDays * 24 * time.Hour),
		Status:            model.LaundryStatus["underway"],
	}

	// step6. 使用事务
	// 1. 修改礼服状态 礼服送洗次数+1
	// 2. 礼服所属品类送洗次数+1
	// 3. 创建送洗记录
	laundryRecordOrm := laundryRecordBiz.CreateORMForLaundry()
	err = dressOrm.UpdateDressStatusAndCreateLaundryRecord(categoryOrm, laundryRecordOrm)
	if err != nil {
		return &sysError.DbError{RealError: err}
	}

	d.Status = model.DressStatus["laundry"]
	return nil
}

func (d *Dress) canBeLaundry() bool {
	canBeLaundryStatuses := []string{
		model.DressStatus["onSale"],
		model.DressStatus["preRent"],
		model.DressStatus["preOnSale"],
	}

	for _, canBeLaundryStatus := range canBeLaundryStatuses {
		if d.Status == canBeLaundryStatus {
			return true
		}
	}

	return false
}

func (d *Dress) Maintain(param *requestParam.MaintainParam) error {
	// step1. 查询礼服是否存在
	dressOrm := &model.Dress{Id: param.Dress.Id}
	err := dressOrm.FindById()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return &sysError.DbError{RealError: err}
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &sysError.DressNotExistError{Id: param.Dress.Id}
	}

	// step2. 确认礼服状态 当礼服状态不为 在售/预租赁/预上架 时 不允许维护
	d.fill(dressOrm)
	if !d.canBeMaintain() {
		return &sysError.MaintainStatusError{DressNowStatus: d.Status}
	}

	// step3. 修改礼服ORM 设置礼服状态为维护中 维护次数+1
	dressOrm.Status = model.DressStatus["maintain"]
	dressOrm.MaintainCounter += 1

	// step4. 对品类ORM 维护次数+1
	categoryOrm := &model.DressCategory{
		Id:              d.Category.Id,
		MaintainCounter: d.Category.MaintainCounter + 1,
	}

	// step5. 创建维护记录
	dailyMaintainBiz := &DailyMaintainRecord{
		Source:              model.MaintainSource["daily"],
		Dress:               d,
		MaintainPositionImg: param.MaintainDetail.MaintainPositionImg,
		Note:                param.MaintainDetail.Note,
		StartMaintainDate:   time.Now(),
		PlanEndMaintainDate: time.Now().Add(MaintainPlanDurationDays * 24 * time.Hour),
		Status:              model.MaintainStatus["underway"],
	}
	maintainOrm := dailyMaintainBiz.CreateORMForDailyMaintain()

	// step6. 使用事务
	// 1. 修改礼服状态 礼服维护次数+1
	// 2. 礼服所属品类维护次数+1
	// 3. 创建维护记录
	err = dressOrm.UpdateDressStatusAndCreateMaintainRecord(categoryOrm, maintainOrm)
	if err != nil {
		return &sysError.DbError{RealError: err}
	}

	d.Status = model.DressStatus["maintain"]
	return nil
}

func (d *Dress) canBeMaintain() bool {
	canBeMaintainStatuses := []string{
		model.DressStatus["onSale"],
		model.DressStatus["preRent"],
		model.DressStatus["preOnSale"],
	}

	for _, canBeMaintainStatus := range canBeMaintainStatuses {
		if d.Status == canBeMaintainStatus {
			return true
		}
	}

	return false
}

func (d *Dress) ShowOne(param *requestParam.ShowOneParam) error {
	orm := &model.Dress{Id: param.Dress.Id}
	err := orm.FindById()
	fmt.Printf("%#v\n", orm.Category.Kind)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return &sysError.DbError{RealError: err}
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &sysError.DressNotExistError{Id: param.Dress.Id}
	}

	d.fill(orm)
	return nil
}

func (d *Dress) Update(param *requestParam.UpdateParam) error {
	orm := &model.Dress{Id: param.Dress.Id}
	err := orm.FindById()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return &sysError.DbError{RealError: err}
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &sysError.DressNotExistError{Id: param.Dress.Id}
	}

	orm.Size = param.Dress.Size
	orm.CoverImg = param.Dress.CoverImg
	orm.SecondaryImg = sliceHelper.ImpactSliceToStr(param.Dress.SecondaryImg, "|")
	err = orm.Updates()
	if err != nil {
		return &sysError.DbError{RealError: err}
	}

	d.fill(orm)
	return nil
}

func (d *Dress) ShowUnusable(param *requestParam.ShowUnusableParam) (category *Category, unusableDresses []*Dress, totalPage int, count int64, err error) {
	// step1. 查品类信息是否存在
	categoryOrm := &model.DressCategory{Id: param.Category.Id}
	err = categoryOrm.FindById()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil, 0, 0, &sysError.DbError{RealError: err}
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil, 0, 0, &sysError.CategoryNotExistError{Id: param.Category.Id}
	}

	// step2. 查询总页数
	dressOrm := &model.Dress{CategoryId: param.Category.Id}
	count, err = dressOrm.CountUnusableByCategoryId()
	if err != nil {
		return nil, nil, 0, 0, &sysError.DbError{RealError: err}
	}
	totalPage = pagination.CalcTotalPage(count, param.Pagination.ItemPerPage)

	// step3. 根据品类信息分页查询礼服
	// Tips: 查询总页数时使用的orm由于已经被用作查询过 所以导致其内部有Id字段等信息 故此处需重新创建一个orm
	dressOrm = &model.Dress{CategoryId: param.Category.Id}
	unusableDressOrms, err := dressOrm.FindUnusableByCategoryId(param.Pagination.CurrentPage, param.Pagination.ItemPerPage)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil, 0, 0, &sysError.DbError{RealError: err}
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil, 0, 0, &sysError.DressNotExistError{}
	}

	category = &Category{}
	category.fill(categoryOrm)
	unusableDresses = make([]*Dress, 0, len(unusableDressOrms))
	for _, unusableDressOrm := range unusableDressOrms {
		unusableDress := &Dress{}
		unusableDress.fill(unusableDressOrm)
		unusableDresses = append(unusableDresses, unusableDress)
	}

	return category, unusableDresses, totalPage, count, nil
}
