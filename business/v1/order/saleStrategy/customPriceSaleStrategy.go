package saleStrategy

type customPriceSaleStrategy struct {
	originalCharterMoney int
	originalCashPledge   int
	customCharterMoney   int
	customCashPledge     int
	duePayCharterMoney   int
	duePayCashPledge     int
}

func (c *customPriceSaleStrategy) calcPrice(strategy *SaleStrategy) {
	c.originalCharterMoney = strategy.OriginalCharterMoney
	c.originalCashPledge = strategy.OriginalCashPledge
	c.customCharterMoney = strategy.CustomCharterMoney
	c.customCashPledge = strategy.CustomCashPledge

	c.calcCharterMoney()
	strategy.DuePayCharterMoney = c.duePayCharterMoney

	c.calcCashPledge()
	strategy.DuePayCashPledge = c.duePayCashPledge
}

// calcCharterMoney 自定义租金与押金策略下计算租金
// 规则:自定义租金为最终应付租金
func (c *customPriceSaleStrategy) calcCharterMoney() {
	c.duePayCharterMoney = c.customCharterMoney
}

// calcCashPledge 自定义租金与押金策略下计算押金
// 规则:自定义押金为最终应付押金
func (c *customPriceSaleStrategy) calcCashPledge() {
	c.duePayCashPledge = c.customCashPledge
}
