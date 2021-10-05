package model

import (
	"WeddingDressManage/lib/db"
	"time"
)

var DressStatus = map[string]string{
	// 可租借
	"rentable": "rentable",
	// 租借中
	"rentOut": "rentOut",
	// 送洗中
	"laundry": "laundry",
	// 已赠与
	"gift": "gift",
}

type DressDetail struct {
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

func (d *DressDetail) Create(categoryId int, size string) (err error) {
	d.CategoryId = categoryId
	d.Size = size
	d.Status = DressStatus["rentable"]
	res := db.Db.Save(d)
	return res.Error
}
