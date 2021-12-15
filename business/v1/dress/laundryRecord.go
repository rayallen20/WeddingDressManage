package dress

import (
	"WeddingDressManage/lib/helper/sliceHelper"
	"WeddingDressManage/model"
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
