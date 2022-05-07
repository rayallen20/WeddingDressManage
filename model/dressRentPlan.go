package model

import (
	"time"
)

// DressRentPlanStatus 礼服租赁计划状态 当客户修改订单中的礼服时 原订单中的礼服租赁状态会失效
// valid:计划有效
// invalid:计划无效
var DressRentPlanStatus map[string]string = map[string]string{
	"valid":   "valid",
	"invalid": "invalid",
}

type DressRentPlan struct {
	Id                int
	OrderId           int
	OrderSerialNumber string
	Order             *Order `gorm:"foreignKey:OrderId"`
	OrderItemId       int
	Item              *OrderItem `gorm:"foreignKey:OrderItemId"`
	DressId           int
	Dress             *Dress `gorm:"foreignKey:DressId"`
	PreRentStartDate  time.Time
	WeddingDate       time.Time
	PreOnSaleEndDate  time.Time
	Status            string
}
