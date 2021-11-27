package category

import (
	"WeddingDressManage/business/v1/dress"
	"WeddingDressManage/lib/sysError"
	categoryRequest "WeddingDressManage/param/v1/request/category"
	categoryResponse "WeddingDressManage/param/v1/resps/category"
	"WeddingDressManage/param/v1/resps/pagination"
	"WeddingDressManage/response"
	"WeddingDressManage/syslog"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Add 添加新品类礼服
func Add(c *gin.Context) {
	var param *categoryRequest.AddParam = &categoryRequest.AddParam{}
	var resp *response.RespBody = &response.RespBody{}

	// 记录请求参数
	logger := &syslog.CreateCategory{}
	logger.GetData(c)

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
	logger.TargetId = category.Id
	logger.Logger()

	resp.Success(map[string]interface{}{})
	c.JSON(http.StatusOK, resp)
	return
}

// Show 礼服品类展示
func Show(c *gin.Context) {
	var param *categoryRequest.ShowParam = &categoryRequest.ShowParam{}
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
	categories, totalPage, err := category.Show(param)
	if err != nil {
		if dbError, ok := err.(*sysError.DbError); ok {
			resp.DbError(dbError)
			c.JSON(http.StatusOK, resp)
			return
		}
	}

	categoryParam := &categoryResponse.Response{}
	categoryParams := categoryParam.Generate(categories)
	paginationParam := &pagination.Response{
		CurrentPage: param.Pagination.CurrentPage,
		ItemCounter: param.Pagination.ItemPerPage,
		TotalPage:   totalPage,
	}
	resp.Success(map[string]interface{}{
		"categories":categoryParams,
		"pagination":paginationParam,
	})
	c.JSON(http.StatusOK, resp)
	return
}