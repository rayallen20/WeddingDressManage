package order

import (
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/lib/validator"
	"github.com/gin-gonic/gin"
)

type DeliveryDetailParam struct {
	Order *DeliveryDetailOrder `form:"order" binding:"required" errFiled:"order"`
}

type DeliveryDetailOrder struct {
	Id int `form:"id" binding:"gt=0,required" errField:"id"`
}

func (d *DeliveryDetailParam) Bind(c *gin.Context) error {
	return validator.Bind(d, []interface{}{d, &DeliveryDetailOrder{}}, c)
}

func (d *DeliveryDetailParam) Validate(errs error) []*sysError.ValidateError {
	return validator.Validate(errs)
}
