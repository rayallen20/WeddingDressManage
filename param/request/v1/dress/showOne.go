package dress

import (
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/lib/validator"
	"github.com/gin-gonic/gin"
)

type ShowOneParam struct {
	Dress *ShowOneDressParam `form:"dress" binding:"required" errField:"dress"`
}

type ShowOneDressParam struct {
	Id int `form:"id" binding:"gt=0,required" errField:"id"`
}

func (s *ShowOneParam) Bind(c *gin.Context) error {
	return validator.Bind(s, []interface{}{s, &ShowOneParam{}}, c)
}

func (s *ShowOneParam) Validate(err error) []*sysError.ValidateError {
	return validator.Validate(err)
}
