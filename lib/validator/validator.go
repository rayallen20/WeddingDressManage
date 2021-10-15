package validator

import (
	"WeddingDressManage/lib/wdmError"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strconv"
	"strings"
)

var trans ut.Translator

// InitTrans 初始化全局错误翻译器
func InitTrans(locale string) (err error) {
	if trans != nil {
		return
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		// 获取自定义tag:errField的方法
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("errField"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		enT := en.New()
		uni := ut.New(enT, enT)

		var ok bool
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		// 注册翻译器
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = zhTranslations.RegisterDefaultTranslations(v, trans)
		}
		return
	}

	return
}

// removeField 从错误信息中移除字段信息 只保留报错信息
func removeField(errs map[string]string) []string {
	res := make([]string, 0, len(errs))
	for _, err := range errs {
		res = append(res, err)
	}
	return res
}

// GenerateErrsInfo 生成数据校验错误信息
func GenerateErrsInfo(err error) (errInfos []string, ok bool) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		// 表示错误类型转化失败的信息
		return nil, ok
	}

	return removeField(errs.Translate(trans)), ok
}

// ValidateParam 对请求参数做校验
func ValidateParam(param interface{}, c *gin.Context) (err error) {
	err = c.ShouldBindJSON(param)
	if err != nil {
		if UnmarshalTypeErr, ok := err.(*json.UnmarshalTypeError); ok {
			paramTypeErr := &wdmError.ParamTypeError {
				Message: err.Error(),
				StructFieldName: UnmarshalTypeErr.Field,
			}
			paramTypeErr.GetFormFieldAndShouldType(param)
			return paramTypeErr
		}

		errInfos, ok := GenerateErrsInfo(err)
		if !ok {
			err = wdmError.BindingValidatorError{Message: err.Error()}
			return
		} else {
			err = wdmError.ParamValueError {
				Message: "",
				Details: errInfos,
			}
			return
		}
	}

	nonNumericFieldsName := getNonNumericFieldsName(param)
	if nonNumericFieldsName != nil {
		err = wdmError.NumericStringError {
			Message: "",
			NotNumericFields: nonNumericFieldsName,
		}
		return
	}

	return nil
}

// getNonNumericFieldsName 获取请求参数中 所有要求字面量为数字的string类型 但不符合要求的字段名
func getNonNumericFieldsName(param interface{}) []string {
	var nonNumericFieldsName []string
	structInfo := reflect.TypeOf(param).Elem()
	valueInfo := reflect.ValueOf(param).Elem()
	for i := 0; i < structInfo.NumField(); i++ {
		numericTagContent := structInfo.Field(i).Tag.Get("numeric")
		if numericTagContent == "true" {
			numericFieldName := structInfo.Field(i).Name
			numericFieldValue := valueInfo.FieldByName(numericFieldName).String()
			if _, err := strconv.Atoi(numericFieldValue); err != nil {
				nonNumericFieldsName = append(nonNumericFieldsName, numericFieldName)
			}
		}
	}
	return nonNumericFieldsName
}
