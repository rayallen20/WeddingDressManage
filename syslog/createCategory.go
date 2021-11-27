package syslog

import (
	"WeddingDressManage/model"
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type CreateCategory struct {
	data string
	TargetId int
}

func (c *CreateCategory) GetData(ginContext *gin.Context)  {
	// 复制一份请求体 用作后续记录日志
	// ioutil.ReadAll()会将c.Request.body的内容直接提取出来 而非复制一份
	// 所以后续还要把提取出来的内容还原到c.Request.body上
	bodyBytes, _ := ioutil.ReadAll(ginContext.Request.Body)
	c.data = string(bodyBytes)
	ginContext.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
}

func (c *CreateCategory) Logger()  {
	logModel := &model.OperationLog{
		Kind:        model.OperationType["createCategoryAndDress"],
		TargetId:    c.TargetId,
		Data:        c.data,
	}
	logModel.Save()
}