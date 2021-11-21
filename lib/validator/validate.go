package validator

import (
	"WeddingDressManage/lib/sysError"
	"github.com/go-playground/validator/v10"
	"strings"
)

// Validate 检查已经完成绑定操作的错误 若完成绑定后存在校验错误 则返回系统校验错误
func Validate(err error) []*sysError.ValidateError {
	if err == nil {
		return nil
	}

	if errs, ok := err.(validator.ValidationErrors); ok {
		validateErrors := generateValidateErrors(errs.Translate(trans))
		return validateErrors
	}
	return nil
}

// generateValidateErrors 根据校验失败错误 创建系统校验失败错误
func generateValidateErrors(validateFailMap map[string]string) []*sysError.ValidateError {
	validateErrors := make([]*sysError.ValidateError, 0, len(validateFailMap))

	for failField, failDetail := range validateFailMap {
		// 获取字段对应的errField值
		failKey := ""
		failFieldByLevel := strings.Split(failField, ".")
		// Tips:此处由于参数均为2层 所以不会出现分割错误字段后仅有1个元素的情况
		for i := 1; i < len(failFieldByLevel); i++ {
			failKey += failFieldByLevel[i]
			if i != len(failFieldByLevel) - 1 {
				failKey += "."
			}
		}

		failDetailWords := strings.Split(failDetail, " ")
		failMsg := ""
		for j := 1; j < len(failDetailWords); j++ {
			failMsg += failDetailWords[j]
			if j != len(failDetailWords) - 1 {
				failMsg += " "
			}
		}

		validateError := &sysError.ValidateError{
			Key: failKey,
			Msg: failMsg,
		}
		validateErrors = append(validateErrors, validateError)
	}

	return validateErrors
}
