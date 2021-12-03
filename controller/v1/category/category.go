package category

import (
	"WeddingDressManage/business/v1/dress"
	"WeddingDressManage/controller"
	"WeddingDressManage/lib/sysError"
	categoryRequest "WeddingDressManage/param/request/v1/category"
	categoryResponse "WeddingDressManage/param/resps/v1/category"
	"WeddingDressManage/param/resps/v1/pagination"
	"WeddingDressManage/response"
	"WeddingDressManage/syslog"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Add 添加新品类礼服
func Add(c *gin.Context) {
	var param *categoryRequest.AddParam = &categoryRequest.AddParam{}
	var logger *syslog.CreateCategory = &syslog.CreateCategory{}

	resp := controller.CheckParam(param, c, logger)
	if resp != nil {
		c.JSON(http.StatusOK, resp)
		return
	}

	resp = &response.RespBody{}
	param.ExtractUri()

	category := &dress.Category{}
	err := category.Add(param)

	if err != nil {
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
	resp := controller.CheckParam(param, c, nil)
	if resp != nil {
		c.JSON(http.StatusOK, resp)
		return
	}

	resp = &response.RespBody{}
	category := &dress.Category{}
	categories, totalPage, err := category.Show(param)
	if err != nil {
		if dbError, ok := err.(*sysError.DbError); ok {
			resp.DbError(dbError)
			c.JSON(http.StatusOK, resp)
			return
		}
	}

	categoryParam := &categoryResponse.ShowResponse{}
	categoryParams := categoryParam.Generate(categories)
	paginationParam := &pagination.Response{
		CurrentPage: param.Pagination.CurrentPage,
		ItemPerPage: param.Pagination.ItemPerPage,
		TotalPage:   totalPage,
	}
	resp.Success(map[string]interface{}{
		"categories":categoryParams,
		"pagination":paginationParam,
	})
	c.JSON(http.StatusOK, resp)
	return
}

// ShowOne 展示1条品类信息
func ShowOne(c *gin.Context) {
	var param *categoryRequest.ShowOneParam = &categoryRequest.ShowOneParam{}
	resp := controller.CheckParam(param, c, nil)
	if resp != nil {
		c.JSON(http.StatusOK, resp)
		return
	}

	resp = &response.RespBody{}
	categoryBiz := &dress.Category{}
	err := categoryBiz.ShowOne(param)

	if err != nil {
		if dbErr, ok := err.(*sysError.DbError); ok {
			resp.DbError(dbErr)
			c.JSON(http.StatusOK, resp)
			return
		}

		if notExistErr, ok := err.(*sysError.CategoryNotExistError); ok {
			resp.CategoryNotExistError(notExistErr)
			c.JSON(http.StatusOK, resp)
			return
		}
	}

	respParam := &categoryResponse.ShowOneResponse{}
	respParam.Fill(categoryBiz)
	data := map[string]interface{}{
		"category":respParam,
	}

	resp.Success(data)
	c.JSON(http.StatusOK, resp)
	return
}

// Update 修改品类信息
func Update(c *gin.Context) {
	var param *categoryRequest.UpdateParam = &categoryRequest.UpdateParam{}
	var logger *syslog.UpdateCategory = &syslog.UpdateCategory{}

	resp := controller.CheckParam(param, c, logger)
	if resp != nil {
		c.JSON(http.StatusOK, resp)
		return
	}

	resp = &response.RespBody{}
	param.ExtractUri()

	categoryBiz := &dress.Category{}
	err := categoryBiz.Update(param)
	if err != nil {
		if dbErr, ok := err.(*sysError.DbError); ok {
			resp.DbError(dbErr)
			c.JSON(http.StatusOK, resp)
			return
		}

		if notExistErr, ok := err.(*sysError.CategoryNotExistError); ok {
			resp.CategoryNotExistError(notExistErr)
			c.JSON(http.StatusOK, resp)
			return
		}

		if dbErr, ok := err.(*sysError.DbError); ok {
			resp.DbError(dbErr)
			c.JSON(http.StatusOK, resp)
			return
		}
	}

	logger.TargetId = categoryBiz.Id
	logger.Logger()

	resp.Success(map[string]interface{}{})
	c.JSON(http.StatusOK, resp)
	return
}