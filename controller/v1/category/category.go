package category

import (
	"WeddingDressManage/business/v1/dress"
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/param/v1/dress/category/request"
	"WeddingDressManage/param/v1/dress/category/resps"
	"WeddingDressManage/response"
	"WeddingDressManage/syslog"
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

// Add 添加新品类礼服
func Add(c *gin.Context) {
	var param *request.AddParam = &request.AddParam{}
	var resp *response.RespBody = &response.RespBody{}

	// 复制一份请求体 用作后续记录日志
	// ioutil.ReadAll()会将c.Request.body的内容直接提取出来 而非复制一份
	// 所以后续还要把提取出来的内容还原到c.Request.body上
	bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

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
	log := &syslog.CreateCategory{
		Data:     string(bodyBytes),
		TargetId: category.Id,
	}
	log.Logger()

	resp.Success(map[string]interface{}{})
	c.JSON(http.StatusOK, resp)
	return
}

// Show 礼服品类展示
func Show(c *gin.Context) {
	var param *request.ShowParam = &request.ShowParam{}
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

	category := &dress.Category{}
	categories, err := category.Show(param)
	if err != nil {
		if dbError, ok := err.(*sysError.DbError); ok {
			resp.DbError(dbError)
			c.JSON(http.StatusOK, resp)
			return
		}
	}

	respParam := &resps.CategoryResponse{}
	respParams := respParam.Generate(categories)
	resp.Success(map[string]interface{}{
		"categories":respParams,
	})
	c.JSON(http.StatusOK, resp)
	return
}