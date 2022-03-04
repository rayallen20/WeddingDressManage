package syslog

import (
	"WeddingDressManage/model"
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type LaundryGiveBack struct {
	data     string
	TargetId int
}

func (l *LaundryGiveBack) GetData(c *gin.Context) {
	bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
	l.data = string(bodyBytes)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
}

func (l *LaundryGiveBack) Logger() {
	logModel := &model.OperationLog{
		Kind:     model.OperationType["itemMaintainGiveBack"],
		TargetId: l.TargetId,
		Data:     l.data,
	}
	logModel.Save()
}
