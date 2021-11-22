package model

import (
	"WeddingDressManage/lib/db"
	"time"
)

// CategoryStatus 品类状态
// onSale:在售
// soldOut:脱销
// haltSales:下架
var CategoryStatus = map[string]string {
	"onSale":"onSale",
	"soldOut":"soldOut",
	"haltSales":"haltSales",
}

// DressCategory dress_category表的ORM
type DressCategory struct {
	Id int
	KindId int
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
	SecondaryImg string `gorm:"text"`
	Status string
	CreatedTime time.Time `gorm:"autoCreateTime"`
	UpdatedTime time.Time `gorm:"autoUpdateTime"`
}

// FindBySerialNumber 根据礼服品类编码 查找1条品类信息 默认为非脱销状态
func (c *DressCategory) FindBySerialNumber() error {
	res := db.Db.Not("status = ?", CategoryStatus["haltSales"]).Where(c).First(c)
	return res.Error
}
