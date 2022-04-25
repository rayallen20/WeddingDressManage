package dress

import (
	"WeddingDressManage/business/v1/dress"
	"WeddingDressManage/controller"
	"WeddingDressManage/lib/sysError"
	laundryRequest "WeddingDressManage/param/request/v1/dress"
	laundryResponse "WeddingDressManage/param/resps/v1/dress"
	"WeddingDressManage/response"
	"WeddingDressManage/syslog"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowLaundry(c *gin.Context) {
	var param *laundryRequest.ShowLaundryParam = &laundryRequest.ShowLaundryParam{}
	resp := controller.CheckParam(param, c, nil)
	if resp != nil {
		c.JSON(http.StatusOK, resp)
		return
	}

	resp = &response.RespBody{}
	laundryRecordBiz := &dress.LaundryRecord{}
	laundryRecords, totalPage, count, err := laundryRecordBiz.Show(param)
	if err != nil {
		if dbError, ok := err.(*sysError.DbError); ok {
			resp.DbError(dbError)
			c.JSON(http.StatusOK, resp)
			return
		}
	}

	showLaundryResp := &laundryResponse.ShowLaundryResponse{}
	showLaundryResp.Fill(laundryRecords, param.Pagination.CurrentPage, totalPage, count, param.Pagination.ItemPerPage)
	resp.Success(map[string]interface{}{
		"pagination": showLaundryResp.Pagination,
		"laundries":  showLaundryResp.Laundries,
	})
	c.JSON(http.StatusOK, resp)
	return
}

func LaundryGiveBack(c *gin.Context) {
	var param *laundryRequest.LaundryGiveBackParam = &laundryRequest.LaundryGiveBackParam{}
	var logger *syslog.LaundryGiveBack = &syslog.LaundryGiveBack{}
	resp := controller.CheckParam(param, c, logger)
	if resp != nil {
		c.JSON(http.StatusOK, resp)
		return
	}

	resp = &response.RespBody{}

	laundryRecordBiz := &dress.LaundryRecord{}
	err := laundryRecordBiz.GiveBack(param)
	if err != nil {
		if dbError, ok := err.(*sysError.DbError); ok {
			resp.DbError(dbError)
			c.JSON(http.StatusOK, resp)
			return
		}

		if laundryRecordNotExistErr, ok := err.(*sysError.LaundryRecordNotExistError); ok {
			resp.LaundryRecordNotExistError(laundryRecordNotExistErr)
			c.JSON(http.StatusOK, resp)
			return
		}

		if dressNotExistErr, ok := err.(*sysError.DressNotExistError); ok {
			resp.DressNotExistError(dressNotExistErr)
			c.JSON(http.StatusOK, resp)
			return
		}

		if dressIsNotLaunderingErr, ok := err.(*sysError.DressIsNotLaunderingError); ok {
			resp.DressIsNotLaunderingError(dressIsNotLaunderingErr)
			c.JSON(http.StatusOK, resp)
			return
		}
	}

	logger.TargetId = laundryRecordBiz.Id
	logger.Logger()

	resp.Success(map[string]interface{}{})
	c.JSON(http.StatusOK, resp)
	return
}
