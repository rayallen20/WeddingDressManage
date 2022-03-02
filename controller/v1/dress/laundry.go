package dress

import (
	"WeddingDressManage/business/v1/dress"
	"WeddingDressManage/controller"
	"WeddingDressManage/lib/sysError"
	laundryRequest "WeddingDressManage/param/request/v1/dress"
	laundryResponse "WeddingDressManage/param/resps/v1/dress"
	"WeddingDressManage/response"
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
	laundryRecords, totalPage, err := laundryRecordBiz.Show(param)
	if err != nil {
		if dbError, ok := err.(*sysError.DbError); ok {
			resp.DbError(dbError)
			c.JSON(http.StatusOK, resp)
			return
		}
	}

	showLaundryResp := &laundryResponse.ShowLaundryResponse{}
	showLaundryResp.Fill(laundryRecords, param.Pagination.CurrentPage, totalPage, param.Pagination.ItemPerPage)
	resp.Success(map[string]interface{}{"data": showLaundryResp})
	c.JSON(http.StatusOK, resp)
	return
}
