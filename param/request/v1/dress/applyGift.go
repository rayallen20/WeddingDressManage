package dress

import (
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/lib/validator"
	"github.com/gin-gonic/gin"
)

type ApplyGiftParam struct {
	Dress    *ApplyGiftDressParam    `form:"dress" binding:"required" errField:"dress"`
	Customer *ApplyGiftCustomerParam `form:"customer" binding:"required" errField:"customer"`
	GiftAsk  *ApplyGiftGiftAskParam  `form:"giftAsk" binding:"required" errField:"giftAsk"`
}

type ApplyGiftDressParam struct {
	Id int `form:"id" binding:"gt=0,required" errField:"id"`
}

type ApplyGiftCustomerParam struct {
	Name   string `from:"name" binding:"gt=0,required" errField:"name"`
	Mobile string `from:"mobile" binding:"gt=0,mobile,required" errField:"mobile"`
}

type ApplyGiftGiftAskParam struct {
	Note string `form:"note" binding:"gt=0,required" errField:"note"`
}

func (a *ApplyGiftParam) Bind(c *gin.Context) error {
	return validator.Bind(a, []interface{}{a, &ApplyGiftDressParam{}, &ApplyGiftCustomerParam{}, &ApplyGiftGiftAskParam{}}, c)
}

func (a ApplyGiftParam) Validate(err error) []*sysError.ValidateError {
	return validator.Validate(err)
}
