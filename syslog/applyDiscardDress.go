package syslog

import (
	"WeddingDressManage/model"
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type ApplyDiscardDress struct {
	data     string
	TargetId int
}

func (a *ApplyDiscardDress) GetData(c *gin.Context) {
	bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
	a.data = string(bodyBytes)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
}

func (a *ApplyDiscardDress) Logger() {
	logModel := &model.OperationLog{
		Kind:     model.OperationType["applyDiscardDress"],
		TargetId: a.TargetId,
		Data:     a.data,
	}
	logModel.Save()
}
