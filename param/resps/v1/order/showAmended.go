package order

import (
	"WeddingDressManage/business/v1/order"
	"WeddingDressManage/lib/helper/paramHelper"
	"strconv"
)

type ShowAmendedResponse struct {
	Order    *ShowAmendedOrder     `json:"order"`
	Customer *ShowAmendedCustomer  `json:"customer"`
	Dresses  []*ShowAmendedDresses `json:"dresses"`
}

type ShowAmendedOrder struct {
	Id          int    `json:"id"`
	WeddingDate string `json:"weddingDate"`
}

type ShowAmendedCustomer struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
}

type ShowAmendedDresses struct {
	Id           int                  `json:"id"`
	Category     *ShowAmendedCategory `json:"category"`
	Size         string               `json:"size"`
	SerialNumber int                  `json:"serialNumber"`
}

type ShowAmendedCategory struct {
	Id           int    `json:"id"`
	SerialNumber string `json:"serialNumber"`
	CharterMoney string `json:"charterMoney"`
}

func (s *ShowAmendedResponse) Fill(order *order.Order) {
	s.Order = &ShowAmendedOrder{
		Id:          order.Id,
		WeddingDate: order.WeddingDate.Format("2006-01-02"),
	}

	s.Customer = &ShowAmendedCustomer{
		Id:     order.Customer.Id,
		Name:   order.Customer.Name,
		Mobile: order.Customer.Mobile,
	}

	s.Dresses = make([]*ShowAmendedDresses, 0, len(order.Items))
	for _, item := range order.Items {
		dress := &ShowAmendedDresses{
			Id: item.Dress.Id,
			Category: &ShowAmendedCategory{
				Id:           item.Dress.Category.Id,
				SerialNumber: item.Dress.Category.SerialNumber,
				CharterMoney: paramHelper.ConvertPennyToYuan(strconv.Itoa(item.Dress.Category.CharterMoney)),
			},
			Size:         item.Dress.Size,
			SerialNumber: item.Dress.SerialNumber,
		}
		s.Dresses = append(s.Dresses, dress)
	}
}
