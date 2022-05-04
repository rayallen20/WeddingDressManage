package model

import (
	"time"
)

// SaleStrategy 销售策略
// originalPrice:原价
// discount:打折
// customPrice:自定义租金与押金
var SaleStrategy map[string]string = map[string]string{
	"originalPrice": "originalPrice",
	"discount":      "discount",
	"customPrice":   "customPrice",
}

// OrderStatus 订单状态
// unsettled:未付租金
// notYetDelivery:租金已付,未出件
// deliving:出件中
// deliveryFinish:出件结束
// retrieving:还件中
// finish:订单结束
var OrderStatus map[string]string = map[string]string{
	"unsettled":      "unsettled",
	"notYetDelivery": "notYetDelivery",
	"deliving":       "deliving",
	"deliveryFinish": "deliveryFinish",
	"retrieving":     "retrieving",
	"finish":         "finish",
}

type Order struct {
	Id                     int
	CustomerId             int
	Customer               *Customer `gorm:"foreignKey:CustomerId"`
	SerialNumber           string
	WeddingDate            time.Time
	OriginalCharterMoney   int
	OriginalCashPledge     int
	SaleStrategy           string
	Discount               float64
	DuePayCharterMoney     int
	DuePayCashPledge       int
	DueRefundCashPledge    int
	ActualPayCharterMoney  int
	ActualPayCashPledge    int
	ActualRefundCashPledge int
	TotalMaintainFee       int
	Status                 string
	CreatedTime            time.Time `gorm:"autoCreateTime"`
	UpdatedTime            time.Time `gorm:"autoUpdateTime"`
}
