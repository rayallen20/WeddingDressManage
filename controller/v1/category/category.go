package category

import (
	"WeddingDressManage/business/v1/dress"
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/param/v1/dress/category/request"
	"WeddingDressManage/response"
	"WeddingDressManage/syslog"
	"github.com/gin-gonic/gin"
	"io/ioutil"
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
	param.ExtractUri()

	category := &dress.Category{}
	err = category.Add(param)

	// 数据库错误
	if dbError, ok := err.(*sysError.DbError); ok {
		resp.DbError(dbError)
		c.JSON(http.StatusOK, resp)
		return
	}

	// kind不存在错误
	if kindNotExistError, ok := err.(*sysError.KindNotExistError); ok {
		resp.KindNotExistError(kindNotExistError)
		c.JSON(http.StatusOK, resp)
		return
	}

	// category已存在错误
	if categoryHasExistError, ok := err.(*sysError.CategoryHasExistError); ok {
		resp.CategoryHasExistError(categoryHasExistError)
		c.JSON(http.StatusOK, resp)
		return
	}

	// 记录日志
	bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
	body := string(bodyBytes[:])
	log := &syslog.CreateCategory{
		Data:     body,
		TargetId: category.Id,
	}
	log.Logger()

	resp.Success(map[string]interface{}{})
	c.JSON(http.StatusOK, resp)
	return
}