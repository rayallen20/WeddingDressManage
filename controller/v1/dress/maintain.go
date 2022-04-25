package dress

import (
	"WeddingDressManage/business/v1/dress"
	"WeddingDressManage/controller"
	"WeddingDressManage/lib/sysError"
	maintainRequest "WeddingDressManage/param/request/v1/dress"
	maintainResponse "WeddingDressManage/param/resps/v1/dress"
	"WeddingDressManage/response"
	"WeddingDressManage/syslog"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DailyMaintainGiveBack(c *gin.Context) {
	var param *maintainRequest.MaintainGiveBackParam = &maintainRequest.MaintainGiveBackParam{}
	// TODO:此处日志类型 写死了日常维护 订单模块实现后要修改
	var logger *syslog.DailyMaintainGiveBack = &syslog.DailyMaintainGiveBack{}
	resp := controller.CheckParam(param, c, logger)
	if resp != nil {
		c.JSON(http.StatusOK, resp)
		return
	}

	resp = &response.RespBody{}
	maintainBiz := &dress.MaintainRecord{}
	err := maintainBiz.GiveBack(param)
	if err != nil {
		if dbErr, ok := err.(*sysError.DbError); ok {
			resp.DbError(dbErr)
			c.JSON(http.StatusOK, resp)
			return
		}

		if maintainRecordNotExistErr, ok := err.(*sysError.MaintainRecordNotExistError); ok {
			resp.MaintainRecordNotExistError(maintainRecordNotExistErr)
			c.JSON(http.StatusOK, resp)
			return
		}

		if dressNotExistErr, ok := err.(*sysError.DressNotExistError); ok {
			resp.DressNotExistError(dressNotExistErr)
			c.JSON(http.StatusOK, resp)
			return
		}

		if dressIsNotMaintainingErr, ok := err.(*sysError.DressIsNotMaintainingError); ok {
			resp.DressIsNotMaintainingError(dressIsNotMaintainingErr)
			c.JSON(http.StatusOK, resp)
			return
		}
	}

	logger.TargetId = maintainBiz.Id
	logger.Logger()

	resp.Success(map[string]interface{}{})
	c.JSON(http.StatusOK, resp)
	return
}

func ShowMaintain(c *gin.Context) {
	var param *maintainRequest.ShowMaintainParam = &maintainRequest.ShowMaintainParam{}
	resp := controller.CheckParam(param, c, nil)
	if resp != nil {
		c.JSON(http.StatusOK, resp)
		return
	}

	resp = &response.RespBody{}
	maintainBiz := &dress.MaintainRecord{}
	maintainBizs, totalPage, count, err := maintainBiz.Show(param)
	if err != nil {
		if dbErr, ok := err.(*sysError.DbError); ok {
			resp.DbError(dbErr)
			c.JSON(http.StatusOK, resp)
			return
		}
	}

	showMaintainResp := &maintainResponse.ShowMaintainResponse{}
	showMaintainResp.Fill(maintainBizs, param.Pagination.CurrentPage, totalPage, count, param.Pagination.ItemPerPage)
	resp.Success(map[string]interface{}{
		"pagination": showMaintainResp.Pagination,
		"maintains":  showMaintainResp.Maintains,
	})
	c.JSON(http.StatusOK, resp)
	return
}
