package dress

import (
	"WeddingDressManage/business/v1/dress"
	"WeddingDressManage/lib/sysError"
	dressRequest "WeddingDressManage/param/request/v1/dress"
	dressResponse "WeddingDressManage/param/resps/v1/dress"
	"WeddingDressManage/param/resps/v1/pagination"
	"WeddingDressManage/response"
	"WeddingDressManage/syslog"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Add 在已有品类下添加礼服
func Add(c *gin.Context) {
	var param *dressRequest.AddParam = &dressRequest.AddParam{}
	var resp *response.RespBody = &response.RespBody{}

	// 记录请求参数
	logger := &syslog.AddDress{}
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

	dressBiz := &dress.Dress{}
	dresses, err := dressBiz.Add(param)

	// 数据库错误
	if dbError, ok := err.(*sysError.DbError); ok {
		resp.DbError(dbError)
		c.JSON(http.StatusOK, resp)
		return
	}

	// 品类信息不存在错误
	if categoryNotExistError, ok := err.(*sysError.CategoryNotExistError); ok {
		resp.CategoryNotExistError(categoryNotExistError)
		c.JSON(http.StatusOK, resp)
		return
	}

	// 记录日志
	dressIds := make([]int, 0, len(dresses))

	for _, dress := range dresses {
		dressIds = append(dressIds, dress.Id)
	}

	logger.TargetIds = dressIds
	logger.Logger()

	resp.Success(map[string]interface{}{})
	c.JSON(http.StatusOK, resp)
}

// ShowUsable 展示指定品类下的可用礼服信息
func ShowUsable(c *gin.Context) {
	param := &dressRequest.ShowUsableParam{}
	resp := &response.RespBody{}

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

	dressBiz := &dress.Dress{}
	categoryBiz, dressBizs, totalPage, err := dressBiz.ShowUsable(param)
	if dbErr, ok := err.(*sysError.DbError); ok {
		resp.DbError(dbErr)
	}

	paginationResp := &pagination.Response{
		CurrentPage: param.Pagination.CurrentPage,
		ItemPerPage: param.Pagination.ItemPerPage,
		TotalPage:   totalPage,
	}

	respParam := &dressResponse.ShowUsableResponse{}
	respParam.Fill(categoryBiz, dressBizs, paginationResp)
	data := map[string]interface{}{
		"data":respParam,
	}
	resp.Success(data)
	c.JSON(http.StatusOK, resp)
	return
}