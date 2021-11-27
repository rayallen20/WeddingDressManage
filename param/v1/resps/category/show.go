package category

import "WeddingDressManage/business/v1/dress"

// Response 响应体中的品类信息部分
type Response struct {
	Id               int           `json:"id"`
	Kind             *kindResponse `json:"kind"`
	SerialNumber     string        `json:"serialNumber"`
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

type kindResponse struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Status string `json:"status"`
}

func (r *Response) fill(category *dress.Category) {
	r.Id = category.Id
	r.Kind = &kindResponse{
		Id:     category.Kind.Id,
		Name:   category.Kind.Name,
		Code:   category.Kind.Code,
		Status: category.Kind.Status,
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

func (c Response) Generate(categories []*dress.Category) (resps []*Response) {
	resps = make([]*Response, 0, len(categories))
	for _, category := range categories {
		resp := &Response{}
		resp.fill(category)
		resps = append(resps, resp)
	}
	return resps
}
