package dress

import (
	"WeddingDressManage/lib/helper/sliceHelper"
	"WeddingDressManage/lib/helper/urlHelper"
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/model"
	"WeddingDressManage/param/request/v1/dress"
	"WeddingDressManage/param/resps/v1/pagination"
	"errors"
	"gorm.io/gorm"
	"strings"
	"time"
)

// LaundryPlanDurationDays 送洗状态预计持续时间 单位:天
const LaundryPlanDurationDays = 3

type LaundryRecord struct {
	Id                int
	Dress             *Dress
	DirtyPositionImg  []string
	Note              string
	StartLaundryDate  time.Time
	DueEndLaundryDate time.Time
	EndLaundryDate    time.Time
	Status            string
}

func (l *LaundryRecord) CreateORMForLaundry() *model.LaundryRecord {
	return &model.LaundryRecord{
		DressId:           l.Dress.Id,
		DirtyPositionImg:  sliceHelper.ImpactSliceToStr(l.DirtyPositionImg, "|"),
		Note:              l.Note,
		StartLaundryDate:  l.StartLaundryDate,
		DueEndLaundryDate: l.DueEndLaundryDate,
		Status:            l.Status,
	}
}

func (l *LaundryRecord) Show(param *dress.ShowLaundryParam) (laundryRecords []*LaundryRecord, totalPage int, count int64, err error) {
	laundryRecordOrm := &model.LaundryRecord{}
	orms, err := laundryRecordOrm.FindUnderway(param.Pagination.CurrentPage, param.Pagination.ItemPerPage)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		dbErr := &sysError.DbError{RealError: err}
		return nil, 0, 0, dbErr
	}

	laundryRecords = make([]*LaundryRecord, 0, len(orms))
	for _, orm := range orms {
		laundryRecord := &LaundryRecord{}
		laundryRecord.fill(orm)
		laundryRecords = append(laundryRecords, laundryRecord)
	}

	count, err = laundryRecordOrm.CountUnderway()
	if err != nil {
		return nil, 0, 0, &sysError.DbError{RealError: err}
	}
	totalPage = pagination.CalcTotalPage(count, param.Pagination.ItemPerPage)
	return laundryRecords, totalPage, count, nil
}

func (l *LaundryRecord) GiveBack(param *dress.GiveBackParam) error {
	laundryOrm := &model.LaundryRecord{Id: param.Laundry.Id}
	err := laundryOrm.FindById()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return &sysError.DbError{RealError: err}
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &sysError.LaundryRecordNotExistError{NotExistId: param.Laundry.Id}
	}

	if laundryOrm.Dress == nil {
		return &sysError.DressNotExistError{Id: laundryOrm.DressId}
	}

	if laundryOrm.Dress.Status != model.DressStatus["laundry"] {
		return &sysError.DressIsNotLaunderingError{NotLaunderingId: laundryOrm.Dress.Id}
	}

	laundryOrm.Status = model.LaundryStatus["finish"]
	laundryOrm.Dress.Status = model.DressStatus["onSale"]
	categoryOrm := &model.DressCategory{
		Id:               laundryOrm.Dress.Category.Id,
		RentableQuantity: laundryOrm.Dress.Category.RentableQuantity + 1,
	}

	err = laundryOrm.Finish(laundryOrm.Dress, categoryOrm)
	if err != nil {
		return &sysError.DbError{RealError: err}
	}

	// 此处为防止日后变更导致控制器层面需要使用biz 故填充 此时暂无实际用途
	l.fill(laundryOrm)
	dressBiz := Dress{}
	dressBiz.fill(laundryOrm.Dress)

	return nil
}

func (l *LaundryRecord) fill(orm *model.LaundryRecord) {
	l.Id = orm.Id
	if orm.Dress != nil {
		l.Dress = &Dress{
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
	}

	if orm.Dress.Category != nil {
		l.Dress.Category = &Category{
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
			l.Dress.Category.Kind = &Kind{
				Id:     orm.Dress.Category.Kind.Id,
				Name:   orm.Dress.Category.Kind.Name,
				Code:   orm.Dress.Category.Kind.Code,
				Status: orm.Dress.Category.Kind.Status,
			}
		}
	}

	l.DirtyPositionImg = urlHelper.GenFullImgWebSites(strings.Split(orm.DirtyPositionImg, "|"))
	l.Note = orm.Note
	l.StartLaundryDate = orm.StartLaundryDate
	l.Status = orm.Status
}
