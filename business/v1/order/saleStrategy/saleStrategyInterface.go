package saleStrategy

type StrategyInterface interface {
	// calcPrice 计算原价(租金和押金)
	calcPrice(strategy *SaleStrategy)
	// calcCharterMoney 计算折扣策略后的租金
	calcCharterMoney()
	// calcCashPledge 计算扣策略后的押金
	calcCashPledge()
}
