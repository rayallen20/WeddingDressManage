package order

import (
	"WeddingDressManage/business/v1/dress"
	"WeddingDressManage/business/v1/order/saleStrategy"
)

type BillLog struct {
	Id                int
	Bill              *Bill
	OrderItem         *Item
	Dress             *dress.Dress
	TransactionAmount int
}

// calcCharterTransactionAmount 计算租金的交易记录
func (b *BillLog) calcCharterMoney(order *Order, bill *Bill) []*BillLog {
	if order.SaleStrategy == saleStrategy.OriginalPrice {
		// 原租金的交易记录
		return b.calcOriginalPriceTransactionAmount(order, bill)
	}

	if order.SaleStrategy == saleStrategy.Discount {
		// 计算打折交易记录
		return b.calcDiscountTransactionAmount(order, bill)
	}

	if order.SaleStrategy == saleStrategy.CustomPrice {
		// 自定义的交易记录
		return b.calcCustomPriceTransactionAmount(bill)
	}

	return nil
}

// calcOriginalPriceTransactionAmount 计算优惠策略为原价时 应付租金的交易记录
func (b *BillLog) calcOriginalPriceTransactionAmount(order *Order, bill *Bill) (billLogs []*BillLog) {
	billLogs = make([]*BillLog, 0, len(order.Items))
	for _, item := range order.Items {
		billLog := &BillLog{
			Bill:              bill,
			OrderItem:         item,
			Dress:             item.Dress,
			TransactionAmount: item.Dress.Category.CharterMoney,
		}
		billLogs = append(billLogs, billLog)
	}
	return billLogs
}

// calcDiscountTransactionAmount 计算优惠策略为打折时 应付租金的交易记录
// 由于打折策略下 订单中最终成交的租金价格截取了角 分 和 元的个位数
// 因此最终成交的租金金额小于每件礼服打折后的金额之和
// 这部分偏差值计算在每笔订单的最后一件礼服的租金上
func (b *BillLog) calcDiscountTransactionAmount(order *Order, bill *Bill) (billLogs []*BillLog) {
	billLogs = make([]*BillLog, 0, len(order.Items))
	// 若订单只有1件礼服 则该件礼服的租金交易记录金额即为订单最终成交的租金价格
	if len(billLogs) == 1 {
		billLog := &BillLog{
			Bill:              bill,
			OrderItem:         order.Items[0],
			Dress:             order.Items[0].Dress,
			TransactionAmount: order.ActualPayCharterMoney,
		}
		billLogs = append(billLogs, billLog)
		return billLogs
	}

	// step1. 按每件礼服折扣后的价格创建交易记录
	for _, item := range order.Items {
		billLog := &BillLog{
			Bill:      bill,
			OrderItem: item,
			Dress:     item.Dress,
		}
		billLog.TransactionAmount = int(float64(item.Dress.Category.CharterMoney) * order.Discount)
		billLogs = append(billLogs, billLog)
	}

	// step2. 用订单最终成交租金 - (前 n - 1)件礼服的折扣后租金 即为最后一件礼服的成交金额
	lastDressTransactionAmount := order.ActualPayCharterMoney
	for i := 0; i <= len(billLogs)-2; i++ {
		lastDressTransactionAmount -= billLogs[i].TransactionAmount
	}
	billLogs[len(billLogs)-1].TransactionAmount = lastDressTransactionAmount
	return billLogs
}

// calcCustomPriceTransactionAmount 计算优惠策略为自定义租金时 应付租金的交易记录
// 此策略下的交易记录不对应具体的item 因为自定义租金的粒度为订单而非订单内的每一件礼服 故只记录交易金额
func (b *BillLog) calcCustomPriceTransactionAmount(bill *Bill) (billLogs []*BillLog) {
	billLogs = make([]*BillLog, 0, 1)
	billLog := &BillLog{
		Bill:              bill,
		TransactionAmount: bill.AmountPaid,
	}
	billLogs = append(billLogs, billLog)
	return billLogs
}

// calcCashPledge 计算押金的交易记录
// 若优惠策略为自定义押金 则只生成1条押金交易记录 对应整个订单的应付押金
// 否则按每件礼服的押金计算交易金额
func (b *BillLog) calcCashPledge(order *Order, bill *Bill) []*BillLog {
	if order.SaleStrategy == saleStrategy.CustomPrice {
		billLogs := make([]*BillLog, 0, 1)
		billLog := &BillLog{
			Bill:              bill,
			TransactionAmount: bill.AmountPaid,
		}
		billLogs = append(billLogs, billLog)
		return billLogs
	} else {
		billLogs := make([]*BillLog, 0, len(order.Items))
		for _, item := range order.Items {
			billLog := &BillLog{
				Bill:              bill,
				OrderItem:         item,
				Dress:             item.Dress,
				TransactionAmount: item.Dress.Category.CashPledge,
			}
			billLogs = append(billLogs, billLog)
		}
		return billLogs
	}
}
