package order

import (
	"WeddingDressManage/business/v1/order"
	"WeddingDressManage/lib/helper/paramHelper"
	"strconv"
)

type DiscountResponse struct {
	Order *DiscountOrderResponse `json:"order"`
}

type DiscountOrderResponse struct {
	OriginalCharterMoney string `json:"originalCharterMoney"`
	OriginalCashPledge   string `json:"originalCashPledge"`
	Discount             string `json:"discount"`
	DuePayCharterMoney   string `json:"duePayCharterMoney"`
	DuePayCashPledge     string `json:"duePayCashPledge"`
}

func (d *DiscountResponse) Fill(order *order.Order) {
	d.Order = &DiscountOrderResponse{}
	d.Order.OriginalCharterMoney = paramHelper.ConvertPennyToYuan(strconv.Itoa(order.OriginalCharterMoney))
	d.Order.OriginalCashPledge = paramHelper.ConvertPennyToYuan(strconv.Itoa(order.OriginalCashPledge))
	d.Order.DuePayCharterMoney = paramHelper.ConvertPennyToYuan(strconv.Itoa(order.DuePayCharterMoney))
	d.Order.DuePayCashPledge = paramHelper.ConvertPennyToYuan(strconv.Itoa(order.DuePayCashPledge))
	d.Order.Discount = strconv.FormatFloat(order.Discount, 'f', 2, 64)
}
