package resps

import "WeddingDressManage/business/v1/dress"

// CategoryResponse 响应体中的品类信息部分
type CategoryResponse struct {
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

func (c *CategoryResponse) fill(category *dress.Category) {
	c.Id = category.Id
	c.Kind = &kindResponse{
		Id:     category.Kind.Id,
		Name:   category.Kind.Name,
		Code:   category.Kind.Code,
		Status: category.Kind.Status,
	}
	c.SerialNumber = category.SerialNumber
	c.Quantity = category.Quantity
	c.RentableQuantity = category.RentableQuantity
	c.CharterMoney = category.CharterMoney
	c.AvgCharterMoney = category.AvgCharterMoney
	c.CashPledge = category.CashPledge
	c.RentCounter = category.RentCounter
	c.LaundryCounter = category.LaundryCounter
	c.MaintainCounter = category.MaintainCounter
	c.CoverImg = category.CoverImg
	c.SecondaryImg = category.SecondaryImg
	c.Status = category.Status
}

func (c CategoryResponse) Generate(categories []*dress.Category) (resps []*CategoryResponse) {
	resps = make([]*CategoryResponse, 0, len(categories))
	for _, category := range categories {
		resp := &CategoryResponse{}
		resp.fill(category)
		resps = append(resps, resp)
	}
	return resps
}
