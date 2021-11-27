package dress

import (
	"WeddingDressManage/lib/helper/sliceHelper"
	"WeddingDressManage/lib/helper/urlHelper"
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/model"
	"WeddingDressManage/param/v1/request/dress"
	"errors"
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
		d.Category = &Category{
			Id:               orm.Category.Id,
			Kind:             &Kind{
				Id:     orm.Category.Kind.Id,
				Name:   orm.Category.Kind.Name,
				Code:   orm.Category.Kind.Code,
				Status: orm.Category.Kind.Status,
			},
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