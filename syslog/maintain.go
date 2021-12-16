package syslog

import (
	"WeddingDressManage/model"
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type Maintain struct {
	data     string
	TargetId int
}

func (m *Maintain) GetData(c *gin.Context) {
	bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
	m.data = string(bodyBytes)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
}

func (m *Maintain) Logger() {
	logModel := &model.OperationLog{
		Kind:     model.OperationType["dailyMaintain"],
		TargetId: m.TargetId,
		Data:     m.data,
	}
	logModel.Save()
}
