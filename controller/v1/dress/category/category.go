package category

import (
	"WeddingDressManage/business/v1/dress/category"
	"WeddingDressManage/business/v1/dress/kind"
	"WeddingDressManage/business/v1/dress/unit"
	"WeddingDressManage/lib/response"
	"WeddingDressManage/lib/validator"
	"WeddingDressManage/lib/wdmError"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddParams struct {
	// 礼服品类名称
	KindName string `form:"kindName" binding:"gt=0,required" errField:"kindName"`
	// 礼服品类编码(编码前缀)
	Code string `form:"code" binding:"gt=0,required" errField:"code"`
	// 礼服品类编号
	SerialNumber string `form:"serialNumber" binding:"gt=0,required" errField:"serialNumber"`
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

func Add(c *gin.Context) {
	resp := &response.ResBody{}
	// 校验参数
	param := &AddParams{}
	err := validateAddParam(param, c)
	if err != nil {
		genRespByErrForAdd(resp, err)
		c.JSON(http.StatusOK, resp)
		return
	}

	// step1. 查询 kind & code 是否存在
	kind := &kind.Kind{
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
	category := &category.Category{
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

	// 填充其他品类信息
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

func validateAddParam(param *AddParams, c *gin.Context) (err error) {
	err = c.ShouldBindJSON(param)
	if err != nil {
		if UnmarshalTypeErr, ok := err.(*json.UnmarshalTypeError); ok {
			paramTypeErr := &wdmError.ParamTypeError {
				Message: err.Error(),
				StructFieldName: UnmarshalTypeErr.Field,
			}
			paramTypeErr.GetFormFieldAndShouldType(param)
			return paramTypeErr
		}

		errInfos, ok := validator.GenerateErrsInfo(err)
		if !ok {
			err = wdmError.BindingValidatorError{Message: err.Error()}
			return
		} else {
			err = wdmError.ParamValueError {
				Message: "",
				Details: errInfos,
			}
			return
		}
	}

	var notNumericFields []string

	serialNumber := param.SerialNumber
	if !validator.StringIsNumeric(serialNumber) {
		notNumericFields = append(notNumericFields, "serialNumber")
	}

	if len(notNumericFields) != 0 {
		err = wdmError.NumericStringError {
			Message: "",
			NotNumericFields: notNumericFields,
		}
		return
	}

	return nil
}

func genRespByErrForAdd(resp *response.ResBody, err error) {
	if paramTypeError,ok := err.(*wdmError.ParamTypeError); ok {
		paramTypeError.GetFormFieldAndShouldType(&AddParams{})
		resp.ParamTypeError(paramTypeError)
	} else if bindingErr, ok := err.(wdmError.BindingValidatorError); ok {
		resp.BindingValidatorError(bindingErr, map[string]interface{}{})
	} else if paramValueError, ok := err.(wdmError.ParamValueError); ok {
		resp.ParamValueError(paramValueError)
	} else if numericStringError, ok := err.(wdmError.NumericStringError); ok {
		resp.NumericStringError(numericStringError)
	}
}

func fillCategoryForAdd(category *category.Category, param *AddParams) {
	category.RentableQuantity = param.UnitNumber
	category.Quantity = param.UnitNumber
	category.CharterMoney = param.CharterMoney
	category.CashPledge = param.CashPledge
	category.CoverImg = param.CoverImg
	category.SecondaryImg = param.SecondaryImg
}

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
