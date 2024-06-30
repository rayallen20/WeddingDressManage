package dress

import (
	"WeddingDressManage/business/v1/dress"
	"WeddingDressManage/controller"
	"WeddingDressManage/lib/sysError"
	laundryRequest "WeddingDressManage/param/request/v1/dress"
	laundryResponse "WeddingDressManage/param/resps/v1/dress"
	"WeddingDressManage/response"
	"WeddingDressManage/syslog"
	"errors"
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
		var dbError *sysError.DbError
		if errors.As(err, &dbError) {
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
		var dbError *sysError.DbError
		if errors.As(err, &dbError) {
			resp.DbError(dbError)
			c.JSON(http.StatusOK, resp)
			return
		}

		var laundryRecordNotExistErr *sysError.LaundryRecordNotExistError
		if errors.As(err, &laundryRecordNotExistErr) {
			resp.LaundryRecordNotExistError(laundryRecordNotExistErr)
			c.JSON(http.StatusOK, resp)
			return
		}

		var dressNotExistErr *sysError.DressNotExistError
		if errors.As(err, &dressNotExistErr) {
			resp.DressNotExistError(dressNotExistErr)
			c.JSON(http.StatusOK, resp)
			return
		}

		var dressIsNotLaunderingErr *sysError.DressIsNotLaunderingError
		if errors.As(err, &dressIsNotLaunderingErr) {
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
