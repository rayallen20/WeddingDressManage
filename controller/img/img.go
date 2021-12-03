package img

import (
	business "WeddingDressManage/business/img"
	"WeddingDressManage/controller"
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/param/request/img"
	imgResponse "WeddingDressManage/param/resps/img"
	"WeddingDressManage/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Upload(c *gin.Context)  {
	var param *img.UploadParam = &img.UploadParam{}
	resp := controller.CheckParam(param, c, nil)
	if resp != nil {
		c.JSON(http.StatusOK, resp)
	}
	resp = &response.RespBody{}

	imgBiz := &business.Img{}
	err := imgBiz.Upload(param)
	if err != nil {
		if saveFileErr, ok := err.(*sysError.SaveFileError); ok {
			resp.SaveFileError(saveFileErr)
			c.JSON(http.StatusOK, resp)
			return
		}
	}

	respParam := &imgResponse.UploadResponse{Url: imgBiz.Url}
	resp.Success(map[string]interface{}{"url":respParam.Url})
	c.JSON(http.StatusOK, resp)
	return
}
