package dress

import (
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/lib/validator"
	"github.com/gin-gonic/gin"
)

type ApplyGiftParam struct {
	Dress    *applyGiftDressParam    `form:"dress" binding:"required" errField:"dress"`
	Customer *applyGiftCustomerParam `form:"customer" binding:"required" errField:"customer"`
	GiftAsk  *applyGiftGiftAskParam  `form:"giftAsk" binding:"required" errField:"giftAsk"`
}

type applyGiftDressParam struct {
	Id int `form:"id" binding:"gt=0,required" errField:"id"`
}

type applyGiftCustomerParam struct {
	Name   string `from:"name" binding:"gt=0,required" errField:"name"`
	Mobile string `from:"mobile" binding:"gt=0,mobile,required" errField:"mobile"`
}

type applyGiftGiftAskParam struct {
	Note string `form:"note" binding:"gt=0,required" errField:"note"`
}

func (a *ApplyGiftParam) Bind(c *gin.Context) error {
	return c.ShouldBindJSON(a)
}

func (a ApplyGiftParam) Validate(err error) []*sysError.ValidateError {
	return validator.Validate(err)
}
