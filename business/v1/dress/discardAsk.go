package dress

import (
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/model"
)

type DiscardAsk struct {
	Dress *Dress
	Note  string
}

func (a *DiscardAsk) Apply() error {
	orm := &model.DressDiscardRecord{
		DressCategoryId: a.Dress.CategoryId,
		DressId:         a.Dress.Id,
		Note:            a.Note,
		Status:          model.DiscardRecordStatus["pendingApproval"],
	}

	err := orm.Save()
	if err != nil {
		return &sysError.DbError{RealError: err}
	}

	return nil
}
