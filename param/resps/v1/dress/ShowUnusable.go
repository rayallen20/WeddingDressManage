package dress

import (
	"WeddingDressManage/business/v1/dress"
	"WeddingDressManage/param/resps/v1/pagination"
)

type ShowUnusableResponse struct {
	Category   *showUnusableCategoryResponse  `json:"category"`
	Dresses    []*showUnusableDressesResponse `json:"dresses"`
	Pagination *pagination.Response           `json:"pagination"`
}

type showUnusableCategoryResponse struct {
	Id               int                       `json:"id"`
	Kind             *showUnusableKindResponse `json:"kind"`
	SerialNumber     string                    `json:"serialNumber"`
	Quantity         int                       `json:"quantity"`
	RentableQuantity int                       `json:"rentableQuantity"`
	CharterMoney     int                       `json:"charterMoney"`
	AvgCharterMoney  int                       `json:"avgCharterMoney"`
	CashPledge       int                       `json:"cashPledge"`
	RentCounter      int                       `json:"rentCounter"`
	LaundryCounter   int                       `json:"laundryCounter"`
	MaintainCounter  int                       `json:"maintainCounter"`
	CoverImg         string                    `json:"coverImg"`
	SecondaryImg     []string                  `json:"secondaryImg"`
	Status           string                    `json:"status"`
}

type showUnusableKindResponse struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Status string `json:"status"`
}

type showUnusableDressesResponse struct {
	Id              int      `json:"id"`
	SerialNumber    int      `json:"serialNumber"`
	Size            string   `json:"size"`
	RentCounter     int      `json:"rentCounter"`
	LaundryCounter  int      `json:"laundryCounter"`
	MaintainCounter int      `json:"maintainCounter"`
	CoverImg        string   `json:"coverImg"`
	SecondaryImg    []string `json:"secondaryImg"`
	Status          string   `json:"status"`
}

func (s *ShowUnusableResponse) Fill(category *dress.Category, dresses []*dress.Dress, pagination *pagination.Response) {
	s.fillCategory(category)
	s.fillDresses(dresses)
	s.Pagination = pagination
}

func (s *ShowUnusableResponse) fillCategory(category *dress.Category) {
	s.Category = &showUnusableCategoryResponse{
		Id: category.Id,
		Kind: &showUnusableKindResponse{
			Id:     category.Kind.Id,
			Name:   category.Kind.Name,
			Code:   category.Kind.Code,
			Status: category.Status,
		},
		SerialNumber:     category.SerialNumber,
		Quantity:         category.Quantity,
		RentableQuantity: category.RentableQuantity,
		CharterMoney:     category.CharterMoney,
		AvgCharterMoney:  category.AvgCharterMoney,
		CashPledge:       category.CashPledge,
		RentCounter:      category.RentCounter,
		LaundryCounter:   category.LaundryCounter,
		MaintainCounter:  category.MaintainCounter,
		CoverImg:         category.CoverImg,
		SecondaryImg:     category.SecondaryImg,
		Status:           category.Status,
	}
}

func (s *ShowUnusableResponse) fillDresses(dresses []*dress.Dress) {
	s.Dresses = make([]*showUnusableDressesResponse, 0, len(dresses))
	for _, dressBiz := range dresses {
		dressResp := &showUnusableDressesResponse{
			Id:              dressBiz.Id,
			SerialNumber:    dressBiz.SerialNumber,
			Size:            dressBiz.Size,
			RentCounter:     dressBiz.RentCounter,
			LaundryCounter:  dressBiz.LaundryCounter,
			MaintainCounter: dressBiz.MaintainCounter,
			CoverImg:        dressBiz.CoverImg,
			SecondaryImg:    dressBiz.SecondaryImg,
			Status:          dressBiz.Status,
		}
		s.Dresses = append(s.Dresses, dressResp)
	}
}
