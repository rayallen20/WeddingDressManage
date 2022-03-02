package dress

import (
	"WeddingDressManage/business/v1/dress"
	"WeddingDressManage/param/resps/v1/pagination"
)

type ShowLaundryResponse struct {
	Pagination *pagination.Response     `json:"pagination"`
	Laundries  []*laundryRecordResponse `json:"laundries"`
}

type laundryRecordResponse struct {
	Laundry *laundryResponse `json:"laundry"`
	Dress   *dressResponse   `json:"dress"`
}

type laundryResponse struct {
	Id               int      `json:"id"`
	StartLaundryDate string   `json:"startLaundryDate"`
	DirtyPositionImg []string `json:"dirtyPositionImg"`
	Note             string   `json:"note"`
	Status           string   `json:"status"`
}

type dressResponse struct {
	Id              int               `json:"id"`
	Category        *categoryResponse `json:"category"`
	SerialNumber    int               `json:"serialNumber"`
	Size            string            `json:"size"`
	RentCounter     int               `json:"rentCounter"`
	LaundryCounter  int               `json:"laundryCounter"`
	MaintainCounter int               `json:"maintainCounter"`
	CoverImg        string            `json:"coverImg"`
	SecondaryImg    []string          `json:"secondaryImg"`
	Status          string            `json:"status"`
}

type categoryResponse struct {
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

func (s *ShowLaundryResponse) Fill(laundries []*dress.LaundryRecord, currentPage, totalPage, itemPerPage int) {
	s.Laundries = make([]*laundryRecordResponse, 0, len(laundries))
	for _, laundryBiz := range laundries {
		laundryResp := &laundryRecordResponse{
			Laundry: &laundryResponse{
				Id:               laundryBiz.Id,
				StartLaundryDate: laundryBiz.StartLaundryDate.Format("2006-01-02"),
				DirtyPositionImg: laundryBiz.DirtyPositionImg,
				Note:             laundryBiz.Note,
				Status:           laundryBiz.Status,
			},
			Dress: &dressResponse{
				Id: laundryBiz.Dress.Id,
				Category: &categoryResponse{
					Id: laundryBiz.Dress.Category.Id,
					Kind: &kindResponse{
						Id:     laundryBiz.Dress.Category.Kind.Id,
						Name:   laundryBiz.Dress.Category.Kind.Name,
						Code:   laundryBiz.Dress.Category.Kind.Code,
						Status: laundryBiz.Dress.Category.Kind.Status,
					},
					SerialNumber:     laundryBiz.Dress.Category.SerialNumber,
					Quantity:         laundryBiz.Dress.Category.Quantity,
					RentableQuantity: laundryBiz.Dress.Category.RentableQuantity,
					CharterMoney:     laundryBiz.Dress.Category.CharterMoney,
					AvgCharterMoney:  laundryBiz.Dress.Category.AvgCharterMoney,
					CashPledge:       laundryBiz.Dress.Category.CashPledge,
					RentCounter:      laundryBiz.Dress.Category.RentCounter,
					LaundryCounter:   laundryBiz.Dress.Category.LaundryCounter,
					MaintainCounter:  laundryBiz.Dress.Category.MaintainCounter,
					CoverImg:         laundryBiz.Dress.Category.CoverImg,
					SecondaryImg:     laundryBiz.Dress.Category.SecondaryImg,
					Status:           laundryBiz.Dress.Category.Status,
				},
				SerialNumber:    laundryBiz.Dress.SerialNumber,
				Size:            laundryBiz.Dress.Size,
				RentCounter:     laundryBiz.Dress.RentCounter,
				LaundryCounter:  laundryBiz.Dress.LaundryCounter,
				MaintainCounter: laundryBiz.Dress.MaintainCounter,
				CoverImg:        laundryBiz.Dress.CoverImg,
				SecondaryImg:    laundryBiz.Dress.SecondaryImg,
				Status:          laundryBiz.Dress.Status,
			},
		}
		s.Laundries = append(s.Laundries, laundryResp)
	}

	s.Pagination = &pagination.Response{
		CurrentPage: currentPage,
		ItemPerPage: itemPerPage,
		TotalPage:   totalPage,
	}
}
