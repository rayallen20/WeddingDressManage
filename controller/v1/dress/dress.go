package dress

import (
	"WeddingDressManage/business/v1/dress"
	"WeddingDressManage/lib/response"
	"WeddingDressManage/lib/validator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddRequest /v1/dress/create请求参数结构体
type AddRequest struct {
	KindId int `form:"kindId" binding:"required,gt=0" errField:"kindId"`	// 礼服品类id
	Code string `form:"code" binding:"required,oneof=Q T J X BL BN C XF XK P" errField:"code"`			// 礼服编码
	SerialNumber string `form:"serialNumber" binding:"required" errField:"serialNumber"`	// 礼服编号
	CashPledge int `form:"cashPledge" binding:"required,gt=0" errField:"cashPledge"` // 押金
	RentMoney int `form:"rentMoney" binding:"required,gt=0" errField:"rentMoney"`	// 租金
	Size string `form:"size" binding:"required,oneof=S s M m F f L l XL xl XXL xxl D d" errField:"size"`	// 尺码
}

// Add 添加礼服
func Add(c *gin.Context) {
	// step1. 接收参数并校验
	var params AddRequest
	errResp := validateAddParam(c, &params)
	if errResp != nil {
		c.JSON(http.StatusOK, errResp)
		return
	}

	// step2. 业务逻辑

	// 查询品类信息 若不存在则创建 若存在则将该品类的可租数量+1 总数量+1
	var category *dress.Category = &dress.Category{}
	err := category.Add(params.KindId, params.CashPledge, params.RentMoney, params.Code, params.SerialNumber)
	if err != nil {
		errResp = response.DBErrorResp(err, []interface{}{})
		c.JSON(http.StatusOK, errResp)
		return
	}

	// 添加礼服信息
	var detail *dress.Detail = &dress.Detail{}
	err = detail.Add(category.Id, params.Size)
	if err != nil {
		errResp = response.DBErrorResp(err, []interface{}{})
		c.JSON(http.StatusOK, errResp)
		return
	}

	c.JSON(http.StatusOK, response.SuccessResp([]interface{}{}))
	return
}

// validateAddParam 为CreateAction校验参数
func validateAddParam(c *gin.Context, requestParams *AddRequest) gin.H {
	// 先校验整型 否则后续ShouldBind的err无法转换为validator的err
	kindId := c.PostForm("kindId")
	if kindId != "" {
		_, err := strconv.Atoi(kindId)
		if err != nil {
			return response.FieldNotIntResp("kindId", []interface{}{})
		}
	}

	cashPledge := c.PostForm("cashPledge")
	if cashPledge != "" {
		_, err := strconv.Atoi(cashPledge)
		if err != nil {
			return response.FieldNotIntResp("cashPledge", []interface{}{})
		}
	}

	rentMoney := c.PostForm("rentMoney")
	if rentMoney != "" {
		_, err := strconv.Atoi(rentMoney)
		if err != nil {
			return response.FieldNotIntResp("rentMoney", []interface{}{})
		}
	}

	err := c.ShouldBind(requestParams)

	if err != nil {
		errsInfo, ok := validator.GenerateErrsInfo(err)
		if !ok {
			return response.ConvertBindErrFailedResp([]interface{}{})
		}

		return response.ParamsInvalidResp(errsInfo, []interface{}{})
	}

	_, err = strconv.Atoi(requestParams.SerialNumber)
	if err != nil {
		return response.SerialNumberInvalidResp([]interface{}{})
	}

	return  nil
}
