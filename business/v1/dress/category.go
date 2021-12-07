package dress

import (
	"WeddingDressManage/lib/helper/sliceHelper"
	"WeddingDressManage/lib/helper/urlHelper"
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/model"
	categoryRequest "WeddingDressManage/param/request/v1/category"
	"errors"
	"gorm.io/gorm"
	"math"
	"strings"
)

// Category 礼服品类 相当于超市中商品的二级分类 例如:"食品-速冻食品" "食品-零食" "日用品-洗衣液"等类别
type Category struct {
	Id               int
	Kind             *Kind
	SerialNumber     string
	Quantity         int
	RentableQuantity int
	CharterMoney     int
	AvgCharterMoney  int
	CashPledge       int
	RentCounter      int
	LaundryCounter   int
	MaintainCounter  int
	CoverImg         string
	SecondaryImg     []string
	Status           string
}

// Add 创建新品类并在该品类下添加礼服
func (c *Category) Add(param *categoryRequest.AddParam) error {
	// step1. 校验kind是否存在 若不存在则报错
	kind := &Kind{
		Id: param.Kind.Id,
	}
	err := kind.FindById()
	if err != nil {
		return err
	}

	// step2. 校验category是否存在 若存在则报错
	c.Kind = kind
	c.SerialNumber = kind.Code + "-" + param.Category.SequenceNumber
	categoryModel := &model.DressCategory{SerialNumber: c.SerialNumber}
	err = categoryModel.FindBySerialNumber()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return &sysError.DbError{RealError: err}
	}

	if categoryModel.Id != 0 {
		return &sysError.CategoryHasExistError{SerialNumber: c.SerialNumber}
	}

	// step3. 创建category的ORM和dress的ORM
	c.createCategoryORMForAdd(categoryModel, param)
	dressORMs := c.createDressORMForAdd(param)

	// step4. 使用事务创建新品类信息和新礼服信息
	err = categoryModel.AddCategoryAndDresses(dressORMs)
	if err != nil {
		return err
	}

	c.fill(categoryModel)

	return nil
}

// createCategoryORMForAdd 为添加新品类礼服创建品类信息ORM
func (c *Category) createCategoryORMForAdd(categoryORM *model.DressCategory, param *categoryRequest.AddParam) {
	categoryORM.KindId = param.Kind.Id
	categoryORM.Quantity = param.Dress.Number
	categoryORM.RentableQuantity = param.Dress.Number
	categoryORM.CharterMoney = param.Category.CharterMoney
	categoryORM.AvgCharterMoney = 0
	categoryORM.CashPledge = param.Category.CashPledge
	categoryORM.RentCounter = 0
	categoryORM.LaundryCounter = 0
	categoryORM.MaintainCounter = 0
	categoryORM.CoverImg = param.Category.CoverImg
	categoryORM.SecondaryImg = sliceHelper.ImpactSliceToStr(param.Category.SecondaryImg, "|")
	categoryORM.Status = model.CategoryStatus["onSale"]
}

// createDressORMForAdd 为添加新品类礼服创建礼服信息ORM集合
func (c *Category) createDressORMForAdd(param *categoryRequest.AddParam) []*model.Dress {
	dressORMs := make([]*model.Dress, 0, param.Dress.Number)
	for i := 1; i <= param.Dress.Number; i++ {
		dressORM := &model.Dress{
			SerialNumber:    i,
			Size:            param.Dress.Size,
			RentCounter:     0,
			LaundryCounter:  0,
			MaintainCounter: 0,
			CoverImg:        param.Category.CoverImg,
			SecondaryImg:    sliceHelper.ImpactSliceToStr(param.Category.SecondaryImg, "|"),
			Status:          model.DressStatus["onSale"],
		}
		dressORMs = append(dressORMs, dressORM)
	}
	return dressORMs
}

// fill 根据ORM信息填充品类对象
func (c *Category) fill(orm *model.DressCategory) {
	c.Id = orm.Id
	if orm.Kind != nil {
		c.Kind = &Kind{
			Id:     orm.Kind.Id,
			Name:   orm.Kind.Name,
			Code:   orm.Kind.Code,
			Status: orm.Kind.Status,
		}
	}
	c.SerialNumber = orm.SerialNumber
	c.Quantity = orm.Quantity
	c.RentableQuantity = orm.RentableQuantity
	c.CharterMoney = orm.CharterMoney
	c.AvgCharterMoney = orm.AvgCharterMoney
	c.CashPledge = orm.CashPledge
	c.RentCounter = orm.RentCounter
	c.LaundryCounter = orm.LaundryCounter
	c.MaintainCounter = orm.MaintainCounter
	c.CoverImg = urlHelper.GenFullImgWebSite(orm.CoverImg)
	if orm.SecondaryImg != "" {
		c.SecondaryImg = urlHelper.GenFullImgWebSites(strings.Split(orm.SecondaryImg, "|"))
	} else {
		c.SecondaryImg = make([]string, 0)
	}
	c.Status = orm.Status
}

// Show 礼服品类展示
func (c *Category) Show(param *categoryRequest.ShowParam) (categories []*Category, totalPage int, err error) {
	model := &model.DressCategory{}
	orms, err := model.FindNormal(param.Pagination.CurrentPage, param.Pagination.ItemPerPage)
	// TODO:此处要对没查到做handle?
	if err != nil {
		return nil, totalPage, &sysError.DbError{RealError: err}
	}

	categories = make([]*Category, 0, len(orms))

	for _, orm := range orms {
		category := &Category{}
		category.fill(orm)
		categories = append(categories, category)
	}

	// 计算总页数
	count, err := model.CountNormal()
	if err != nil {
		return nil, totalPage, &sysError.DbError{RealError: err}
	}

	totalPage = int(math.Ceil(float64(count) / float64(param.Pagination.ItemPerPage)))
	return categories, totalPage, nil
}

// ShowOne 根据id展示1条品类信息
func (c *Category) ShowOne(param *categoryRequest.ShowOneParam) error {
	orm := &model.DressCategory{Id: param.Category.Id}
	err := orm.FindById()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		dbErr := &sysError.DbError{RealError: err}
		return dbErr
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		notExistErr := &sysError.CategoryNotExistError{Id: param.Category.Id}
		return notExistErr
	}

	c.fill(orm)
	return nil
}

func (c *Category) Update(param *categoryRequest.UpdateParam) error {
	orm := &model.DressCategory{Id: param.Category.Id}
	err := orm.FindById()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		dbErr := &sysError.DbError{RealError: err}
		return dbErr
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		notExistErr := &sysError.CategoryNotExistError{Id: param.Category.Id}
		return notExistErr
	}

	orm.CharterMoney = param.Category.CharterMoney
	orm.CashPledge = param.Category.CashPledge
	orm.CoverImg = param.Category.CoverImg
	orm.SecondaryImg = sliceHelper.ImpactSliceToStr(param.Category.SecondaryImg, "|")
	err = orm.Updates()

	if err != nil {
		dbErr := &sysError.DbError{RealError: err}
		return dbErr
	}

	c.fill(orm)
	return nil
}
