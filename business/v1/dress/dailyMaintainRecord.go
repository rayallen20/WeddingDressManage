package dress

import (
	"WeddingDressManage/lib/helper/sliceHelper"
	"WeddingDressManage/lib/helper/urlHelper"
	"WeddingDressManage/model"
	"strings"
)

type dailyMaintainRecord struct {
	maintainRecord MaintainRecord
}

func (d *dailyMaintainRecord) createORMForDailyMaintain() *model.MaintainRecord {
	return &model.MaintainRecord{
		Source:              d.maintainRecord.Source,
		DressId:             d.maintainRecord.Dress.Id,
		MaintainPositionImg: sliceHelper.ImpactSliceToStr(d.maintainRecord.MaintainPositionImg, "|"),
		Note:                d.maintainRecord.Note,
		StartMaintainDate:   d.maintainRecord.StartMaintainDate,
		PlanEndMaintainDate: d.maintainRecord.PlanEndMaintainDate,
		Status:              d.maintainRecord.Status,
	}
}

func (d *dailyMaintainRecord) fill(orm *model.MaintainRecord) {
	d.maintainRecord.Id = orm.Id
	d.maintainRecord.Source = orm.Source
	if orm.Dress != nil {
		d.maintainRecord.Dress = &Dress{
			Id:         orm.Dress.Id,
			CategoryId: orm.Dress.CategoryId,
			Category: &Category{
				Id: orm.Dress.Category.Id,
				Kind: &Kind{
					Id:     orm.Dress.Category.Kind.Id,
					Name:   orm.Dress.Category.Kind.Name,
					Code:   orm.Dress.Category.Kind.Code,
					Status: orm.Dress.Category.Kind.Status,
				},
				SerialNumber:     orm.Dress.Category.SerialNumber,
				Quantity:         orm.Dress.Category.Quantity,
				RentableQuantity: orm.Dress.Category.RentableQuantity,
				CharterMoney:     orm.Dress.Category.CharterMoney,
				AvgCharterMoney:  orm.Dress.Category.AvgCharterMoney,
				CashPledge:       orm.Dress.Category.CashPledge,
				RentCounter:      orm.Dress.Category.RentCounter,
				LaundryCounter:   orm.Dress.Category.LaundryCounter,
				MaintainCounter:  orm.Dress.Category.MaintainCounter,
				CoverImg:         orm.Dress.Category.CoverImg,
				SecondaryImg:     urlHelper.GenFullImgWebSites(strings.Split(orm.Dress.Category.SecondaryImg, "|")),
				Status:           orm.Dress.Category.Status,
			},
			SerialNumber:    orm.Dress.SerialNumber,
			Size:            orm.Dress.Size,
			RentCounter:     orm.Dress.RentCounter,
			LaundryCounter:  orm.Dress.LaundryCounter,
			MaintainCounter: orm.Dress.MaintainCounter,
			CoverImg:        urlHelper.GenFullImgWebSite(orm.Dress.CoverImg),
			SecondaryImg:    urlHelper.GenFullImgWebSites(strings.Split(orm.Dress.SecondaryImg, "|")),
			Status:          orm.Dress.Status,
		}
	}
	d.maintainRecord.MaintainPositionImg = urlHelper.GenFullImgWebSites(strings.Split(orm.MaintainPositionImg, "|"))
	d.maintainRecord.Note = orm.Note
	d.maintainRecord.StartMaintainDate = orm.StartMaintainDate
	d.maintainRecord.PlanEndMaintainDate = orm.PlanEndMaintainDate
	d.maintainRecord.Status = orm.Status
}
