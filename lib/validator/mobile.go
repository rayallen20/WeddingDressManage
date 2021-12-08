package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

// Mobile 对自定义校验标签mobile的校验 规则:
// 1. 一个字符串以13/14/15/17/18开头
// 2. 字符串长度为11
// 3. 字符串中的每一个字符的字面量均为数字
func Mobile(fl validator.FieldLevel) bool {
	if mobileStr, ok := fl.Field().Interface().(string); ok {
		result, _ := regexp.MatchString(`^(1[3|4|5|7|8][0-9]\d{4,8})$`, mobileStr)
		if result {
			return true
		}
	}
	return false
}
