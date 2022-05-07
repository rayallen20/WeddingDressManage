package saleStrategy

type originalPriceStrategy struct {
	originalCharterMoney int
	originalCashPledge   int
	duePayCharterMoney   int
	duePayCashPledge     int
}

func (o *originalPriceStrategy) calcPrice(strategy *SaleStrategy) {
	o.originalCharterMoney = strategy.OriginalCharterMoney
	o.originalCashPledge = strategy.OriginalCashPledge

	o.calcCharterMoney()
	strategy.DuePayCharterMoney = o.duePayCharterMoney

	o.calcCashPledge()
	strategy.DuePayCashPledge = o.duePayCashPledge
}

// calcCharterMoney 原价策略下计算租金
// 规则:租金不变
func (o *originalPriceStrategy) calcCharterMoney() {
	o.duePayCharterMoney = o.originalCharterMoney
}

// calcCashPledge 原价策略下计算押金
// 规则:押金不变
func (o *originalPriceStrategy) calcCashPledge() {
	o.duePayCashPledge = o.originalCashPledge
}
