package order

import (
	"WeddingDressManage/business/v1/dress"
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/model"
	"errors"
	"gorm.io/gorm"
)

type Bill struct {
	Id             int
	Type           string
	Order          *Order
	MaintainRecord *dress.MaintainRecord
	AmountPayable  int
	AmountPaid     int
	BillLogs       []*BillLog
	Status         string
}

// createCollectCharterMoney 创建实付租金账单与交易记录
func (b *Bill) createCollectCharterMoney(order *Order) {
	b.Type = model.BillType["collectCharterMoney"]
	b.Order = order
	b.AmountPayable = order.DuePayCharterMoney
	b.AmountPaid = order.ActualPayCharterMoney
	b.Status = model.BillStatus["finish"]

	billLogBiz := &BillLog{}
	b.BillLogs = billLogBiz.calcCharterMoney(order, b)
}

// createCollectCashPledge 创建实付押金账单与交易记录
func (b *Bill) createCollectCashPledge(order *Order, isSettled bool) {
	b.Type = model.BillType["collectCashPledge"]
	b.Order = order
	b.AmountPayable = order.DuePayCashPledge
	if isSettled {
		b.AmountPaid = order.ActualPayCashPledge
		b.Status = model.BillStatus["finish"]
		billLogBiz := &BillLog{}
		b.BillLogs = billLogBiz.calcCashPledge(order, b)
	} else {
		b.Status = model.BillStatus["notStarted"]
	}
}

// createRestituteCashPledge 创建退还押金账单
func (b *Bill) createRestituteCashPledge(order *Order) {
	b.Type = model.BillType["restituteCashPledge"]
	b.Order = order
	b.AmountPayable = -order.DueRefundCashPledge
	b.Status = model.BillStatus["notStarted"]
}

func (b *Bill) genORMForCreateOrder() (billORM *model.Bill, billLogORMs []*model.BillLog) {
	billORM = &model.Bill{
		Type:          b.Type,
		AmountPayable: b.AmountPayable,
		AmountPaid:    b.AmountPaid,
		Status:        b.Status,
	}

	billLogORMs = make([]*model.BillLog, 0, len(b.BillLogs))
	for _, billLog := range b.BillLogs {
		billLogORM := &model.BillLog{
			TransactionAmount: billLog.TransactionAmount,
		}

		// 此处由于优惠策略为自定义租金与押金时 billLog无DressId 故需判断
		if billLog.Dress != nil {
			billLogORM.DressId = billLog.Dress.Id
		}

		billLogORMs = append(billLogORMs, billLogORM)
	}
	return billORM, billLogORMs
}

func (b *Bill) FindCashPledge(order *Order) error {
	orm := &model.Bill{
		Type:    model.BillType["collectCashPledge"],
		OrderId: order.Id,
	}
	err := orm.FindByTypeAndOrderId()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return &sysError.DbError{RealError: err}
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &sysError.BillNotFoundError{OrderId: order.Id}
	}
	b.fill(orm)
	return nil
}

func (b *Bill) fill(orm *model.Bill) {
	b.Id = orm.Id
	b.Type = orm.Type
	if orm.Order != nil {
		orderBiz := &Order{}
		orderBiz.fill(orm.Order)
		b.Order = orderBiz
	}

	if orm.BillLogs != nil {
		billLogs := make([]*BillLog, 0, len(orm.BillLogs))
		for _, billLogORM := range orm.BillLogs {
			billLog := &BillLog{}
			billLog.fill(billLogORM)
			billLogs = append(billLogs, billLog)
		}
		b.BillLogs = billLogs
	}

	b.Status = orm.Status
}
