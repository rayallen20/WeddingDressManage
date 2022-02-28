package dress

import (
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/lib/validator"
	"WeddingDressManage/param/request/v1/pagination"
	"github.com/gin-gonic/gin"
)

type ShowUnusableParam struct {
	Pagination pagination.Pagination      `from:"pagination" binding:"required" errField:"pagination"`
	Category   *ShowUnusableCategoryParam `from:"category" binding:"required" errField:"category"`
}

type ShowUnusableCategoryParam struct {
	Id int `from:"id" binding:"gt=0,required" errField:"id"`
}

func (s *ShowUnusableParam) Bind(c *gin.Context) error {
	return validator.Bind(s, []interface{}{s, &ShowUsableCategoryParam{}}, c)
}

func (s *ShowUnusableParam) Validate(err error) []*sysError.ValidateError {
	return validator.Validate(err)
}
