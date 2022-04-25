package img

import (
	"WeddingDressManage/lib/sysError"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"mime/multipart"
	"strings"
)

// imgFileTypes 图片文件类型
var imgFileTypes []string = []string{
	"jpg",
	"png",
	"jpeg",
}

// BytesOfImgSize 图片文件大小限制 单位:Bytes
const bytesOfImgSize = 8 * 1024 * 1024

type UploadParam struct {
	// TODO:此处对自定义标签imgFile定义了校验规则 但校验函数没有被调用 为什么?
	File *multipart.FileHeader `form:"img" binding:"imgFile,required" errField:"img"`
}

func (u *UploadParam) Bind(c *gin.Context) error {
	return c.ShouldBindWith(u, binding.FormMultipart)
}

// Validate TODO:此处由于自定义标签校验规则没有被调用 故又在参数类下实现了校验方法 由controller层调用
// 校验规则:
// 1. 文件类型必须为 jpg,jpeg,png其中之一
// 2. 文件大小不得超过8MB
func (u *UploadParam) Validate(err error) []*sysError.ValidateError {
	// 此处cap定义为2 因为校验只有2种错误
	errs := make([]*sysError.ValidateError, 0, 2)
	if !u.isImg() {
		validateErr := &sysError.ValidateError{
			Key: "img",
			Msg: "file isn't a image file",
		}
		errs = append(errs, validateErr)
	}

	if !u.sizeIsOverLimit() {
		validateErr := &sysError.ValidateError{
			Key: "img",
			Msg: "size must less than 8 MB",
		}

		errs = append(errs, validateErr)
	}

	if len(errs) != 0 {
		return errs
	}
	return nil
}

// isImg 判断上传文件是否为图片
func (u *UploadParam) isImg() bool {
	fileName := u.File.Filename
	fileNameSegments := strings.Split(fileName, ".")
	fileType := fileNameSegments[len(fileNameSegments)-1]
	for _, imgFileType := range imgFileTypes {
		if fileType == imgFileType {
			return true
		}
	}
	return false
}

// sizeIsOverLimit 判断上传文件大小是否超过限制
func (u *UploadParam) sizeIsOverLimit() bool {
	return u.File.Size <= bytesOfImgSize
}
