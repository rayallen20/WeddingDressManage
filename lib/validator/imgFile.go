package validator

import (
	"github.com/go-playground/validator/v10"
	"mime/multipart"
	"strings"
)

// imgFileType 图片文件类型
var imgFileTypes []string = []string{
	"jpg",
	"png",
	"jpeg",
}

// BytesOfImgSize 图片文件大小限制 单位:Bytes
const bytesOfImgSize = 8 * 1024 * 1024

// ImgFile 对自定义校验标签imgFile的校验 规则:
// 1. 文件类型必须为 jpg,jpeg,png其中之一
// 2. 文件大小不得超过8MB
// TODO:此校验规则未生效!
func ImgFile(fl validator.FieldLevel) bool {
	if file, ok := fl.Field().Interface().(*multipart.FileHeader); ok {
		if !isImg(file) {
			return false
		}

		if !sizeIsOverLimit(file) {
			return false
		}

		return true
	}
	return false
}

func isImg(file *multipart.FileHeader) bool {
	fileName := file.Filename
	fileNameSegments := strings.Split(fileName, ".")
	fileType := fileNameSegments[len(fileNameSegments) - 1]
	for _, imgFileType := range imgFileTypes {
		if fileType == imgFileType {
			return true
		}
	}
	return false
}

func sizeIsOverLimit(file *multipart.FileHeader) bool {
	return file.Size <= bytesOfImgSize
}
