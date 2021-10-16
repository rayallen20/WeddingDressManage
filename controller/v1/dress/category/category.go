package category

import (
	"WeddingDressManage/business/v1/dress/category"
	"WeddingDressManage/business/v1/dress/kind"
	"WeddingDressManage/business/v1/dress/unit"
	"WeddingDressManage/conf"
	"WeddingDressManage/lib/response"
	"WeddingDressManage/lib/validator"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
)

type AddParams struct {
	// 礼服品类名称
	KindName string `form:"kindName" binding:"gt=0,required" errField:"kindName"`
	// 礼服品类编码(编码前缀)
	Code string `form:"code" binding:"gt=0,required" errField:"code"`
	// 礼服品类编号
	SerialNumber string `form:"serialNumber" binding:"gt=0,required" errField:"serialNumber" numeric:"true"`
	// 租金
	CharterMoney int `form:"charterMoney" binding:"gt=0,required" errField:"charterMoney"`
	// 押金
	CashPledge int `form:"cashPledge" binding:"gt=0,required" errField:"cashPledge"`
	// 尺码
	Size string `form:"size" binding:"gt=0,required,oneof=S M F L XL XXL D" errField:"size"`
	// 件数
	UnitNumber int `form:"unitNumber" binding:"gt=0,required" errField:"unitNumber"`
	// 封面图
	CoverImg string `form:"coverImg" binding:"gt=0,required" errField:"coverImg"`
	// 副图
	// TODO:此处的元素数量应该从配置中读取
	SecondaryImg []string `form:"secondaryImg" binding:"gt=0,lte=1" errField:"secondaryImg"`
}

// Add 添加礼服品类的同时 添加至少1件礼服
func Add(c *gin.Context) {
	resp := &response.ResBody{}
	// 校验参数
	param := &AddParams{}
	err := validator.ValidateParam(param, c)
	if err != nil {
		resp.GenRespByParamErr(err)
		c.JSON(http.StatusOK, resp)
		return
	}

	// step1. 查询 kind & code 是否存在
	kind := &kind.Kind {
		Name: param.KindName,
		Code: param.Code,
	}
	err = kind.FindByNameAndCode()
	if err != nil {
		resp.DBError(err, map[string]interface{}{})
		c.JSON(http.StatusOK, resp)
		return
	}

	if kind.Id == 0 {
		resp.KindNotExistError(map[string]interface{}{})
		c.JSON(http.StatusOK, resp)
		return
	}

	// step2. 查询 code & serialNumber是否存在
	category := &category.Category {
		KindId: kind.Id,
		Code: param.Code,
		SerialNumber: param.SerialNumber,
	}
	err = category.FindByKindIdAndCodeAndSN()
	if err != nil {
		resp.DBError(err, map[string]interface{}{})
		c.JSON(http.StatusOK, resp)
		return
	}

	if category.Id != 0 {
		resp.CategoryHasExistedError(map[string]interface{}{})
		c.JSON(http.StatusOK, resp)
		return
	}
	// step3. 事务:创建品类 & 创建礼服
	fillCategoryForAdd(category, param)

	// 创建礼服
	units := makeUnitsForAdd(param)

	units, err = category.Add(units)
	if err != nil {
		resp.TransactionError(err, map[string]interface{}{})
		c.JSON(http.StatusOK, resp)
		return
	}

	data := genRespDataForAdd(category)
	resp.Success(data)
	c.JSON(http.StatusOK, resp)
	return
}

//fillCategoryForAdd 为Add函数根据请求参数填充其他品类信息
func fillCategoryForAdd(category *category.Category, param *AddParams) {
	category.RentableQuantity = param.UnitNumber
	category.Quantity = param.UnitNumber
	category.CharterMoney = param.CharterMoney
	category.CashPledge = param.CashPledge
	category.CoverImg = param.CoverImg
	category.SecondaryImg = param.SecondaryImg
}

//makeUnitsForAdd 为Add函数根据请求参数创建具体礼服对象
func makeUnitsForAdd(param *AddParams) []*unit.Unit {
	units := make([]*unit.Unit, 0, param.UnitNumber)
	for i := 0; i < param.UnitNumber; i++ {
		unit := &unit.Unit {
			SerialNumber: i + 1,
			Size: param.Size,
			CoverImg: param.CoverImg,
			SecondaryImg: param.SecondaryImg,
		}
		units = append(units, unit)
	}
	return units
}

type AddRespData struct {
	Id int `json:"id"`
	CodeAndSN string `json:"codeAndSN"`
	RentNumber int `json:"rentNumber"`
	RentableQuantity int `json:"rentableQuantity"`
	AvgRentMoney int `json:"avgRentMoney"`
	CoverImg string `json:"coverImg"`
	SecondaryImg []string `json:"secondaryImg"`
}

// genRespDataForAdd 为Add函数生成响应体的data部分
func genRespDataForAdd(category *category.Category) (data map[string]interface{}) {
	respData := &AddRespData{
		Id:               category.Id,
		CodeAndSN:        category.Code + "-" + category.SerialNumber,
		RentNumber:       category.RentNumber,
		RentableQuantity: category.RentableQuantity,
		AvgRentMoney:     category.AvgRentMoney,
		CoverImg:         category.CoverImg,
		SecondaryImg:     category.SecondaryImg,
	}
	data = map[string]interface{}{
		"categoryInfo":respData,
	}
	return
}

type ShowParam struct {
	Page int `form:"page" binding:"gte=1,required" errField:"page"`
}

// Show 展示所有可用的礼服品类信息
func Show(c *gin.Context) {
	resp := &response.ResBody{}
	param := &ShowParam{}
	err := validator.ValidateParam(param, c)
	if err != nil {
		resp.GenRespByParamErr(err)
		c.JSON(http.StatusOK, resp)
		return
	}

	category := category.Category{}
	totalUsableCategory, err := category.CountTotalUsable()
	if err != nil {
		resp.DBError(err, map[string]interface{}{})
		c.JSON(http.StatusOK, resp)
		return
	}

	totalPage := int(math.Ceil(float64(totalUsableCategory) / float64(conf.Conf.DataBase.PageSize)))

	if param.Page > totalPage {
		resp.PageBeyondMaximumError(totalPage)
		c.JSON(http.StatusOK, resp)
		return
	}

	categoryies, err := category.Show(param.Page)
	if err != nil {
		resp.DBError(err, map[string]interface{}{})
		c.JSON(http.StatusOK, resp)
		return
	}

	data := genRespDataForShow(categoryies)
	data["totalPage"] = totalPage
	resp.Success(data)
	c.JSON(http.StatusOK, resp)
	return
}

type ShowRespData struct {
	SerialNumber string
	RentNumber int
	RentableQuantity int
	AvgRentMoney int
	CoverImg string
	SecondaryImg []string
}

func genRespDataForShow(categoryies []category.Category) (data map[string]interface{}) {
	categoriesResps := make([]ShowRespData, 0, len(categoryies))

	for _, category := range categoryies {
		categoriesResp := ShowRespData{
			SerialNumber:     category.Code + "-" + category.SerialNumber,
			RentNumber:       category.RentNumber,
			RentableQuantity: category.RentableQuantity,
			AvgRentMoney:     category.AvgRentMoney,
			CoverImg: category.CoverImg,
			SecondaryImg: category.SecondaryImg,
		}
		categoriesResps = append(categoriesResps, categoriesResp)
	}

	data = map[string]interface{}{
		"categories":categoriesResps,
	}
	return
}