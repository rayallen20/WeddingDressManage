package category

import (
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/lib/validator"
	"github.com/gin-gonic/gin"
)

type ShowOneParam struct {
	Category *showOneCategoryParam `form:"category" binding:"required" errField:"category"`
}

type showOneCategoryParam struct {
	Id int `form:"id" binding:"gt=0,required" errField:"id"`
}

func (s *ShowOneParam) Bind(c *gin.Context) error {
	return c.ShouldBindJSON(s)
}

func (s *ShowOneParam) Validate(err error) []*sysError.ValidateError {
	return validator.Validate(err)
}
