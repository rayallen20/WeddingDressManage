package order

import (
	"WeddingDressManage/business/v1/order"
	"WeddingDressManage/controller"
	"WeddingDressManage/lib/sysError"
	orderRequest "WeddingDressManage/param/request/v1/order"
	orderResponse "WeddingDressManage/param/resps/v1/order"
	"WeddingDressManage/param/resps/v1/pagination"
	"WeddingDressManage/response"
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
