package dress

import (
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/lib/validator"
	"github.com/gin-gonic/gin"
)

type ApplyDiscardParam struct {
	Dress      *applyDiscardDressParam      `form:"dress" binding:"required" errField:"dress"`
	DiscardAsk *applyDiscardDiscardAskParam `form:"discardAsk" binding:"required" errField:"discardAsk"`
}

type applyDiscardDressParam struct {
	Id int `form:"id" binding:"gt=0,required" errField:"id"`
}

type applyDiscardDiscardAskParam struct {
	Note string `form:"note" binding:"gt=0,required" errField:"note"`
}

func (a *ApplyDiscardParam) Bind(c *gin.Context) error {
	return c.ShouldBindJSON(a)
}

func (a ApplyDiscardParam) Validate(err error) []*sysError.ValidateError {
	return validator.Validate(err)
}
