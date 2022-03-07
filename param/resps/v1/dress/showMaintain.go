package dress

import (
	"WeddingDressManage/business/v1/dress"
	"WeddingDressManage/param/resps/v1/pagination"
)

type ShowMaintainResponse struct {
	Pagination *pagination.Response      `json:"pagination"`
	Maintains  []*maintainRecordResponse `json:"maintains"`
}

type maintainRecordResponse struct {
	Maintain *maintainResponse          `json:"maintain"`
	Dress    *showMaintainDressResponse `json:"dress"`
}

type maintainResponse struct {
	Id                  int      `json:"id"`
	Source              string   `json:"source"`
	StartMaintainDate   string   `json:"startMaintainDate"`
	PlanEndMaintainDate string   `json:"planEndMaintainDate"`
	MaintainPositionImg []string `json:"maintainPositionImg"`
	Note                string   `json:"note"`
	Status              string   `json:"status"`
}

type showMaintainDressResponse struct {
	Id              int                           `json:"id"`
	Category        *showMaintainCategoryResponse `json:"category"`
	SerialNumber    int                           `json:"serialNumber"`
	Size            string                        `json:"size"`
	RentCounter     int                           `json:"rentCounter"`
	LaundryCounter  int                           `json:"laundryCounter"`
	MaintainCounter int                           `json:"maintainCounter"`
	CoverImg        string                        `json:"coverImg"`
	SecondaryImg    []string                      `json:"secondaryImg"`
	Status          string                        `json:"status"`
}

type showMaintainCategoryResponse struct {
	Id               int                       `json:"id"`
	Kind             *showMaintainKindResponse `json:"kind"`
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

type showMaintainKindResponse struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Status string `json:"status"`
}

func (s *ShowMaintainResponse) Fill(maintains []*dress.MaintainRecord, currentPage, totalPage int, count int64, itemPerPage int) {
	s.Maintains = make([]*maintainRecordResponse, 0, len(maintains))
	for _, maintain := range maintains {
		maintainRecordResp := &maintainRecordResponse{
			Maintain: &maintainResponse{
				Id:                  maintain.Id,
				Source:              maintain.Source,
				StartMaintainDate:   maintain.StartMaintainDate.Format("2006-01-02"),
				PlanEndMaintainDate: maintain.PlanEndMaintainDate.Format("2006-01-02"),
				MaintainPositionImg: maintain.MaintainPositionImg,
				Note:                maintain.Note,
				Status:              maintain.Status,
			},
			Dress: &showMaintainDressResponse{
				Id: maintain.Dress.Id,
				Category: &showMaintainCategoryResponse{
					Id: maintain.Dress.Category.Id,
					Kind: &showMaintainKindResponse{
						Id:     maintain.Dress.Category.Kind.Id,
						Name:   maintain.Dress.Category.Kind.Name,
						Code:   maintain.Dress.Category.Kind.Code,
						Status: maintain.Dress.Category.Kind.Status,
					},
					SerialNumber:     maintain.Dress.Category.SerialNumber,
					Quantity:         maintain.Dress.Category.Quantity,
					RentableQuantity: maintain.Dress.Category.RentableQuantity,
					CharterMoney:     maintain.Dress.Category.CharterMoney,
					AvgCharterMoney:  maintain.Dress.Category.AvgCharterMoney,
					CashPledge:       maintain.Dress.Category.CashPledge,
					RentCounter:      maintain.Dress.Category.RentCounter,
					LaundryCounter:   maintain.Dress.Category.LaundryCounter,
					MaintainCounter:  maintain.Dress.Category.MaintainCounter,
					CoverImg:         maintain.Dress.Category.CoverImg,
					SecondaryImg:     maintain.Dress.Category.SecondaryImg,
					Status:           maintain.Dress.Category.Status,
				},
				SerialNumber:    maintain.Dress.SerialNumber,
				Size:            maintain.Dress.Size,
				RentCounter:     maintain.Dress.RentCounter,
				LaundryCounter:  maintain.Dress.LaundryCounter,
				MaintainCounter: maintain.Dress.MaintainCounter,
				CoverImg:        maintain.Dress.CoverImg,
				SecondaryImg:    maintain.Dress.SecondaryImg,
				Status:          maintain.Dress.Status,
			},
		}
		s.Maintains = append(s.Maintains, maintainRecordResp)
	}

	s.Pagination = &pagination.Response{
		CurrentPage: currentPage,
		ItemPerPage: itemPerPage,
		TotalPage:   totalPage,
		TotalItem:   count,
	}
}
