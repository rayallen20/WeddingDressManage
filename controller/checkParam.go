package controller

import (
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/param/request/requestiface"
	"WeddingDressManage/response"
	"WeddingDressManage/syslog/logInterface"
	"github.com/gin-gonic/gin"
	"time"
)

// CheckParam 从请求中获取参数信息并绑定至给定参数对象 若logger不为空 则记录请求参数
// 若对参数的绑定与校验均正确 则返回空 否则返回一个响应体
func CheckParam(param requestiface.RequestParam, c *gin.Context, logger logInterface.SysLog) *response.RespBody {
	// 若logger不为空 则记录请求参数
	if logger != nil {
		recordData(c, logger)
	}

	err := param.Bind(c)

	var resp *response.RespBody = &response.RespBody{}

	if requestNilJsonErr, ok := err.(*sysError.RequestNilJsonError); ok {
		resp.RequestNilJsonError(requestNilJsonErr)
		return resp
	}

	if invalidUnmarshalError, ok := err.(*sysError.InvalidUnmarshalError); ok {
		resp.InvalidUnmarshalError(invalidUnmarshalError)
		return resp
	}

	if unmarshalTypeError, ok := err.(*sysError.UnmarshalTypeError); ok {
		resp.FieldTypeError(unmarshalTypeError)
		return resp
	}

	if timeParseError, ok := err.(*time.ParseError); ok {
		resp.TimeParseError(timeParseError)
		return resp
	}

	validateErrors := param.Validate(err)
	if validateErrors != nil {
		resp.ValidateError(validateErrors)
		return resp
	}
	return nil
}

// recordData 记录请求参数
func recordData(c *gin.Context, logger logInterface.SysLog) {
	logger.GetData(c)
}
