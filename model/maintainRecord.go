package model

import "time"

// MaintainSource 维护记录来源
// item 订单内礼服归还时需要维护(此来源的维护信息需包含订单相关信息 会涉及到钱)
// daily 商家日常维护(此来源的维护信息无订单相关信息 不涉及到钱)
var MaintainSource map[string]string = map[string]string{
	"item":  "item",
	"daily": "daily",
}

// MaintainStatus 维护状态
// underway 维护中
// finish 维护结束
var MaintainStatus map[string]string = map[string]string{
	"underway": "underway",
	"finish":   "finish",
}

type MaintainRecord struct {
	Id     int
	Source string
	// TODO:此处差order的orm 但目前暂时没做到订单模块 故未写
	OrderId int
	// TODO:此处差order_item的orm 但目前暂时没做到订单模块 故未写
	OrderItemId            int
	DressId                int
	Dress                  *Dress `gorm:"foreignKey:DressId"`
	DuePayMaintainMoney    int
	ActualPayMaintainMoney int
	MaintainPositionImg    string `gorm:"text"`
	Note                   string `gorm:"text"`
	StartMaintainDate      time.Time
	PlanEndMaintainDate    time.Time
	EndMaintainDate        time.Time
	// TODO gorm使用问题:若字段为time.Time(其他类型尚未遇到) 则Save()/Create()时 该字段为空值时 默认该字段值不为空 需在更新时 使用Select()方法指定更新字段
	Status      string
	CreatedTime time.Time `gorm:"autoCreateTime"`
	UpdatedTime time.Time `gorm:"autoUpdateTime"`
}
