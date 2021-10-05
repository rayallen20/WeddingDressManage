package dress

import (
	"WeddingDressManage/lib/structHelper"
	"WeddingDressManage/model"
	"time"
)

type Detail struct {
	// 礼服ID
	Id int

	// 礼服品类ID
	CategoryId int

	// 尺码
	Size string

	// 出租次数
	RentNumber int

	// 送洗次数
	LaundryNumber int

	// 状态
	Status string

	CreatedTime time.Time `gorm:"autoCreateTime"`
	UpdatedTime time.Time `gorm:"autoUpdateTime"`
}

func (d *Detail) Add(categoryId int, size string) (err error) {
	var dressDetailModel *model.DressDetail = &model.DressDetail{}
	err = dressDetailModel.Create(categoryId, size)
	structHelper.StructAssign(d, dressDetailModel)
	return err
}
