package order

import (
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/lib/validator"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type DiscountParam struct {
	Dresses  []*DiscountDress `form:"dresses" binding:"gt=0,required,unique,dive" errField:"dresses"`
	Discount string           `form:"discount" errField:"discount"`
}

type DiscountDress struct {
	Id int `form:"id" binding:"gt=0,required" errField:"id"`
}

func (d *DiscountParam) Bind(c *gin.Context) error {
	return validator.Bind(d, []interface{}{&DiscountDress{}}, c)
}

func (d *DiscountParam) Validate(errs error) []*sysError.ValidateError {
	validateErrors := validator.Validate(errs)
	discountError := d.validateDiscount()
	if discountError != nil {
		validateErrors = append(validateErrors, discountError)
	}
	return validateErrors
}

func (d *DiscountParam) validateDiscount() *sysError.ValidateError {
	if d.Discount != "" {
		_, err := strconv.ParseFloat(d.Discount, 64)
		if err != nil {
			fmt.Printf("%#v\n", err)
			return &sysError.ValidateError{
				Key: "Discount",
				Msg: "must be a numeric string",
			}
		}
	}
	return nil
}
