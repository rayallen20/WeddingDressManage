package urlHelper

import (
	"WeddingDressManage/conf"
	"strings"
)

// GetUriFromWebsite 从一个网址中截取uri部分
func GetUriFromWebsite(website string) (uri string) {
	webSiteSegments := strings.Split(website, "//")[1]
	uriSegments := strings.Split(webSiteSegments, "/")[1:]

	uri = "/"
	for k, v := range  uriSegments {
		uri += v
		if k != len(uriSegments) - 1 {
			uri += "/"
		}
	}
	return uri
}


// GetUrlFromWebSite 从一个网址中截取url部分
func GetUrlFromWebSite(website string) (url string) {
	webSiteSegments := strings.Split(website, "//")
	url = strings.Split(webSiteSegments[1], "/")[0]
	return url
}

// GenFullImgWebSite 根据uri生成一个完整图片地址
func GenFullImgWebSite(uri string) string {
	return conf.Conf.File.Protocol + conf.Conf.File.DomainName + conf.Conf.File.ImgUri + uri
}

// GenFullImgWebSites 根据给定uri集合生成一个完整图片地址集合
func GenFullImgWebSites(uris []string) []string {
	webSites := make([]string, 0, len(uris))

	for _, uri := range uris {
		webSite := conf.Conf.File.Protocol + conf.Conf.File.DomainName + conf.Conf.File.ImgUri + uri
		webSites = append(webSites, webSite)
	}
	return webSites
}