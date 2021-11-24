package validator

import (
	"github.com/go-playground/validator/v10"
	"strconv"
)

// Numeric 对自定义校验标签numeric的校验 规则:若一个字符串可以被转化为整型 则校验成功 否则校验失败
func Numeric(fl validator.FieldLevel) bool {
	// 尝试将具有numeric校验tag的字段转化为string
	if numericStr, ok := fl.Field().Interface().(string); ok {
		_, err := strconv.Atoi(numericStr)
		if err == nil {
			return true
		}
	}
	return false
}
