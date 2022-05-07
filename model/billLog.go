package model

import "time"

type BillLog struct {
	Id                int
	BillId            int
	Bill              *Bill `gorm:"foreignKey:BillId"`
	OrderItemId       int
	OrderItem         *OrderItem `gorm:"foreignKey:OrderItemId"`
	DressId           int
	Dress             *Dress `gorm:"foreignKey:DressId"`
	TransactionAmount int
	CreatedTime       time.Time `gorm:"autoCreateTime"`
	UpdatedTime       time.Time `gorm:"autoUpdateTime"`
}
