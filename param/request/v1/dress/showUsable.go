package dress

import (
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/lib/validator"
	"WeddingDressManage/param/request/v1/pagination"
	"github.com/gin-gonic/gin"
)

type ShowUsableParam struct {
	Pagination pagination.Pagination    `form:"pagination" binding:"required" errField:"pagination"`
	Category   *ShowUsableCategoryParam `form:"category" binding:"required" errField:"category"`
}

type ShowUsableCategoryParam struct {
	Id int `form:"id" binding:"gt=0,required" errField:"id"`
}

func (s *ShowUsableParam) Bind(c *gin.Context) error {
	return validator.Bind(s, []interface{}{s, &ShowUsableCategoryParam{}}, c)
}

func (s *ShowUsableParam) Validate(err error) []*sysError.ValidateError {
	return validator.Validate(err)
}
