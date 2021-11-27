package syslog

import (
	"WeddingDressManage/model"
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type AddDress struct {
	data      string
	TargetIds []int
}

func (a *AddDress) GetData(c *gin.Context)  {
	bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
	a.data = string(bodyBytes)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
}

func (a *AddDress) Logger()  {
	logModel := &model.OperationLog{
		Kind:        model.OperationType["addDress"],
		Data:        a.data,
	}

	secondaryLogModels := make([]*model.OperationSecondaryEntity, 0, len(a.TargetIds))

	for _, targetId := range a.TargetIds {
		secondaryLogModel := &model.OperationSecondaryEntity{
			SecondaryEntityType: model.SecondaryEntityType["dress"],
			SecondaryEntityId:   targetId,
		}
		secondaryLogModels = append(secondaryLogModels, secondaryLogModel)
	}


	logModel.SaveWithSecondaryLog(secondaryLogModels)
}
