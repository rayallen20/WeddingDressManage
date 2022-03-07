package dress

import (
	"WeddingDressManage/lib/helper/sliceHelper"
	"WeddingDressManage/lib/helper/urlHelper"
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/model"
	"strings"
)

type dailyMaintainRecord struct {
	MaintainRecord *MaintainRecord
}

func (d *dailyMaintainRecord) createORMForDailyMaintain() *model.MaintainRecord {
	return &model.MaintainRecord{
		Source:              d.MaintainRecord.Source,
		DressId:             d.MaintainRecord.Dress.Id,
		MaintainPositionImg: sliceHelper.ImpactSliceToStr(d.MaintainRecord.MaintainPositionImg, "|"),
		Note:                d.MaintainRecord.Note,
		StartMaintainDate:   d.MaintainRecord.StartMaintainDate,
		PlanEndMaintainDate: d.MaintainRecord.PlanEndMaintainDate,
		Status:              d.MaintainRecord.Status,
	}
}

func (d *dailyMaintainRecord) fill(orm *model.MaintainRecord) {
	d.MaintainRecord = &MaintainRecord{
		Id:                  orm.Id,
		Source:              orm.Source,
		MaintainPositionImg: urlHelper.GenFullImgWebSites(strings.Split(orm.MaintainPositionImg, "|")),
		Note:                orm.Note,
		StartMaintainDate:   orm.StartMaintainDate,
		PlanEndMaintainDate: orm.PlanEndMaintainDate,
		Status:              orm.Status,
	}

	if orm.Dress != nil {
		d.MaintainRecord.Dress = &Dress{
			Id:              orm.Dress.Id,
			CategoryId:      orm.Dress.CategoryId,
			SerialNumber:    orm.Dress.SerialNumber,
			Size:            orm.Dress.Size,
			RentCounter:     orm.Dress.RentCounter,
			LaundryCounter:  orm.Dress.LaundryCounter,
			MaintainCounter: orm.Dress.MaintainCounter,
			CoverImg:        urlHelper.GenFullImgWebSite(orm.Dress.CoverImg),
			SecondaryImg:    urlHelper.GenFullImgWebSites(strings.Split(orm.Dress.SecondaryImg, "|")),
			Status:          orm.Dress.Status,
		}

		if orm.Dress.Category != nil {
			d.MaintainRecord.Dress.Category = &Category{
				Id:               orm.Dress.Category.Id,
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
			}

			if orm.Dress.Category.Kind != nil {
				d.MaintainRecord.Dress.Category.Kind = &Kind{
					Id:     orm.Dress.Category.Kind.Id,
					Name:   orm.Dress.Category.Kind.Name,
					Code:   orm.Dress.Category.Kind.Code,
					Status: orm.Dress.Category.Kind.Status,
				}
			}
		}
	}
}

func (d *dailyMaintainRecord) giveBack(orm *model.MaintainRecord) (*dailyMaintainRecord, error) {
	orm.Status = model.MaintainStatus["finish"]
	orm.Dress.Status = model.DressStatus["onSale"]
	categoryOrm := &model.DressCategory{
		Id:               orm.Dress.Category.Id,
		RentableQuantity: orm.Dress.Category.RentableQuantity + 1,
	}

	err := orm.Finish(orm.Dress, categoryOrm)
	if err != nil {
		return nil, &sysError.DbError{RealError: err}
	}

	d.fill(orm)
	return d, nil
}
