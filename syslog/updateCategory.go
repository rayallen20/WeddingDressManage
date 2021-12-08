package syslog

import (
	"WeddingDressManage/model"
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type UpdateCategory struct {
	data     string
	TargetId int
}

func (u *UpdateCategory) GetData(c *gin.Context) {
	bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
	u.data = string(bodyBytes)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
}

func (u *UpdateCategory) Logger() {
	logModel := &model.OperationLog{
		Kind:     model.OperationType["updateCategory"],
		TargetId: u.TargetId,
		Data:     u.data,
	}
	logModel.Save()
}
