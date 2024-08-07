package dress

import (
	"WeddingDressManage/business/v1/dress"
	"WeddingDressManage/controller"
	"WeddingDressManage/lib/sysError"
	dressRequest "WeddingDressManage/param/request/v1/dress"
	dressResponse "WeddingDressManage/param/resps/v1/dress"
	"WeddingDressManage/param/resps/v1/pagination"
	"WeddingDressManage/response"
	"WeddingDressManage/syslog"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Add 在已有品类下添加礼服
func Add(c *gin.Context) {
	var param *dressRequest.AddParam = &dressRequest.AddParam{}
	var logger *syslog.AddDress = &syslog.AddDress{}

	resp := controller.CheckParam(param, c, logger)
	if resp != nil {
		c.JSON(http.StatusOK, resp)
		return
	}

	resp = &response.RespBody{}
	param.ExtractUri()

	dressBiz := &dress.Dress{}
	dressBizs, err := dressBiz.Add(param)

	if err != nil {
		// 数据库错误
		var dbError *sysError.DbError
		if errors.As(err, &dbError) {
			resp.DbError(dbError)
			c.JSON(http.StatusOK, resp)
			return
		}

		// 品类信息不存在错误
		var categoryNotExistError *sysError.CategoryNotExistError
		if errors.As(err, &categoryNotExistError) {
			resp.CategoryNotExistError(categoryNotExistError)
			c.JSON(http.StatusOK, resp)
			return
		}
	}

	// 记录日志
	dressIds := make([]int, 0, len(dressBizs))

	// Tips:此处由于之前已经使用过dressBiz 这一变量名 故命名为dressObj是无奈之举
	for _, dressObj := range dressBizs {
		dressIds = append(dressIds, dressObj.Id)
	}

	logger.TargetIds = dressIds
	logger.Logger()

	resp.Success(map[string]interface{}{})
	c.JSON(http.StatusOK, resp)
}

// ShowUsable 展示指定品类下的可用礼服信息
// TODO: 该方法是否放在品类下比较合适?
func ShowUsable(c *gin.Context) {
	var param *dressRequest.ShowUsableParam = &dressRequest.ShowUsableParam{}
	resp := controller.CheckParam(param, c, nil)
	if resp != nil {
		c.JSON(http.StatusOK, resp)
		return
	}

	resp = &response.RespBody{}

	dressBiz := &dress.Dress{}
	categoryBiz, dressBizs, totalPage, count, err := dressBiz.ShowUsable(param)
	if err != nil {
		var dbErr *sysError.DbError
		if errors.As(err, &dbErr) {
			resp.DbError(dbErr)
			c.JSON(http.StatusOK, resp)
			return
		}

		var categoryNotExistErr *sysError.CategoryNotExistError
		if errors.As(err, &categoryNotExistErr) {
			resp.CategoryNotExistError(categoryNotExistErr)
			c.JSON(http.StatusOK, resp)
			return
		}

		var dressNotExistErr *sysError.DressNotExistError
		if errors.As(err, &dressNotExistErr) {
			resp.DressNotExistError(dressNotExistErr)
			c.JSON(http.StatusOK, resp)
			return
		}
	}

	paginationResp := &pagination.Response{
		CurrentPage: param.Pagination.CurrentPage,
		ItemPerPage: param.Pagination.ItemPerPage,
		TotalPage:   totalPage,
		TotalItem:   count,
	}

	respParam := &dressResponse.ShowUsableResponse{}
	respParam.Fill(categoryBiz, dressBizs, paginationResp)
	data := map[string]interface{}{
		"category":   respParam.Category,
		"dresses":    respParam.Dresses,
		"pagination": respParam.Pagination,
	}
	resp.Success(data)
	c.JSON(http.StatusOK, resp)
	return
}

// ApplyDiscard 礼服销库申请
func ApplyDiscard(c *gin.Context) {
	var param *dressRequest.ApplyDiscardParam = &dressRequest.ApplyDiscardParam{}
	var logger *syslog.ApplyDiscardDress = &syslog.ApplyDiscardDress{}
	resp := controller.CheckParam(param, c, logger)
	if resp != nil {
		c.JSON(http.StatusOK, resp)
		return
	}

	resp = &response.RespBody{}
	dressBiz := &dress.Dress{}
	err := dressBiz.ApplyDiscard(param)

	if err != nil {
		var dbErr *sysError.DbError
		if errors.As(err, &dbErr) {
			resp.DbError(dbErr)
			c.JSON(http.StatusOK, resp)
			return
		}

		var dressNotExistErr *sysError.DressNotExistError
		if errors.As(err, &dressNotExistErr) {
			resp.DressNotExistError(dressNotExistErr)
			c.JSON(http.StatusOK, resp)
			return
		}

		var hasGiftedErr *sysError.DressHasGiftedError
		if errors.As(err, &hasGiftedErr) {
			resp.DressHasGiftedError(hasGiftedErr)
			c.JSON(http.StatusOK, resp)
			return
		}

		var hasDiscardedErr *sysError.DressHasDiscardedError
		if errors.As(err, &hasDiscardedErr) {
			resp.DressHasDiscardedError(hasDiscardedErr)
			c.JSON(http.StatusOK, resp)
			return
		}
	}

	logger.TargetId = dressBiz.Id
	logger.Logger()

	resp.Success(map[string]interface{}{})
	c.JSON(http.StatusOK, resp)
	return
}

func ApplyGift(c *gin.Context) {
	var param *dressRequest.ApplyGiftParam = &dressRequest.ApplyGiftParam{}
	var logger *syslog.ApplyGiftDress = &syslog.ApplyGiftDress{}
	resp := controller.CheckParam(param, c, logger)
	if resp != nil {
		c.JSON(http.StatusOK, resp)
		return
	}

	resp = &response.RespBody{}
	dressBiz := &dress.Dress{}
	err := dressBiz.ApplyGift(param)

	if err != nil {
		var dbErr *sysError.DbError
		if errors.As(err, &dbErr) {
			resp.DbError(dbErr)
			c.JSON(http.StatusOK, resp)
			return
		}

		var dressNotExistErr *sysError.DressNotExistError
		if errors.As(err, &dressNotExistErr) {
			resp.DressNotExistError(dressNotExistErr)
			c.JSON(http.StatusOK, resp)
			return
		}

		var hasGiftedErr *sysError.DressHasGiftedError
		if errors.As(err, &hasGiftedErr) {
			resp.DressHasGiftedError(hasGiftedErr)
			c.JSON(http.StatusOK, resp)
			return
		}

		var hasDiscardedErr *sysError.DressHasDiscardedError
		if errors.As(err, &hasDiscardedErr) {
			resp.DressHasDiscardedError(hasDiscardedErr)
			c.JSON(http.StatusOK, resp)
			return
		}

		var customerNotExistErr *sysError.CustomerNotExistError
		if errors.As(err, &customerNotExistErr) {
			resp.CustomerNotExistError(customerNotExistErr)
			c.JSON(http.StatusOK, resp)
			return
		}
	}

	logger.TargetId = dressBiz.Id
	logger.Logger()

	resp.Success(map[string]interface{}{})
	c.JSON(http.StatusOK, resp)
	return
}

func Laundry(c *gin.Context) {
	var param *dressRequest.LaundryParam = &dressRequest.LaundryParam{}
	var logger *syslog.Laundry = &syslog.Laundry{}

	resp := controller.CheckParam(param, c, logger)
	if resp != nil {
		c.JSON(http.StatusOK, resp)
		return
	}

	resp = &response.RespBody{}
	param.ExtractUri()
	dressBiz := &dress.Dress{}
	err := dressBiz.Laundry(param)
	if err != nil {
		var dbErr *sysError.DbError
		if errors.As(err, &dbErr) {
			resp.DbError(dbErr)
			c.JSON(http.StatusOK, resp)
			return
		}

		var dressNotExistErr *sysError.DressNotExistError
		if errors.As(err, &dressNotExistErr) {
			resp.DressNotExistError(dressNotExistErr)
			c.JSON(http.StatusOK, resp)
			return
		}

		var laundryStatusErr *sysError.LaundryStatusError
		if errors.As(err, &laundryStatusErr) {
			resp.LaundryStatusError(laundryStatusErr)
			c.JSON(http.StatusOK, resp)
			return
		}
	}

	logger.TargetId = dressBiz.Id
	logger.Logger()

	data := map[string]interface{}{}
	resp.Success(data)
	c.JSON(http.StatusOK, resp)
	return
}

func Maintain(c *gin.Context) {
	var param *dressRequest.MaintainParam = &dressRequest.MaintainParam{}
	var logger *syslog.Maintain = &syslog.Maintain{}

	resp := controller.CheckParam(param, c, logger)
	if resp != nil {
		c.JSON(http.StatusOK, resp)
		return
	}

	resp = &response.RespBody{}
	param.ExtractUri()

	dressBiz := &dress.Dress{}
	err := dressBiz.Maintain(param)
	if err != nil {
		var dbErr *sysError.DbError
		if errors.As(err, &dbErr) {
			resp.DbError(dbErr)
			c.JSON(http.StatusOK, resp)
			return
		}

		var dressNotExistErr *sysError.DressNotExistError
		if errors.As(err, &dressNotExistErr) {
			resp.DressNotExistError(dressNotExistErr)
			c.JSON(http.StatusOK, resp)
			return
		}

		var maintainStatusErr *sysError.MaintainStatusError
		if errors.As(err, &maintainStatusErr) {
			resp.MaintainStatusError(maintainStatusErr)
			c.JSON(http.StatusOK, resp)
			return
		}
	}

	logger.TargetId = dressBiz.Id
	logger.Logger()

	data := map[string]interface{}{}
	resp.Success(data)
	c.JSON(http.StatusOK, resp)
	return
}

func ShowOne(c *gin.Context) {
	var param *dressRequest.ShowOneParam = &dressRequest.ShowOneParam{}
	resp := controller.CheckParam(param, c, nil)
	if resp != nil {
		c.JSON(http.StatusOK, resp)
		return
	}

	resp = &response.RespBody{}
	dressBiz := &dress.Dress{}
	err := dressBiz.ShowOne(param)
	if err != nil {
		var dbErr *sysError.DbError
		if errors.As(err, &dbErr) {
			resp.DbError(dbErr)
			c.JSON(http.StatusOK, resp)
			return
		}

		var dressNotExistErr *sysError.DressNotExistError
		if errors.As(err, &dressNotExistErr) {
			resp.DressNotExistError(dressNotExistErr)
			c.JSON(http.StatusOK, resp)
			return
		}
	}

	respParam := &dressResponse.ShowOneResponse{}
	respParam.Fill(dressBiz)
	data := map[string]interface{}{
		"category": respParam.Category,
		"dress":    respParam.Dress,
	}
	resp.Success(data)
	c.JSON(http.StatusOK, resp)
	return
}

func Update(c *gin.Context) {
	var param *dressRequest.UpdateParam = &dressRequest.UpdateParam{}
	var logger *syslog.UpdateDress = &syslog.UpdateDress{}
	resp := controller.CheckParam(param, c, logger)
	if resp != nil {
		c.JSON(http.StatusOK, resp)
		return
	}

	resp = &response.RespBody{}
	param.ExtractUri()
	dressBiz := &dress.Dress{}
	err := dressBiz.Update(param)
	if err != nil {
		var dbErr *sysError.DbError
		if errors.As(err, &dbErr) {
			resp.DbError(dbErr)
			c.JSON(http.StatusOK, resp)
			return
		}

		var dressNotExistErr *sysError.DressNotExistError
		if errors.As(err, &dressNotExistErr) {
			resp.DressNotExistError(dressNotExistErr)
			c.JSON(http.StatusOK, resp)
			return
		}
	}

	logger.TargetId = param.Dress.Id
	logger.Logger()

	resp.Success(map[string]interface{}{})
	c.JSON(http.StatusOK, resp)
	return
}

// ShowUnusable 展示指定品类下的不可用礼服信息
func ShowUnusable(c *gin.Context) {
	var param *dressRequest.ShowUnusableParam = &dressRequest.ShowUnusableParam{}
	resp := controller.CheckParam(param, c, nil)
	if resp != nil {
		c.JSON(http.StatusOK, resp)
		return
	}

	resp = &response.RespBody{}
	dressBiz := &dress.Dress{}
	categoryBiz, dressBizs, totalPage, count, err := dressBiz.ShowUnusable(param)
	if err != nil {
		var dbErr *sysError.DbError
		if errors.As(err, &dbErr) {
			resp.DbError(dbErr)
			c.JSON(http.StatusOK, resp)
			return
		}

		var categoryNotExistErr *sysError.CategoryNotExistError
		if errors.As(err, &categoryNotExistErr) {
			resp.CategoryNotExistError(categoryNotExistErr)
			c.JSON(http.StatusOK, resp)
			return
		}

		var dressNotExistErr *sysError.DressNotExistError
		if errors.As(err, &dressNotExistErr) {
			resp.DressNotExistError(dressNotExistErr)
			c.JSON(http.StatusOK, resp)
			return
		}
	}

	paginationResp := &pagination.Response{
		CurrentPage: param.Pagination.CurrentPage,
		ItemPerPage: param.Pagination.ItemPerPage,
		TotalPage:   totalPage,
		TotalItem:   count,
	}

	respParam := &dressResponse.ShowUnusableResponse{}
	respParam.Fill(categoryBiz, dressBizs, paginationResp)
	data := map[string]interface{}{
		"category":   respParam.Category,
		"dresses":    respParam.Dresses,
		"pagination": respParam.Pagination,
	}
	resp.Success(data)
	c.JSON(http.StatusOK, resp)
	return
}
