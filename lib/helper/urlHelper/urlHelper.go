package urlHelper

import (
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