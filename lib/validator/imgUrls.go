package validator

import (
	"WeddingDressManage/conf"
	"WeddingDressManage/lib/helper/urlHelper"
	"github.com/go-playground/validator/v10"
)

// ImgUrls 对自定义校验标签imgUrl的校验 规则:若一个包含多个网址的字符串切片中 每个网址的url均为图片服务器url 则校验通过 否则校验失败
func ImgUrls(fl validator.FieldLevel) bool {
	if webSites, ok := fl.Field().Interface().([]string); ok {
		for _, webSite := range webSites {
			url := urlHelper.GetUrlFromWebSite(webSite)
			if url != conf.Conf.File.DomainName {
				return false
			}
		}
		return true
	}
	return false
}
