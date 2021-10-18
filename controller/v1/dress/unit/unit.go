package unit

import (
	"WeddingDressManage/business/v1/dress/category"
	"WeddingDressManage/business/v1/dress/unit"
	"WeddingDressManage/conf"
	"WeddingDressManage/lib/response"
	"WeddingDressManage/lib/validator"
	"WeddingDressManage/model"
	"github.com/gin-gonic/gin"
	"math"
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
	err := validator.ValidateParam(param, c)
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

	category := &category.Category {
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

type ShowUsableParams struct {
	CategoryId int `form:"categoryId" binding:"gte=1,required" errField:"categoryId"`
	SerialNumber string `form:"serialNumber" binding:"gte=1,required" errField:"serialNumber"`
	Page int `form:"page" binding:"gte=1,required" errField:"page"`
}

// ShowUsable 查看指定品类下可用(非赠与且非废弃状态)的礼服信息集合
func ShowUsable(c *gin.Context) {
	resp := &response.ResBody{}
	param := &ShowUsableParams{}
	err := validator.ValidateParam(param, c)
	if err != nil {
		resp.GenRespByParamErr(err)
		c.JSON(http.StatusOK, resp)
		return
	}

	// 查询品类信息是否存在
	category := &category.Category{
		Id: param.CategoryId,
	}
	err = category.ExistById()
	if err != nil {
		resp.DBError(err, map[string]interface{}{})
		c.JSON(http.StatusOK, resp)
		return
	}

	sn := category.Code + "-" + category.SerialNumber
	if category.SerialNumber == "" || param.SerialNumber != sn {
		resp.CategoryHasNotExist(map[string]interface{}{})
		c.JSON(http.StatusOK, resp)
		return
	}

	unit := unit.Unit{}
	units, err := unit.ShowUsable(param.CategoryId, param.Page)
	if err != nil {
		resp.DBError(err, map[string]interface{}{})
		c.JSON(http.StatusOK, resp)
		return
	}

	usableCount, err := unit.CountCategoryUsable(param.CategoryId)
	if err != nil {
		resp.DBError(err, map[string]interface{}{})
		c.JSON(http.StatusOK, resp)
		return
	}

	totalPage := int(math.Ceil(float64(usableCount) / float64(conf.Conf.DataBase.PageSize)))

	data := map[string]interface{}{
		"units":genRespDataForShowUsable(units, sn),
		"totalPage":totalPage,
	}
	resp.Success(data)
	c.JSON(http.StatusOK, resp)
	return
}

type ShowUsableRespData struct {
	UnitId int `json:"unitId"`
	CategorySerialNumber string `json:"CategorySerialNumber"`
	UnitSerialNumber int `json:"unitSerialNumber"`
	Size string `json:"size"`
	Status string `json:"status"`
	CoverImg string `json:"coverImg"`
	SecondaryImg []string `json:"secondaryImg"`
}

func genRespDataForShowUsable(units []unit.Unit, categorySN string) []ShowUsableRespData {
	respDatas := make([]ShowUsableRespData, 0, len(units))

	for i := 0; i < len(units); i++ {
		respData := ShowUsableRespData{
			UnitId:               units[i].Id,
			CategorySerialNumber: categorySN,
			UnitSerialNumber:     units[i].SerialNumber,
			Size:                 units[i].Size,
			Status:               units[i].Status,
			CoverImg:             units[i].CoverImg,
			SecondaryImg:         units[i].SecondaryImg,
		}
		respDatas = append(respDatas, respData)
	}
	return respDatas
}

type ShowUnusableParams struct {
	CategoryId int `form:"categoryId" binding:"gte=1,required" errField:"categoryId"`
	SerialNumber string `form:"serialNumber" binding:"gte=1,required" errField:"serialNumber"`
	Page int `form:"page" binding:"gte=1,required" errField:"page"`
}

// ShowUnusable 查看指定品类下不可用(赠与 或 废弃状态)的礼服信息集合
func ShowUnusable(c *gin.Context) {
	resp := &response.ResBody{}
	param := &ShowUsableParams{}
	err := validator.ValidateParam(param, c)
	if err != nil {
		resp.GenRespByParamErr(err)
		c.JSON(http.StatusOK, resp)
		return
	}

	// 查询品类信息是否存在
	category := &category.Category{
		Id: param.CategoryId,
	}
	err = category.ExistById()
	if err != nil {
		resp.DBError(err, map[string]interface{}{})
		c.JSON(http.StatusOK, resp)
		return
	}

	sn := category.Code + "-" + category.SerialNumber
	if category.SerialNumber == "" || param.SerialNumber != sn {
		resp.CategoryHasNotExist(map[string]interface{}{})
		c.JSON(http.StatusOK, resp)
		return
	}

	unit := unit.Unit{}
	units, err := unit.ShowUnusable(param.CategoryId, param.Page)
	if err != nil {
		resp.DBError(err, map[string]interface{}{})
		c.JSON(http.StatusOK, resp)
		return
	}

	usableCount, err := unit.CountCategoryUnusable(param.CategoryId)
	if err != nil {
		resp.DBError(err, map[string]interface{}{})
		c.JSON(http.StatusOK, resp)
		return
	}

	totalPage := int(math.Ceil(float64(usableCount) / float64(conf.Conf.DataBase.PageSize)))

	data := map[string]interface{}{
		"units":genRespDataForShowUnusable(units, sn),
		"totalPage":totalPage,
	}
	resp.Success(data)
	c.JSON(http.StatusOK, resp)
	return
}

type ShowUnusableRespData struct {
	UnitId int `json:"unitId"`
	CategorySerialNumber string `json:"CategorySerialNumber"`
	UnitSerialNumber int `json:"unitSerialNumber"`
	Size string `json:"size"`
	Status string `json:"status"`
	CoverImg string `json:"coverImg"`
	SecondaryImg []string `json:"secondaryImg"`
}

func genRespDataForShowUnusable(units []unit.Unit, categorySN string) []ShowUsableRespData {
	respDatas := make([]ShowUsableRespData, 0, len(units))

	for i := 0; i < len(units); i++ {
		respData := ShowUsableRespData{
			UnitId:               units[i].Id,
			CategorySerialNumber: categorySN,
			UnitSerialNumber:     units[i].SerialNumber,
			Size:                 units[i].Size,
			Status:               units[i].Status,
			CoverImg:             units[i].CoverImg,
			SecondaryImg:         units[i].SecondaryImg,
		}
		respDatas = append(respDatas, respData)
	}
	return respDatas
}