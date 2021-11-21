package validator

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

var trans ut.Translator

// 默认将错误翻译为英文
var locale string = "en"

// init 初始化全局错误翻译器
func init() {
	if trans != nil {
		return
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		// 获取自定义的tag errField的方法
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("errField"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		// 注册自定义校验tag与校验函数的映射
		err := v.RegisterValidation("numeric", Numeric)
		if err != nil {
			panic("register custom validation failed")
		}

		enT := en.New()
		uni := ut.New(enT, enT)

		translator, translatorOk := uni.GetTranslator(locale)
		if !translatorOk {
			panic(fmt.Sprintf("init translator for validateor failed: uni.GetTranslator(%s) failed", locale))
		}

		trans = translator

		// 注册翻译器
		switch locale {
		case "en":
			err := enTranslations.RegisterDefaultTranslations(v, trans)
			if err != nil {
				panic("init translator for validator failed:" + err.Error())
			}
		default:
			err := zhTranslations.RegisterDefaultTranslations(v, trans)
			if err != nil {
				panic("init translator for validator failed:" + err.Error())
			}
		}
		return
	}

	panic("init translator for validator failed:can not convert Validator.Engine to *Validate")
}