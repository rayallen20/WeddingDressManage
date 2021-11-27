package model

import "time"

// SecondaryEntityType 二级实体类型
// dressCategory:礼服品类信息
// dress:礼服信息
// order:订单信息
// orderItem:订单条目信息
// dailyMaintainRecord:日常维护信息
// itemMaintainRecord:订单内礼服归还时维护信息
// laundryRecord:送洗信息
var SecondaryEntityType map[string]string = map[string]string{
	"dressCategory":"dressCategory",
	"dress":"dress",
	"order":"order",
	"orderItem":"orderItem",
	"dailyMaintainRecord":"dailyMaintainRecord",
	"itemMaintainRecord":"itemMaintainRecord",
	"laundryRecord":"laundryRecord",
}

type OperationSecondaryEntity struct {
	Id int
	OperationLogId int
	OperationLog *OperationLog `gorm:"foreignKey:OperationLogId"`
	SecondaryEntityType string
	SecondaryEntityId int
	CreatedTime time.Time `gorm:"autoCreateTime"`
	UpdatedTime time.Time `gorm:"autoUpdateTime"`
}
