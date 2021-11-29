package img

import (
	business "WeddingDressManage/business/img"
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/param/request/img"
	imgResponse "WeddingDressManage/param/resps/img"
	"WeddingDressManage/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Upload(c *gin.Context)  {
	param := &img.UploadParam{}
	resp := &response.RespBody{}
	err := param.Bind(c)

	// 接收文件错误
	if err != nil {
		receiveErr := &sysError.ReceiveFileError{RealError: err}
		resp.ReceiveFileError(receiveErr)
		c.JSON(http.StatusOK, resp)
		return
	}

	// 校验
	errs := param.Validate(err)
	if errs != nil {
		resp.ValidateError(errs)
		c.JSON(http.StatusOK, resp)
		return
	}

	imgBiz := &business.Img{}
	err = imgBiz.Upload(param)
	if err != nil {
		if saveFileErr, ok := err.(*sysError.SaveFileError); ok {
			resp.SaveFileError(saveFileErr)
			c.JSON(http.StatusOK, resp)
			return
		}
	}

	respParam := &imgResponse.Response{Url: imgBiz.Url}
	resp.Success(map[string]interface{}{"url":respParam.Url})
	c.JSON(http.StatusOK, resp)
	return
}
