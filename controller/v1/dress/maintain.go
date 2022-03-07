package dress

import (
	"WeddingDressManage/business/v1/dress"
	"WeddingDressManage/controller"
	"WeddingDressManage/lib/sysError"
	maintainRequest "WeddingDressManage/param/request/v1/dress"
	maintainResponse "WeddingDressManage/param/resps/v1/dress"
	"WeddingDressManage/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DailyMaintainGiveBack(c *gin.Context) {

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
	resp.Success(map[string]interface{}{"data": showMaintainResp})
	c.JSON(http.StatusOK, resp)
	return
}
