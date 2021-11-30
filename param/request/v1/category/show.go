package category

import (
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/lib/validator"
	"WeddingDressManage/param/request/v1/pagination"
	"github.com/gin-gonic/gin"
)

type ShowParam struct {
	Pagination *pagination.Pagination `from:"pagination" binding:"required"`
}

func (s *ShowParam) Bind(c *gin.Context) error {
	return validator.Bind(s, []interface{}{&pagination.Pagination{}}, c)
}

func (s *ShowParam) Validate(err error) []*sysError.ValidateError {
	return validator.Validate(err)
}