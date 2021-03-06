package category

import "WeddingDressManage/business/v1/dress"

// ShowResponse 响应体中的品类信息部分
type ShowResponse struct {
	Id               int               `json:"id"`
	Kind             *showKindResponse `json:"kind"`
	SerialNumber     string            `json:"serialNumber"`
	Quantity         int           `json:"quantity"`
	RentableQuantity int           `json:"rentableQuantity"`
	CharterMoney     int           `json:"charterMoney"`
	AvgCharterMoney  int           `json:"avgCharterMoney"`
	CashPledge       int           `json:"cashPledge"`
	RentCounter      int           `json:"rentCounter"`
	LaundryCounter   int           `json:"laundryCounter"`
	MaintainCounter  int           `json:"maintainCounter"`
	CoverImg         string        `json:"coverImg"`
	SecondaryImg     []string      `json:"secondaryImg"`
	Status           string        `json:"status"`
}

type showKindResponse struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Status string `json:"status"`
}

func (r *ShowResponse) fill(category *dress.Category) {
	r.Id = category.Id
	if category.Kind != nil {
		r.Kind = &showKindResponse{
			Id:     category.Kind.Id,
			Name:   category.Kind.Name,
			Code:   category.Kind.Code,
			Status: category.Kind.Status,
		}
	}
	r.SerialNumber = category.SerialNumber
	r.Quantity = category.Quantity
	r.RentableQuantity = category.RentableQuantity
	r.CharterMoney = category.CharterMoney
	r.AvgCharterMoney = category.AvgCharterMoney
	r.CashPledge = category.CashPledge
	r.RentCounter = category.RentCounter
	r.LaundryCounter = category.LaundryCounter
	r.MaintainCounter = category.MaintainCounter
	r.CoverImg = category.CoverImg
	r.SecondaryImg = category.SecondaryImg
	r.Status = category.Status
}

func (r *ShowResponse) Generate(categories []*dress.Category) (resps []*ShowResponse) {
	resps = make([]*ShowResponse, 0, len(categories))
	for _, category := range categories {
		resp := &ShowResponse{}
		resp.fill(category)
		resps = append(resps, resp)
	}
	return resps
}
