package dress

import (
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/lib/validator"
	"WeddingDressManage/param/request/v1/pagination"
	"github.com/gin-gonic/gin"
)

type ShowUsableParam struct {
	Pagination pagination.Pagination `form:"pagination" binding:"required" errField:"pagination"`
	Category *showUsableCategoryParam `form:"category" binding:"required" errField:"category"`
}

type showUsableCategoryParam struct {
	Id int `form:"id" binding:"gt=0,required" errField:"id"`
}

func (s *ShowUsableParam) Bind(c *gin.Context) error {
	return c.ShouldBindJSON(s)
}

func (s *ShowUsableParam) Validate(err error) []*sysError.ValidateError {
	return validator.Validate(err)
}