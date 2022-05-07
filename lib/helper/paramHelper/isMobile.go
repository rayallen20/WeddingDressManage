package paramHelper

import (
	"regexp"
)

func IsMobile(mobile string) bool {
	// 匹配规则
	// ^1:				第一位为一
	// [345789]{1}:		后接一位345789 的数字
	// \\d:				\d的转义 表示数字 {9} 接9位
	// $:				结束符
	regRuler := "^1[345789]{1}\\d{9}$"
	reg := regexp.MustCompile(regRuler)
	return reg.MatchString(mobile)
}
