package saleStrategy

import (
	"WeddingDressManage/business/v1/dress"
	"WeddingDressManage/lib/sysError"
)

// OriginalPrice 原价
const OriginalPrice = "originalPrice"

// Discount 打折
const Discount = "discount"

// CustomPrice 自定义租金与押金
const CustomPrice = "customPrice"

type SaleStrategy struct {
	Type                 string
	Discount             float64
	OriginalCharterMoney int
	OriginalCashPledge   int
	CustomCharterMoney   int
	CustomCashPledge     int
	DuePayCharterMoney   int
	DuePayCashPledge     int
	Dresses              []*dress.Dress
}

func (s *SaleStrategy) CalcPrice() error {
	strategy := s.genStrategy()
	if strategy == nil {
		return &sysError.StrategyNotExistError{Type: s.Type}
	}
	strategy.calcPrice(s)
	return nil
}

func (s *SaleStrategy) genStrategy() StrategyInterface {
	if s.Type == OriginalPrice {
		strategy := &originalPriceStrategy{
			originalCharterMoney: s.OriginalCharterMoney,
			originalCashPledge:   s.OriginalCashPledge,
		}
		return strategy
	}

	if s.Type == Discount {
		strategy := &discountStrategy{
			discount:             s.Discount,
			originalCharterMoney: s.OriginalCharterMoney,
			originalCashPledge:   s.OriginalCashPledge,
			dresses:              s.Dresses,
		}
		return strategy
	}

	if s.Type == CustomPrice {
		strategy := &customPriceSaleStrategy{
			originalCharterMoney: s.OriginalCharterMoney,
			originalCashPledge:   s.OriginalCashPledge,
			customCharterMoney:   s.CustomCharterMoney,
			customCashPledge:     s.CustomCashPledge,
		}
		return strategy
	}
	return nil
}
