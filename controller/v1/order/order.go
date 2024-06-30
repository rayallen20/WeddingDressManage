package order

import (
	"WeddingDressManage/business/v1/order"
	"WeddingDressManage/controller"
	"WeddingDressManage/lib/sysError"
	orderRequest "WeddingDressManage/param/request/v1/order"
	orderResponse "WeddingDressManage/param/resps/v1/order"
	"WeddingDressManage/param/resps/v1/pagination"
	"WeddingDressManage/response"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Search(c *gin.Context) {
	var param *orderRequest.SearchParam = &orderRequest.SearchParam{}

	resp := controller.CheckParam(param, c, nil)
	if resp != nil {
		c.JSON(http.StatusOK, resp)
		return
	}

	resp = &response.RespBody{}
	// 判断婚期是否早于当天
	if param.SearchCondition.WeddingDate != nil {
		t := time.Now().AddDate(0, 0, 1)
		today := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
		if param.SearchCondition.WeddingDate.IsBefore(today) {
			err := &sysError.DateBeforeTodayError{
				Field: "weddingDate",
			}
			resp.WeddingDateBeforeTodayError(err)
			c.JSON(http.StatusOK, resp)
			return
		}
	}

	// TODO:此处并未实现筛选逻辑 仅将所有品类信息查询并返回了
	orderBiz := &order.Order{}
	categories, totalPage, count, err := orderBiz.Search(param)
	if err != nil {
		if dbError, ok := err.(*sysError.DbError); ok {
			resp.DbError(dbError)
			c.JSON(http.StatusOK, resp)
			return
		}
	}

	searchResp := &orderResponse.SearchResponse{}
	searchResps := searchResp.Generate(categories)
	paginationParam := &pagination.Response{
		CurrentPage: param.Pagination.CurrentPage,
		ItemPerPage: param.Pagination.ItemPerPage,
		TotalPage:   totalPage,
		TotalItem:   count,
	}
	resp.Success(map[string]interface{}{
		"categories": searchResps,
		"pagination": paginationParam,
	})
	c.JSON(http.StatusOK, resp)
	return
}

func PreCreate(c *gin.Context) {
	var param *orderRequest.PreCreateParam = &orderRequest.PreCreateParam{}
	resp := controller.CheckParam(param, c, nil)
	if resp != nil {
		c.JSON(http.StatusOK, resp)
		return
	}

	resp = &response.RespBody{}

	orderBiz := &order.Order{}
	err := orderBiz.PreCreate(param)

	if dbErr, ok := err.(*sysError.DbError); ok {
		resp.DbError(dbErr)
		c.JSON(http.StatusOK, resp)
		return
	}

	if dressNotExistErr, ok := err.(*sysError.DressNotExistError); ok {
		resp.DressNotExistError(dressNotExistErr)
		c.JSON(http.StatusOK, resp)
		return
	}

	respParam := &orderResponse.PreCreateResponse{}
	respParam.Fill(orderBiz)

	resp.Success(map[string]interface{}{
		"order": respParam.Order,
	})
	c.JSON(http.StatusOK, resp)
	return
}

func Discount(c *gin.Context) {
	param := &orderRequest.DiscountParam{}
	resp := controller.CheckParam(param, c, nil)
	if resp != nil {
		c.JSON(http.StatusOK, resp)
		return
	}

	resp = &response.RespBody{}

	orderBiz := &order.Order{}
	err := orderBiz.CalcDiscount(param)

	if err != nil {
		if dbErr, ok := err.(*sysError.DbError); ok {
			resp.DbError(dbErr)
			c.JSON(http.StatusOK, resp)
			return
		}

		if dressNotExistErr, ok := err.(*sysError.DressNotExistError); ok {
			resp.DressNotExistError(dressNotExistErr)
			c.JSON(http.StatusOK, resp)
			return
		}

		if discountInvalidErr, ok := err.(*sysError.DiscountInvalidError); ok {
			resp.DiscountInvalidError(discountInvalidErr)
			c.JSON(http.StatusOK, resp)
			return
		}
	}
	respParam := &orderResponse.DiscountResponse{}
	respParam.Fill(orderBiz)
	data := map[string]interface{}{
		"order": respParam.Order,
	}
	resp.Success(data)
	c.JSON(http.StatusOK, resp)
	return
}

func Create(c *gin.Context) {
	var param *orderRequest.CreateParam = &orderRequest.CreateParam{}
	// TODO:此操作会影响DB变化 需要log
	resp := controller.CheckParam(param, c, nil)
	if resp != nil {
		c.JSON(http.StatusOK, resp)
		return
	}

	resp = &response.RespBody{}

	orderBiz := &order.Order{}
	err := orderBiz.Create(param)
	if err != nil {
		var dbErr *sysError.DbError
		if errors.As(err, &dbErr) {
			resp.DbError(dbErr)
			c.JSON(http.StatusOK, resp)
			return
		}

		var customerBeBannedErr *sysError.CustomerBeBannedError
		if errors.As(err, &customerBeBannedErr) {
			resp.CustomerBeBannedError(customerBeBannedErr)
			c.JSON(http.StatusOK, resp)
			return
		}

		var dressNotExistErr *sysError.DressNotExistError
		if errors.As(err, &dressNotExistErr) {
			resp.DressNotExistError(dressNotExistErr)
			c.JSON(http.StatusOK, resp)
			return
		}

		var dateBeforeTodayErr *sysError.DateBeforeTodayError
		if errors.As(err, &dateBeforeTodayErr) {
			resp.WeddingDateBeforeTodayError(dateBeforeTodayErr)
			c.JSON(http.StatusOK, resp)
			return
		}

		var discountInvalidErr *sysError.DiscountInvalidError
		if errors.As(err, &discountInvalidErr) {
			resp.DiscountInvalidError(discountInvalidErr)
			c.JSON(http.StatusOK, resp)
			return
		}

		var strategyNotExistErr *sysError.StrategyNotExistError
		if errors.As(err, &strategyNotExistErr) {
			resp.StrategyNotExistError(strategyNotExistErr)
			c.JSON(http.StatusOK, resp)
			return
		}

		var customPriceTooFewErr *sysError.CustomPriceTooFewError
		if errors.As(err, &customPriceTooFewErr) {
			resp.CustomPriceTooFewError(customPriceTooFewErr)
			c.JSON(http.StatusOK, resp)
			return
		}
	}

	respParam := &orderResponse.CreateResponse{}
	respParam.Fill(orderBiz, param.Order.PledgeIsSettled)
	data := map[string]interface{}{
		"order": respParam,
	}
	resp.Success(data)
	c.JSON(http.StatusOK, resp)
	return
}

func ShowDelivery(c *gin.Context) {
	var param *orderRequest.ShowDeliveryParam = &orderRequest.ShowDeliveryParam{}
	resp := controller.CheckParam(param, c, nil)
	if resp != nil {
		c.JSON(http.StatusOK, resp)
		return
	}

	resp = &response.RespBody{}
	orderBiz := &order.Order{}
	orders, totalPage, count, err := orderBiz.ShowDelivery(param)
	if err != nil {
		if dbErr, ok := err.(*sysError.DbError); ok {
			resp.DbError(dbErr)
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

	respParam := &orderResponse.ShowDeliveryResponse{}
	respParam.Fill(orders, paginationResp)
	data := map[string]interface{}{
		"orders":     respParam.Orders,
		"pagination": respParam.Pagination,
	}
	resp.Success(data)
	c.JSON(http.StatusOK, resp)
	return
}

func DeliveryDetail(c *gin.Context) {
	var param *orderRequest.DeliveryDetailParam = &orderRequest.DeliveryDetailParam{}
	resp := controller.CheckParam(param, c, nil)
	if resp != nil {
		c.JSON(http.StatusOK, resp)
		return
	}

	resp = &response.RespBody{}
	orderBiz := &order.Order{}
	err := orderBiz.DeliveryDetail(param)
	if err != nil {
		var dbErr *sysError.DbError
		if errors.As(err, &dbErr) {
			resp.DbError(dbErr)
			c.JSON(http.StatusOK, resp)
			return
		}

		var orderNotExistErr *sysError.DeliveryOrderNotExist
		if errors.As(err, &orderNotExistErr) {
			resp.DeliveryOrderNotExistError(orderNotExistErr)
			c.JSON(http.StatusOK, resp)
			return
		}
	}

	respParam := &orderResponse.DeliveryDetailResponse{}
	respParam.Fill(orderBiz)
	data := map[string]interface{}{
		"order": respParam.Order,
	}
	resp.Success(data)
	c.JSON(http.StatusOK, resp)
	return
}

func ShowAmended(c *gin.Context) {
	var param *orderRequest.ShowAmendedParam = &orderRequest.ShowAmendedParam{}
	resp := controller.CheckParam(param, c, nil)
	if resp != nil {
		c.JSON(http.StatusOK, resp)
		return
	}

	resp = &response.RespBody{}
	orderBiz := &order.Order{}
	err := orderBiz.ShowAmended(param)
	if err != nil {
		var dbErr *sysError.DbError
		if errors.As(err, &dbErr) {
			resp.DbError(dbErr)
			c.JSON(http.StatusOK, resp)
			return
		}

		var orderNotExistErr *sysError.DeliveryOrderNotExist
		if errors.As(err, &orderNotExistErr) {
			resp.DeliveryOrderNotExistError(orderNotExistErr)
			c.JSON(http.StatusOK, resp)
			return
		}
	}

	respParam := &orderResponse.ShowAmendedResponse{}
	respParam.Fill(orderBiz)
	data := map[string]interface{}{
		"order":    respParam.Order,
		"customer": respParam.Customer,
		"dresses":  respParam.Dresses,
	}
	resp.Success(data)
	c.JSON(http.StatusOK, resp)
	return
}
