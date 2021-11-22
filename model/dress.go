package model

import "time"

// DressStatus 礼服状态
// onSale:在售
// preRent:预租
// rentOut:租借中
// pending:回收后待处理
// preOnSale:预上架
// laundry:送洗中
// maintain:维护中
var DressStatus = map[string]string{
	"onSale":"onSale",
	"preRent":"preRent",
	"rentOut":"rentOut",
	"pending":"pending",
	"preOnSale":"preOnSale",
	"laundry":"laundry",
	"maintain":"maintain",
}

// Dress dress表的ORM
type Dress struct {
	Id int
	CategoryId int
	SerialNumber int
	RentCounter int
	LaundryCounter int
	MaintainCounter int
	CoverImg string
	SecondaryImg string `gorm:"text"`
	Status string
	CreatedTime time.Time `gorm:"autoCreateTime"`
	UpdatedTime time.Time `gorm:"autoUpdatedTime"`
}
