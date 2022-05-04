package model

import (
	"time"
)

// OrderItemStatus 订单条目状态
// valid:有效
// invalid:无效
var OrderItemStatus map[string]string = map[string]string{
	"valid":   "valid",
	"invalid": "invalid",
}

type OrderItem struct {
	Id          int
	OrderId     int
	Order       *Order `gorm:"foreignKey:OrderId"`
	DressId     int
	Dress       *Dress `gorm:"foreignKey:DressId"`
	MaintainFee int
	Status      string
	CreatedTime time.Time `gorm:"autoCreateTime"`
	UpdatedTime time.Time `gorm:"autoUpdateTime"`
}
