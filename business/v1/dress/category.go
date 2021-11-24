package dress

import (
	"WeddingDressManage/lib/helper/sliceHelper"
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/model"
	"WeddingDressManage/param/v1/dress/category/request"
	"errors"
	"gorm.io/gorm"
	"strings"
)

// Category 礼服品类 相当于超市中商品的二级分类 例如:"食品-速冻食品" "食品-零食" "日用品-洗衣液"等类别
type Category struct {
	Id int
	Kind *Kind
	SerialNumber string
	Quantity int
	RentableQuantity int
	CharterMoney int
	AvgCharterMoney int
	CashPledge int
	RentCounter int
	LaundryCounter int
	MaintainCounter int
	CoverImg string
	SecondaryImg []string
	Status string
}

func (c *Category) Add(param *request.AddParam) error {
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
func(c *Category) createCategoryORMForAdd(categoryORM *model.DressCategory, param *request.AddParam) {
	categoryORM.KindId = param.Kind.Id
	categoryORM.Quantity = param.Dress.Number
	categoryORM.RentableQuantity = param.Dress.Number
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
func(c *Category) createDressORMForAdd(param *request.AddParam) []*model.Dress {
	dressORMs := make([]*model.Dress, 0, param.Dress.Number)
	for i := 1; i <= param.Dress.Number; i++ {
		dressORM := &model.Dress{
			SerialNumber:    i,
			Size: param.Dress.Size,
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
func(c *Category) fill(orm *model.DressCategory)  {
	c.Id = orm.Id
	c.Quantity = orm.Quantity
	c.RentableQuantity = orm.RentableQuantity
	c.CharterMoney = orm.CharterMoney
	c.AvgCharterMoney = orm.AvgCharterMoney
	c.CashPledge = orm.CashPledge
	c.RentCounter = orm.RentCounter
	c.LaundryCounter = orm.LaundryCounter
	c.MaintainCounter = orm.MaintainCounter
	c.CoverImg = orm.CoverImg
	c.SecondaryImg = strings.Split(orm.SecondaryImg, "|")
	c.Status = orm.Status
}