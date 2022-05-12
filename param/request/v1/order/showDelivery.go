package order

import (
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/lib/validator"
	"WeddingDressManage/param/request/v1/pagination"
	"github.com/gin-gonic/gin"
)

type ShowDeliveryParam struct {
	Pagination pagination.Pagination `form:"pagination" binding:"required"`
}

func (s *ShowDeliveryParam) Bind(c *gin.Context) error {
	return validator.Bind(s, []interface{}{}, c)
}

func (s *ShowDeliveryParam) Validate(errs error) []*sysError.ValidateError {
	return validator.Validate(errs)
}
