package order

import (
	"WeddingDressManage/business/v1/order"
	"strconv"
)

type PreCreateResponse struct {
	Order *PreCreateOrderResponse `json:"order"`
}

type PreCreateOrderResponse struct {
	Items                []*PreCreateItemResponse `json:"items"`
	OriginalCharterMoney string                   `json:"originalCharterMoney"`
	OriginalCashPledge   string                   `json:"originalCashPledge"`
}

type PreCreateItemResponse struct {
	Dress       *PreCreateItemDressResponse `json:"dress"`
	Status      string                      `json:"status"`
	MaintainFee string                      `json:"maintainFee"`
}

type PreCreateItemDressResponse struct {
	Id              int                            `json:"id"`
	Category        *PreCreateItemCategoryResponse `json:"category"`
	SerialNumber    int                            `json:"serialNumber"`
	Size            string                         `json:"size"`
	RentCounter     int                            `json:"rentCounter"`
	LaundryCounter  int                            `json:"laundryCounter"`
	MaintainCounter int                            `json:"maintainCounter"`
	CoverImg        string                         `json:"coverImg"`
	SecondaryImg    []string                       `json:"secondaryImg"`
	Status          string                         `json:"status"`
}

type PreCreateItemCategoryResponse struct {
	Id               int                        `json:"id"`
	Kind             *PreCreateItemKindResponse `json:"kind"`
	SerialNumber     string                     `json:"serialNumber"`
	Quantity         int                        `json:"quantity"`
	RentableQuantity int                        `json:"rentableQuantity"`
	CharterMoney     int                        `json:"charterMoney"`
	AvgCharterMoney  int                        `json:"avgCharterMoney"`
	CashPledge       int                        `json:"cashPledge"`
	RentCounter      int                        `json:"rentCounter"`
	LaundryCounter   int                        `json:"laundryCounter"`
	MaintainCounter  int                        `json:"maintainCounter"`
	CoverImg         string                     `json:"coverImg"`
	SecondaryImg     []string                   `json:"secondaryImg"`
	Status           string                     `json:"status"`
}

type PreCreateItemKindResponse struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Status string `json:"status"`
}

func (p *PreCreateResponse) Fill(order order.Order) {
	p.fillItems(order.Items)
}

func (p *PreCreateResponse) fillItems(items []*order.Item) {
	p.Order.Items = make([]*PreCreateItemResponse, 0, len(items))
	for _, item := range items {
		itemResponse := &PreCreateItemResponse{
			Dress: &PreCreateItemDressResponse{
				Id: item.Dress.Id,
				Category: &PreCreateItemCategoryResponse{
					Id: item.Dress.Category.Id,
					Kind: &PreCreateItemKindResponse{
						Id:     item.Dress.Category.Kind.Id,
						Name:   item.Dress.Category.Kind.Name,
						Code:   item.Dress.Category.Kind.Code,
						Status: item.Dress.Category.Kind.Status,
					},
					SerialNumber:     item.Dress.Category.SerialNumber,
					Quantity:         item.Dress.Category.Quantity,
					RentableQuantity: item.Dress.Category.RentableQuantity,
					CharterMoney:     item.Dress.Category.CharterMoney,
					AvgCharterMoney:  item.Dress.Category.AvgCharterMoney,
					CashPledge:       item.Dress.Category.CashPledge,
					RentCounter:      item.Dress.Category.RentCounter,
					LaundryCounter:   item.Dress.Category.LaundryCounter,
					MaintainCounter:  item.Dress.Category.MaintainCounter,
					CoverImg:         item.Dress.Category.CoverImg,
					SecondaryImg:     item.Dress.Category.SecondaryImg,
					Status:           item.Dress.Category.Status,
				},
				SerialNumber:    item.Dress.SerialNumber,
				Size:            item.Dress.Size,
				RentCounter:     item.Dress.RentCounter,
				LaundryCounter:  item.Dress.LaundryCounter,
				MaintainCounter: item.Dress.MaintainCounter,
				CoverImg:        item.Dress.CoverImg,
				SecondaryImg:    item.Dress.SecondaryImg,
				Status:          item.Dress.Status,
			},
			Status:      item.Status,
			MaintainFee: strconv.Itoa(item.MaintainFee),
		}
		p.Order.Items = append(p.Order.Items, itemResponse)
	}
}

func (p PreCreateResponse) fill() {

}
