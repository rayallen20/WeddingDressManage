package order

import (
	"WeddingDressManage/business/v1/order"
	"WeddingDressManage/lib/helper/paramHelper"
	"WeddingDressManage/model"
	"WeddingDressManage/param/resps/v1/pagination"
	"strconv"
)

// PledgeSettled 押金支付情况
// true 押金已收
// false 押金未收
// fractional 已收部分押金 未收全
var PledgeSettled map[string]string = map[string]string{
	"true":       "true",
	"false":      "false",
	"fractional": "fractional",
}

type ShowDeliveryResponse struct {
	Orders     []*ShowDeliveryOrder `json:"orders"`
	Pagination *pagination.Response `json:"pagination"`
}

type ShowDeliveryOrder struct {
	Id                  int                   `json:"id"`
	Customer            *ShowDeliveryCustomer `json:"customer"`
	SerialNumber        string                `json:"serialNumber"`
	WeddingDate         string                `json:"weddingDate"`
	DuePayCharterMoney  string                `json:"duePayCharterMoney"`
	DuePayCashPledge    string                `json:"duePayCashPledge"`
	ActualPayCashPledge string                `json:"actualPayCashPledge"`
	PledgeSettled       string                `json:"pledgeSettled"`
	CanBeChanged        bool                  `json:"canBeChange"`
	CanBeBatchDelivery  bool                  `json:"canBeBatchDelivery"`
}

type ShowDeliveryCustomer struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
}

func (s *ShowDeliveryResponse) Fill(orders []*order.Order, pagination *pagination.Response) {
	s.Orders = make([]*ShowDeliveryOrder, 0, len(orders))
	s.fillOrders(orders)
	s.Pagination = pagination
}

func (s *ShowDeliveryResponse) fillOrders(orders []*order.Order) {
	for _, orderBiz := range orders {
		orderResp := &ShowDeliveryOrder{
			Id: orderBiz.Id,
			Customer: &ShowDeliveryCustomer{
				Id:     orderBiz.Customer.Id,
				Name:   orderBiz.Customer.Name,
				Mobile: orderBiz.Customer.Mobile,
			},
			SerialNumber:        orderBiz.SerialNumber,
			WeddingDate:         orderBiz.WeddingDate.Format("2006-01-02"),
			DuePayCharterMoney:  paramHelper.ConvertPennyToYuan(strconv.Itoa(orderBiz.DuePayCharterMoney)),
			DuePayCashPledge:    paramHelper.ConvertPennyToYuan(strconv.Itoa(orderBiz.DuePayCashPledge)),
			ActualPayCashPledge: paramHelper.ConvertPennyToYuan(strconv.Itoa(orderBiz.ActualPayCashPledge)),
		}

		// 判断押金支付情况
		if orderBiz.ActualPayCashPledge == 0 {
			orderResp.PledgeSettled = PledgeSettled["false"]
		}

		if orderBiz.ActualPayCashPledge != 0 && orderBiz.ActualPayCashPledge != orderBiz.DuePayCashPledge {
			orderResp.PledgeSettled = PledgeSettled["fractional"]
		}

		if orderBiz.ActualPayCashPledge == orderBiz.DuePayCashPledge {
			orderResp.PledgeSettled = PledgeSettled["true"]
		}

		// 判断是否可以修改订单
		// 判断标准:仅在未开始支付押金前可修改订单
		if orderBiz.ActualPayCashPledge == 0 {
			orderResp.CanBeChanged = true
		} else {
			orderResp.CanBeChanged = false
		}

		// 判断是否可以批次出件
		// 判断标准:若优惠策略为打折或原价 则可以批次出件 若优惠策略为自定义租金与押金 则不可以批次出件
		if orderBiz.SaleStrategy == model.SaleStrategy["discount"] || orderBiz.SaleStrategy == model.SaleStrategy["originalPrice"] {
			orderResp.CanBeBatchDelivery = true
		}

		if orderBiz.SaleStrategy == model.SaleStrategy["customPrice"] {
			orderResp.CanBeBatchDelivery = false
		}
		s.Orders = append(s.Orders, orderResp)
	}
}
