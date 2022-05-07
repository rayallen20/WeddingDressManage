package order

import (
	"WeddingDressManage/lib/helper/paramHelper"
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/lib/validator"
	"WeddingDressManage/param"
	"github.com/gin-gonic/gin"
	"strconv"
)

type CreateParam struct {
	Customer *CreateCustomerParam `form:"customer" binding:"required"`
	Order    *CreateOrderParam    `form:"order" binding:"required"`
}

type CreateCustomerParam struct {
	Name   string `form:"name" binding:"required"`
	Mobile string `form:"mobile" binding:"required"` // mobile
}

type CreateOrderParam struct {
	WeddingDate     param.Date         `form:"weddingDate" binding:"required" errField:"weddingDate"`
	Items           []*CreateItemParam `form:"items" binding:"gt=0,required,unique,dive"`
	SaleStrategy    *SaleStrategy      `form:"saleStrategy" binding:"required"`
	PledgeIsSettled bool               `form:"pledgeIsSettled"`
	Comment         string             `form:"comment"`
}

type CreateItemParam struct {
	Dress *CreateDressParam `form:"dress" binding:"required"`
}

type CreateDressParam struct {
	Id int `form:"id" validate:"gt=0,required"`
}

type SaleStrategy struct {
	Type               string `form:"type" binding:"oneof=originalPrice discount customPrice,required"`
	Discount           string `form:"discount" validate:"numeric"`
	CustomPriceCharter string `form:"customPriceCharter" validate:"numeric"`
	CustomPricePledge  string `form:"customPricePledge" validate:"numeric"`
}

func (p *CreateParam) Bind(c *gin.Context) error {
	return validator.Bind(p, []interface{}{&CreateCustomerParam{}, &CreateOrderParam{}, make([]*CreateItemParam, 0, 0),
		&CreateItemParam{}, &CreateDressParam{}, &SaleStrategy{}}, c)
}

func (p *CreateParam) Validate(errs error) []*sysError.ValidateError {
	validateErrors := validator.Validate(errs)

	if p.existRepetitionDressId() {
		validateError := &sysError.ValidateError{
			Key: "Order.Items",
			Msg: "must contains unique element",
		}
		validateErrors = append(validateErrors, validateError)
	}

	if !paramHelper.IsMobile(p.Customer.Mobile) {
		validateError := &sysError.ValidateError{
			Key: "Customer.Mobile",
			Msg: "must be a phone cell-phone number",
		}
		validateErrors = append(validateErrors, validateError)
	}

	strategyErr := p.validateSaleStrategy()
	if strategyErr != nil {
		validateErrors = append(validateErrors, strategyErr)
	}

	return validateErrors
}

// existRepetitionDressId 校验items中的id是否唯一
func (p *CreateParam) existRepetitionDressId() bool {
	uniqueDressIds := make([]int, 0, len(p.Order.Items))
	for _, item := range p.Order.Items {
		for _, uniqueDressId := range uniqueDressIds {
			if item.Dress.Id == uniqueDressId {
				return true
			}
		}
		uniqueDressIds = append(uniqueDressIds, item.Dress.Id)
	}
	return false
}

// strconv.ParseFloat(param.Order.SaleStrategy.Discount, 64)

// validateSaleStrategy 校验优惠策略对象的参数
// 校验规则:
// 若优惠策略为打折 则Discount字段必填
// 若优惠策略为自定义 则CustomPriceCharter字段和CustomPricePledge字段必填
func (p *CreateParam) validateSaleStrategy() *sysError.ValidateError {
	strategy := p.Order.SaleStrategy
	if strategy.Type == "discount" {
		_, err := strconv.ParseFloat(strategy.Discount, 64)
		if err != nil {
			return &sysError.ValidateError{
				Key: "Order.SaleStrategy.Discount",
				Msg: "must be a numeric string",
			}
		}
	}

	if strategy.Type == "customPrice" {
		_, err := strconv.ParseFloat(strategy.CustomPriceCharter, 64)
		if err != nil {
			return &sysError.ValidateError{
				Key: "Order.SaleStrategy.CustomPriceCharter",
				Msg: "must be a numeric string",
			}
		}

		_, err = strconv.ParseFloat(strategy.CustomPricePledge, 64)
		if err != nil {
			return &sysError.ValidateError{
				Key: "Order.SaleStrategy.CustomPricePledge",
				Msg: "must be a numeric string",
			}
		}
	}

	return nil
}
