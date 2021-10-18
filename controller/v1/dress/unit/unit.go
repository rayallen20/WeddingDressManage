package unit

import (
	"WeddingDressManage/business/v1/dress/category"
	"WeddingDressManage/business/v1/dress/unit"
	"WeddingDressManage/lib/response"
	"WeddingDressManage/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AddParams struct {
	CategoryId int `form:"categoryId" binding:"gte=1,required" errField:"categoryId"`
	SerialNumber string `form:"serialNumber" binding:"gte=1,required" errField:"serialNumber"`
	Size string `form:"size" binding:"gt=0,required,oneof=S M F L XL XXL D" errField:"size"`
	UnitNumber int `form:"unitNumber" binding:"gt=0,required" errField:"unitNumber"`
	CoverImg string `form:"coverImg" binding:"gt=0,required" errField:"coverImg"`
	SecondaryImg []string `form:"secondaryImg" binding:"gt=0,lte=1" errField:"secondaryImg"`
}

// Add 在已有的品类下 添加多件礼服
func Add(c *gin.Context) {
	resp := &response.ResBody{}
	param := &AddParams{}
	err := c.ShouldBindJSON(param)
	if err != nil {
		resp.GenRespByParamErr(err)
		c.JSON(http.StatusOK, resp)
		return
	}

	codeAndSN := strings.Split(param.SerialNumber, "-")
	if len(codeAndSN) != 2 {
		resp.SNFormatError(map[string]interface{}{})
		c.JSON(http.StatusOK, resp)
		return
	}

	category := &category.Category{
		Id: param.CategoryId,
	}
	err = category.ExistById()
	if err != nil {
		resp.DBError(err, map[string]interface{}{})
		c.JSON(http.StatusOK, resp)
		return
	}
	
	if category.Status == model.CategoryStatus["unusable"] {
		resp.CategoryIsUnusable(map[string]interface{}{})
		c.JSON(http.StatusOK, resp)
		return
	}

	sn := category.Code + "-" + category.SerialNumber
	if sn != param.SerialNumber {
		resp.CategoryHasNotExist(map[string]interface{}{})
		c.JSON(http.StatusOK, resp)
		return
	}

	category.Quantity += param.UnitNumber
	category.RentableQuantity += param.UnitNumber

	unitModel := &model.DressUnit{
		CategoryId: category.Id,
	}
	err = unitModel.FindMaxSerialNumberByCategoryId()
	if err != nil {
		resp.DBError(err, map[string]interface{}{})
		c.JSON(http.StatusOK, resp)
		return
	}

	units := make([]unit.Unit, 0, param.UnitNumber)

	for i := 0; i < param.UnitNumber; i++ {
		unit := unit.Unit{
			CategoryId: category.Id,
			SerialNumber: unitModel.SerialNumber + i + 1,
			Size: param.Size,
			CoverImg: param.CoverImg,
			SecondaryImg: param.SecondaryImg,
			Status: model.UnitStatus["rentable"],
		}
		units = append(units, unit)
	}

	err = category.AddUnits(units)
	resp.Success(map[string]interface{}{})
	c.JSON(http.StatusOK, resp)
	return
}