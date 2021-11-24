package syslog

import (
	"WeddingDressManage/model"
)

type CreateCategory struct {
	Data string
	TargetId int
}

func (c *CreateCategory) Logger()  {
	logModel := &model.OperationLog{
		Kind:        model.OperationType["createCategoryAndDress"],
		TargetId:    c.TargetId,
		Data:        c.Data,
	}
	logModel.Save()
}