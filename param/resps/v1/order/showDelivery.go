package order

import (
	"WeddingDressManage/business/v1/order"
	"WeddingDressManage/lib/helper/paramHelper"
	"WeddingDressManage/param/resps/v1/pagination"
	"strconv"
)

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
	CanBeChanged        bool                  `json:"canBeChanged"`
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
			PledgeSettled:       orderBiz.PledgeSettledStatus,
			CanBeChanged:        orderBiz.CanBeChanged,
			CanBeBatchDelivery:  orderBiz.CanBeBatchDelivery,
		}
		s.Orders = append(s.Orders, orderResp)
	}
}
