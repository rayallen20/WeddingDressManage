package syslog

import (
	"WeddingDressManage/model"
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type Laundry struct {
	data     string
	TargetId int
}

func (l *Laundry) GetData(c *gin.Context) {
	bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
	l.data = string(bodyBytes)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
}

func (l *Laundry) Logger() {
	logModel := &model.OperationLog{
		Kind:     model.OperationType["laundry"],
		TargetId: l.TargetId,
		Data:     l.data,
	}
	logModel.Save()
}
