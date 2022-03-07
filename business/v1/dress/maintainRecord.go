package dress

import (
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/model"
	requestParam "WeddingDressManage/param/request/v1/dress"
	"WeddingDressManage/param/resps/v1/pagination"
	"time"
)

// MaintainPlanDurationDays 维护状态预计持续时间 单位:天
const MaintainPlanDurationDays = 3

type MaintainRecord struct {
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

func (m *MaintainRecord) CreateORMForDailyMaintain() *model.MaintainRecord {
	dailyMaintainRecord := &dailyMaintainRecord{MaintainRecord: m}
	return dailyMaintainRecord.createORMForDailyMaintain()
}

func (m *MaintainRecord) Show(param *requestParam.ShowMaintainParam) (maintainRecords []*MaintainRecord, totalPage int, count int64, err error) {
	orm := &model.MaintainRecord{}
	maintainRecordOrms, err := orm.FindUnderway(param.Pagination.CurrentPage, param.Pagination.ItemPerPage)
	if err != nil {
		return nil, 0, 0, &sysError.DbError{RealError: err}
	}

	maintainRecords = make([]*MaintainRecord, 0, len(maintainRecordOrms))
	for _, maintainRecordOrm := range maintainRecordOrms {
		if maintainRecordOrm.Source == model.MaintainSource["daily"] {
			dailyMaintainRecord := &dailyMaintainRecord{}
			dailyMaintainRecord.fill(maintainRecordOrm)
			maintainRecord := dailyMaintainRecord.MaintainRecord
			maintainRecords = append(maintainRecords, maintainRecord)
		} else if maintainRecordOrm.Source == model.MaintainSource["item"] {
			// TODO:此处需与订单关联 待订单模块完成后实现
		}
	}

	count, err = orm.CountUnderway()
	if err != nil {
		return nil, 0, 0, &sysError.DbError{RealError: err}
	}
	totalPage = pagination.CalcTotalPage(count, param.Pagination.ItemPerPage)

	return maintainRecords, totalPage, count, nil
}
