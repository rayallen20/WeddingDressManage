package validator

import (
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/param/v1/request/requestiface"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

// Bind 将请求参数反序列化为结构体的实例 返回的错误若没有明确为绑定错误 则请求参数已经绑定到结构体的实例上了
func Bind(param requestiface.RequestParam, allParams []interface{}, c *gin.Context) error {
	err := c.ShouldBindJSON(param)
	if err != nil {
		// 1. nil结构体 *json.InvalidUnmarshalError
		if invalidUnmarshalErr, ok := err.(*json.InvalidUnmarshalError); ok {
			sysInvalidUnmarshalErr := &sysError.InvalidUnmarshalError{
				Type: invalidUnmarshalErr.Type,
			}
			return sysInvalidUnmarshalErr
		}
		// 2. 字段类型错误 *json.UnmarshalTypeError
		if unmarshalTypeError, ok := err.(*json.UnmarshalTypeError); ok {
			sysUnmarshalTypeError := &sysError.UnmarshalTypeError{
				Value:  unmarshalTypeError.Value,
				Type:   unmarshalTypeError.Type,
				Struct: unmarshalTypeError.Struct,
				Field:  unmarshalTypeError.Field,
			}

			sysUnmarshalTypeError.SetMsg(allParams)
			return sysUnmarshalTypeError
		}
	}
	return err
}
