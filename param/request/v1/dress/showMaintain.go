package dress

import (
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/lib/validator"
	"WeddingDressManage/param/request/v1/pagination"
	"github.com/gin-gonic/gin"
)

type ShowMaintainParam struct {
	Pagination pagination.Pagination `json:"pagination" binding:"required" errField:"pagination"`
}

func (m *ShowMaintainParam) Bind(c *gin.Context) error {
	return validator.Bind(m, []interface{}{m, &pagination.Pagination{}}, c)
}

func (m *ShowMaintainParam) Validate(err error) []*sysError.ValidateError {
	return validator.Validate(err)
}
