package kind

import (
	"WeddingDressManage/business/v1/dress/kind"
	"WeddingDressManage/lib/response"
	"WeddingDressManage/lib/wdmError"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Show(c *gin.Context) {
	kind := &kind.Kind{}
	kinds, err := kind.Show()

	resp := &response.ResBody{}
	if err != nil {
		if _, ok := err.(wdmError.DBError);ok {
			// DB错误
			resp.DBError(err, map[string]interface{}{})
			c.JSON(http.StatusOK, resp)
			return
		}
	}

	data := map[string]interface{}{
		"kinds":kinds,
	}
	resp.Success(data)
	c.JSON(http.StatusOK, resp)
	return
}