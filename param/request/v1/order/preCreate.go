package order

import (
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/lib/validator"
	"WeddingDressManage/param"
	"github.com/gin-gonic/gin"
)

type PreCreateParam struct {
	Dresses []*PreCreateDress `form:"dresses" binding:"gt=0,required,unique,dive" errField:"dresses"`
	Order   *PreCreateOrder   `form:"order" binding:"required" errField:"order"`
}

type PreCreateDress struct {
	Id int `form:"id" binding:"gt=0,required" errField:"id"`
}

type PreCreateOrder struct {
	WeddingDate param.Date `form:"weddingDate" binding:"required" errField:"weddingDate"`
}

func (p *PreCreateParam) Bind(c *gin.Context) error {
	return validator.Bind(p, []interface{}{make([]*PreCreateDress, 0, 0), &PreCreateOrder{}}, c)
}

func (p *PreCreateParam) Validate(errs error) []*sysError.ValidateError {
	return validator.Validate(errs)
}
