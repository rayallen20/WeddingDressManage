package img

import (
	file2 "WeddingDressManage/business/v1/file"
	"WeddingDressManage/lib/response"
	"WeddingDressManage/lib/wdmError"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Upload(c *gin.Context) {
	resp := &response.ResBody{}
	file, err := c.FormFile("file")
	if err != nil {
		resp.ReceiveFileError(err, map[string]interface{}{})
		c.JSON(http.StatusOK, resp)
		return
	}

	img := &file2.Img{}
	uri, err := img.Upload(file)
	if err != nil {
		if _, ok := err.(wdmError.FileTypeError); ok {
			resp.FileTypeError(err, map[string]interface{}{})
			c.JSON(http.StatusOK, resp)
			return
		} else if _, ok1 := err.(wdmError.SaveFileError); ok1 {
			resp.SaveFileError(err, map[string]interface{}{})
			c.JSON(http.StatusOK, resp)
			return
		}
	}

	data := map[string]interface{}{
		"uri":uri,
	}
	resp.Success(data)
	c.JSON(http.StatusOK, resp)
	return
}
