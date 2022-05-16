package order

import (
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/lib/validator"
	"github.com/gin-gonic/gin"
)

type ShowAmendedParam struct {
	Order *ShowAmendedOrder `form:"order" binding:"required"`
}

type ShowAmendedOrder struct {
	Id int `form:"id" binging:"gt=0,required"`
}

func (s *ShowAmendedParam) Bind(c *gin.Context) error {
	return validator.Bind(s, []interface{}{s, &ShowAmendedOrder{}}, c)
}

func (s *ShowAmendedParam) Validate(errs error) []*sysError.ValidateError {
	return validator.Validate(errs)
}
