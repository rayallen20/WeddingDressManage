package file

import (
	"WeddingDressManage/business/v1/img"
	"WeddingDressManage/conf"
	"WeddingDressManage/lib/response"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"strings"
)

func UploadImg(c *gin.Context) {
	errResp, files := validateUploadParam(c)
	if errResp != nil {
		c.JSON(http.StatusOK, errResp)
		return
	}

	for _, file := range files {
		if !isImg(file.Filename) {
			// 文件非图片错误
			c.JSON(http.StatusOK, response.FileIsNotImgResp([]interface{}{}))
			return
		}
	}

	imgObj := &img.Img{}
	err := imgObj.Upload(files)
	if err != nil {
		c.JSON(http.StatusOK, response.UploadFileFailedResp(err, []interface{}{}))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResp([]interface{}{}))
	return
}

func isImg(fileName string) bool {
	splitFileName := strings.Split(fileName, ".")
	for _, imgType := range conf.Conf.File.ImgType {
		if splitFileName[len(splitFileName) - 1] == imgType {
			return true
		}
	}
	return false
}

func validateUploadParam(c *gin.Context) (gin.H, []*multipart.FileHeader) {
	form, err := c.MultipartForm()
	// 非multipart/form-data类型表单
	if err != nil {
		return response.RequestContentTypeErrResp(err.Error(), []interface{}{}), nil

	}

	files := form.File["file"]
	// 没有文件
	if len(files) == 0 {
		return response.FileNumZeroResp("file", []interface{}{}), nil
	}

	return nil, files
}
