package user

import (
	"WeddingDressManage/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	res := map[string]interface{}{
		"status":           "ok",
		"type":             "pc",
		"currentAuthority": "admin",
	}

	resp := &response.RespBody{}
	resp.Success(res)
	c.JSON(http.StatusOK, resp)
	return
}
