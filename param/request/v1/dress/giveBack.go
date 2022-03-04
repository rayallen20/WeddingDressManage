package dress

import (
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/lib/validator"
	"github.com/gin-gonic/gin"
)

type GiveBackParam struct {
	Laundry laundryParam `json:"laundry" binding:"required" errField:"laundry"`
}

type laundryParam struct {
	Id int `json:"id" binding:"gt=0,required" errField:"id"`
}

func (g *GiveBackParam) Bind(c *gin.Context) error {
	return validator.Bind(g, []interface{}{g, &laundryParam{}}, c)
}

func (g GiveBackParam) Validate(err error) []*sysError.ValidateError {
	return validator.Validate(err)
}
