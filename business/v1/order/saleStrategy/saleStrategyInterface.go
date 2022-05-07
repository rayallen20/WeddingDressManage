package saleStrategy

type StrategyInterface interface {
	calcPrice(strategy *SaleStrategy)
	calcCharterMoney()
	calcCashPledge()
}
