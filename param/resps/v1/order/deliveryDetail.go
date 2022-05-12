package order

import (
	"WeddingDressManage/business/v1/order"
	"WeddingDressManage/lib/helper/paramHelper"
	"strconv"
)

type DeliveryDetailResponse struct {
	Order *DeliveryDetailOrder `json:"order"`
}

type DeliveryDetailOrder struct {
	Id                  int                     `json:"id"`
	Customer            *DeliveryDetailCustomer `json:"customer"`
	SerialNumber        string                  `json:"serialNumber"`
	WeddingDate         string                  `json:"weddingDate"`
	DuePayCharterMoney  string                  `json:"duePayCharterMoney"`
	DuePayCashPledge    string                  `json:"duePayCashPledge"`
	ActualPayCashPledge string                  `json:"actualPayCashPledge"`
	Items               []*DeliveryDetailItem   `json:"items"`
}

type DeliveryDetailCustomer struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
}

type DeliveryDetailItem struct {
	Id     int                  `json:"id"`
	Dress  *DeliveryDetailDress `json:"dress"`
	Status string               `json:"status"`
}

type DeliveryDetailDress struct {
	Id           int                     `json:"id"`
	Category     *DeliveryDetailCategory `json:"category"`
	SerialNumber int                     `json:"serialNumber"`
	CoverImg     string                  `json:"coverImg"`
	SecondaryImg []string                `json:"secondaryImg"`
}

type DeliveryDetailCategory struct {
	Id           int    `json:"id"`
	SerialNumber string `json:"serialNumber"`
}

func (d *DeliveryDetailResponse) Fill(order *order.Order) {
	d.Order = &DeliveryDetailOrder{
		Id: order.Id,
		Customer: &DeliveryDetailCustomer{
			Id:     order.Customer.Id,
			Name:   order.Customer.Name,
			Mobile: order.Customer.Mobile,
		},
		SerialNumber:        order.SerialNumber,
		WeddingDate:         order.WeddingDate.Format("2006-01-02"),
		DuePayCharterMoney:  paramHelper.ConvertPennyToYuan(strconv.Itoa(order.DuePayCharterMoney)),
		DuePayCashPledge:    paramHelper.ConvertPennyToYuan(strconv.Itoa(order.DuePayCashPledge)),
		ActualPayCashPledge: paramHelper.ConvertPennyToYuan(strconv.Itoa(order.ActualPayCashPledge)),
	}
	d.Order.Items = make([]*DeliveryDetailItem, 0, len(order.Items))
	for _, itemBiz := range order.Items {
		itemResp := &DeliveryDetailItem{
			Id: itemBiz.Id,
			Dress: &DeliveryDetailDress{
				Id: itemBiz.Dress.Id,
				Category: &DeliveryDetailCategory{
					Id:           itemBiz.Dress.Category.Id,
					SerialNumber: itemBiz.Dress.Category.SerialNumber,
				},
				SerialNumber: itemBiz.Dress.SerialNumber,
				CoverImg:     itemBiz.Dress.CoverImg,
				SecondaryImg: itemBiz.Dress.SecondaryImg,
			},
			Status: itemBiz.RentStatus,
		}
		d.Order.Items = append(d.Order.Items, itemResp)
	}
}
