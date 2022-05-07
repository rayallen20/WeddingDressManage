package model

import "time"

// BillType 账单类型
// collectCharterMoney:收取租金
// collectCashPledge:收取押金
// restituteCashPledge:退还押金
// restituteCharterMoney:退还租金
// maintain:维护费用支出
var BillType map[string]string = map[string]string{
	"collectCharterMoney":   "collectCharterMoney",
	"collectCashPledge":     "collectCashPledge",
	"restituteCashPledge":   "restituteCashPledge",
	"restituteCharterMoney": "restituteCharterMoney",
	"maintain":              "maintain",
}

// BillStatus 账单状态
// notStarted:账单未开始
// underway:账单进行中
// finish:账单结束
// cancel:账单取消 本状态用于修改订单时 将原订单的相关账单取消 仅可在账单的应付金额为0时置为cancel状态
var BillStatus map[string]string = map[string]string{
	"notStarted": "notStarted",
	"underway":   "underway",
	"finish":     "finish",
	"cancel":     "cancel",
}

type Bill struct {
	Id               int
	Type             string
	OrderId          int
	Order            *Order `gorm:"foreignKey:OrderId"`
	MaintainRecordId int
	MaintainRecord   *MaintainRecord `gorm:"foreignKey:MaintainRecordId"`
	AmountPayable    int
	AmountPaid       int
	Status           string
	CreatedTime      time.Time `gorm:"autoCreateTime"`
	UpdatedTime      time.Time `gorm:"autoUpdateTime"`
}
