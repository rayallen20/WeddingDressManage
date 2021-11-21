package validator

import (
	"github.com/go-playground/validator/v10"
	"strconv"
)

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
