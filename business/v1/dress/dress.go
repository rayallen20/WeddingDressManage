package dress

import (
	"WeddingDressManage/lib/helper/sliceHelper"
	"WeddingDressManage/lib/helper/urlHelper"
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/model"
	"WeddingDressManage/param/request/v1/dress"
	"WeddingDressManage/param/resps/v1/pagination"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

// Dress 礼服类 即具体的每一件礼服
type Dress struct {
	Id int
	CategoryId int
	Category *Category
	SerialNumber int
	Size string
	RentCounter int
	LaundryCounter int
	MaintainCounter int
	CoverImg string
	SecondaryImg []string
	Status string
}

func (d *Dress) Add(param *dress.AddParam) ([]*Dress, error) {
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

func (d *Dress) createDressORMForAdd(param *dress.AddParam, maxSerialNumber int) []*model.Dress {
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

func (d *Dress) fill(orm *model.Dress)  {
	d.Id = orm.Id
	d.CategoryId = orm.CategoryId
	if orm.Category != nil {
		d.Category = &Category {
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

func (d *Dress) ShowUsable(param *dress.ShowUsableParam) (category *Category, dresses []*Dress, totalPage int, err error) {
	// step1. 查品类信息是否存在
	categoryOrm := &model.DressCategory{Id: param.Category.Id}
	err = categoryOrm.FindById()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil, 0, &sysError.DbError{RealError: err}
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil, 0, &sysError.CategoryNotExistError{Id: param.Category.Id}
	}

	// step2. 查询总页数
	dressOrm := &model.Dress{CategoryId: param.Category.Id}
	count, err := dressOrm.CountUsableByCategoryId()
	if err != nil {
		return nil, nil, 0, &sysError.DbError{RealError: err}
	}
	totalPage = pagination.CalcTotalPage(count, param.Pagination.ItemPerPage)

	// step3. 根据品类信息分页查询礼服
	// Tips: 查询总页数时使用的orm由于已经被用作查询过 所以导致其内部有Id字段等信息 故此处需重新创建一个orm
	dressOrm = &model.Dress{CategoryId: param.Category.Id}
	usableDressOrms, err := dressOrm.FindUsableByCategoryId(param.Pagination.CurrentPage, param.Pagination.ItemPerPage)
	fmt.Printf("%d\n", len(usableDressOrms))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil, 0, &sysError.DbError{RealError: err}
	}

	category = &Category{}
	category.fill(categoryOrm)
	dresses = make([]*Dress, 0, len(usableDressOrms))
	for _, usableOrm := range usableDressOrms {
		dress := &Dress{}
		dress.fill(usableOrm)
		dresses = append(dresses, dress)
	}

	return category, dresses, totalPage, nil
}