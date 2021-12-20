package dress

import (
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/lib/validator"
	"github.com/gin-gonic/gin"
)

type ApplyDiscardParam struct {
	Dress      *ApplyDiscardDressParam      `form:"dress" binding:"required" errField:"dress"`
	DiscardAsk *ApplyDiscardDiscardAskParam `form:"discardAsk" binding:"required" errField:"discardAsk"`
}

type ApplyDiscardDressParam struct {
	Id int `form:"id" binding:"gt=0,required" errField:"id"`
}

type ApplyDiscardDiscardAskParam struct {
	Note string `form:"note" binding:"gt=0,required" errField:"note"`
}

func (a *ApplyDiscardParam) Bind(c *gin.Context) error {
	return validator.Bind(a, []interface{}{a, &ApplyDiscardDressParam{}, &ApplyDiscardDiscardAskParam{}}, c)
}

func (a ApplyDiscardParam) Validate(err error) []*sysError.ValidateError {
	return validator.Validate(err)
}
