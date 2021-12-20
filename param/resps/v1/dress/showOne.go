package dress

import "WeddingDressManage/business/v1/dress"

type ShowOneResponse struct {
	Category *showOneCategoryResponse `json:"category"`
	Dress    *showOneDressResponse    `json:"dress"`
}

type showOneCategoryResponse struct {
	Id               int                  `json:"id"`
	Kind             *showOneKindResponse `json:"kind"`
	SerialNumber     string               `json:"serialNumber"`
	Quantity         int                  `json:"quantity"`
	RentableQuantity int                  `json:"rentableQuantity"`
	CharterMoney     int                  `json:"charterMoney"`
	AvgCharterMoney  int                  `json:"avgCharterMoney"`
	CashPledge       int                  `json:"cashPledge"`
	RentCounter      int                  `json:"rentCounter"`
	LaundryCounter   int                  `json:"laundryCounter"`
	MaintainCounter  int                  `json:"maintainCounter"`
	CoverImg         string               `json:"coverImg"`
	SecondaryImg     []string             `json:"secondaryImg"`
	Status           string               `json:"status"`
}

type showOneDressResponse struct {
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

type showOneKindResponse struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Status string `json:"status"`
}

func (s *ShowOneResponse) Fill(dress *dress.Dress) {
	s.fillCategory(dress.Category)
	s.fillDress(dress)
}

func (s *ShowOneResponse) fillCategory(category *dress.Category) {
	s.Category = &showOneCategoryResponse{
		Id: category.Id,
		Kind: &showOneKindResponse{
			Id:     category.Kind.Id,
			Name:   category.Kind.Name,
			Code:   category.Kind.Code,
			Status: category.Kind.Status,
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

func (s *ShowOneResponse) fillDress(dress *dress.Dress) {
	s.Dress = &showOneDressResponse{
		Id:              dress.Id,
		SerialNumber:    dress.SerialNumber,
		Size:            dress.Size,
		RentCounter:     dress.RentCounter,
		LaundryCounter:  dress.LaundryCounter,
		MaintainCounter: dress.MaintainCounter,
		CoverImg:        dress.CoverImg,
		SecondaryImg:    dress.SecondaryImg,
		Status:          dress.Status,
	}
}
