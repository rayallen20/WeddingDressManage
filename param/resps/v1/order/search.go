package order

import "WeddingDressManage/business/v1/dress"

type SearchResponse struct {
	Id               int                 `json:"id"`
	Kind             *searchKindResponse `json:"kind"`
	SerialNumber     string              `json:"serialNumber"`
	Quantity         int                 `json:"quantity"`
	RentableQuantity int                 `json:"rentableQuantity"`
	CharterMoney     int                 `json:"charterMoney"`
	AvgCharterMoney  int                 `json:"avgCharterMoney"`
	CashPledge       int                 `json:"cashPledge"`
	RentCounter      int                 `json:"rentCounter"`
	LaundryCounter   int                 `json:"laundryCounter"`
	MaintainCounter  int                 `json:"maintainCounter"`
	CoverImg         string              `json:"coverImg"`
	SecondaryImg     []string            `json:"secondaryImg"`
	Status           string              `json:"status"`
}

type searchKindResponse struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Status string `json:"status"`
}

func (s *SearchResponse) fill(category *dress.Category) {
	s.Id = category.Id
	if category.Kind != nil {
		s.Kind = &searchKindResponse{
			Id:     category.Kind.Id,
			Name:   category.Kind.Name,
			Code:   category.Kind.Code,
			Status: category.Kind.Status,
		}
	}
	s.SerialNumber = category.SerialNumber
	s.Quantity = category.Quantity
	s.RentableQuantity = category.RentableQuantity
	s.CharterMoney = category.CharterMoney
	s.AvgCharterMoney = category.AvgCharterMoney
	s.CashPledge = category.CashPledge
	s.RentCounter = category.RentCounter
	s.LaundryCounter = category.LaundryCounter
	s.MaintainCounter = category.MaintainCounter
	s.CoverImg = category.CoverImg
	s.SecondaryImg = category.SecondaryImg
	s.Status = category.Status
}

func (s *SearchResponse) Generate(categories []*dress.Category) (resps []*SearchResponse) {
	resps = make([]*SearchResponse, 0, len(categories))
	for _, category := range categories {
		resp := &SearchResponse{}
		resp.fill(category)
		resps = append(resps, resp)
	}
	return resps
}
