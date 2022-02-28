package validator

import (
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/param/request/requestiface"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io"
)

// Bind 将请求参数反序列化为结构体的实例 返回的错误若没有明确为绑定错误 则请求参数已经绑定到结构体的实例上了
func Bind(param requestiface.RequestParam, allParams []interface{}, c *gin.Context) error {
	err := c.ShouldBindJSON(param)
	if err != nil {
		// 请求中的JSON为空
		if errors.As(err, &io.EOF) {
			nilJsonErr := &sysError.RequestNilJsonError{}
			return nilJsonErr
		}

		// nil结构体 属于后端GO编码内部错误
		var invalidUnmarshalError *json.InvalidUnmarshalError
		if errors.As(err, &invalidUnmarshalError) {
			invalidUnmarshalErr, _ := err.(*json.InvalidUnmarshalError)
			sysInvalidUnmarshalErr := &sysError.InvalidUnmarshalError{
				Type: invalidUnmarshalErr.Type,
			}
			return sysInvalidUnmarshalErr
		}

		// 字段类型错误
		var unmarshalTypeErr *json.UnmarshalTypeError
		if errors.As(err, &unmarshalTypeErr) {
			unmarshalTypeError, _ := err.(*json.UnmarshalTypeError)
			sysUnmarshalTypeError := &sysError.UnmarshalTypeError{
				Value:  unmarshalTypeError.Value,
				Type:   unmarshalTypeError.Type,
				Struct: unmarshalTypeError.Struct,
				Field:  unmarshalTypeError.Field,
			}
			sysUnmarshalTypeError.SetMsg(allParams)
			return sysUnmarshalTypeError
		}

		// 1. nil结构体 *json.InvalidUnmarshalError
		if invalidUnmarshalErr, ok := err.(*json.InvalidUnmarshalError); ok {
			sysInvalidUnmarshalErr := &sysError.InvalidUnmarshalError{
				Type: invalidUnmarshalErr.Type,
			}
			return sysInvalidUnmarshalErr
		}
	}
	return err
}
