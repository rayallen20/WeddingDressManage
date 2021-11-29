package kind

import (
	"WeddingDressManage/business/v1/dress"
	"WeddingDressManage/lib/sysError"
	responseParam "WeddingDressManage/param/resps/v1/kind"
	"WeddingDressManage/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Show 礼服大类展示Action
func Show(c *gin.Context)  {
	resp := &response.RespBody{}

	kind := &dress.Kind{}
	kinds, err := kind.Show()
	if dbErr, ok := err.(*sysError.DbError); ok {
		resp.DbError(dbErr)
		c.JSON(http.StatusOK, resp)
		return
	}

	respParam := &responseParam.Response{}
	respParams := respParam.Generate(kinds)

	data := map[string]interface{}{
		"kinds": respParams,
	}

	resp.Success(data)
	c.JSON(http.StatusOK, resp)
	return
}
