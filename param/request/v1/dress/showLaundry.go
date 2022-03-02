package dress

import (
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/lib/validator"
	"WeddingDressManage/param/request/v1/pagination"
	"github.com/gin-gonic/gin"
)

type ShowLaundryParam struct {
	Pagination pagination.Pagination `from:"pagination" binding:"required" errField:"pagination"`
}

func (l *ShowLaundryParam) Bind(c *gin.Context) error {
	return validator.Bind(l, []interface{}{l}, c)
}

func (l *ShowLaundryParam) Validate(err error) []*sysError.ValidateError {
	return validator.Validate(err)
}
