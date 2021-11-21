package category

import (
	"WeddingDressManage/lib/param/v1/dress/category/request"
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Add(c *gin.Context) {
	var param *request.AddParam = &request.AddParam{}
	var resp *response.RespBody = &response.RespBody{}
	err := param.Bind(c)

	if invalidUnmarshalError, ok := err.(*sysError.InvalidUnmarshalError); ok {
		resp.InvalidUnmarshalError(invalidUnmarshalError)
		c.JSON(http.StatusOK, resp)
		return
	}

	if unmarshalTypeError, ok := err.(*sysError.UnmarshalTypeError); ok {
		resp.FieldTypeError(unmarshalTypeError)
		c.JSON(http.StatusOK, resp)
		return
	}

	validateErrors := param.Validate(err)
	if validateErrors != nil {
		resp.ValidateError(validateErrors)
		c.JSON(http.StatusOK, resp)
		return
	}
}