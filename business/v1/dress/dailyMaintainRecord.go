package dress

import (
	"WeddingDressManage/lib/helper/sliceHelper"
	"WeddingDressManage/model"
	"time"
)

// MaintainPlanDurationDays 维护状态预计持续时间 单位:天
const MaintainPlanDurationDays = 3

type DailyMaintainRecord struct {
	Id                  int
	Source              string
	Dress               *Dress
	MaintainPositionImg []string
	Note                string
	StartMaintainDate   time.Time
	PlanEndMaintainDate time.Time
	EndMaintainDate     time.Time
	Status              string
}

func (d *DailyMaintainRecord) CreateORMForDailyMaintain() *model.MaintainRecord {
	return &model.MaintainRecord{
		Source:              d.Source,
		DressId:             d.Dress.Id,
		MaintainPositionImg: sliceHelper.ImpactSliceToStr(d.MaintainPositionImg, "|"),
		Note:                d.Note,
		StartMaintainDate:   d.StartMaintainDate,
		PlanEndMaintainDate: d.PlanEndMaintainDate,
		Status:              d.Status,
	}
}
