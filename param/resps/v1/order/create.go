package order

import (
	"WeddingDressManage/business/v1/order"
	"WeddingDressManage/lib/helper/paramHelper"
	"strconv"
)

type CreateResponse struct {
	SerialNumber          string `json:"serialNumber"`
	ActualPayCharterMoney string `json:"actualPayCharterMoney"`
	ActualPayCashPledge   string `json:"actualPayCashPledge"`
	PledgeIsSettled       bool   `json:"pledgeIsSettled"`
}

type CreateOrderResponse struct {
}

func (c *CreateResponse) Fill(order *order.Order, PledgeIsSettled bool) {
	c.SerialNumber = order.SerialNumber
	charterMoneyStr := strconv.Itoa(order.ActualPayCharterMoney)
	c.ActualPayCharterMoney = paramHelper.ConvertPennyToYuan(charterMoneyStr)
	cashPledgeMoneyStr := strconv.Itoa(order.ActualPayCashPledge)
	c.ActualPayCashPledge = paramHelper.ConvertPennyToYuan(cashPledgeMoneyStr)
	c.PledgeIsSettled = PledgeIsSettled
}
