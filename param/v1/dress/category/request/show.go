package request

import (
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/lib/validator"
	"github.com/gin-gonic/gin"
)

type ShowParam struct {
	Pagination *Pagination `from:"pagination" binding:"required"`
}

type Pagination struct {
	CurrentPage int `form:"currentPage" binding:"gt=0,required" errField:"currentPage"`
	ItemPerPage int `form:"itemPerPage" binding:"gt=0,required" errField:"itemPerPage"`
}

func (s *ShowParam) Bind(c *gin.Context) error {
	return validator.Bind(s, []interface{}{&Pagination{}}, c)
}

func (s *ShowParam) Validate(err error) []*sysError.ValidateError {
	return validator.Validate(err)
}