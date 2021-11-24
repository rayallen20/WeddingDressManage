package validator

import (
	"WeddingDressManage/conf"
	"WeddingDressManage/lib/helper/urlHelper"
	"github.com/go-playground/validator/v10"
)

// ImgUrl 对自定义校验标签imgUrl的校验 规则:若一个语义为网址的字符串中的url部分为图片服务器url 则校验成功 否则校验失败
func ImgUrl(fl validator.FieldLevel) bool {
	if webSite, ok := fl.Field().Interface().(string); ok {
		url := urlHelper.GetUrlFromWebSite(webSite)
		if url == conf.Conf.File.DomainName {
			return true
		}
	}

	return false
}
