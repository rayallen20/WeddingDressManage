package dress

import (
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/model"
	"WeddingDressManage/param/v1/dress/category/request"
	"errors"
	"gorm.io/gorm"
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
		Id: param.Category.KindId,
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
	return nil
}