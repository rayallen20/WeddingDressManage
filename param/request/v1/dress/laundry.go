package dress

import (
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/lib/validator"
	"github.com/gin-gonic/gin"
)

type LaundryParam struct {
	Dress         *dressLaundryParam         `form:"dress" binding:"required" errField:"dress"`
	LaundryDetail *laundryDetailLaundryParam `form:"laundryDetail" binding:"required" errField:"laundryDetail"`
}

type dressLaundryParam struct {
	Id int `form:"id" binding:"gt=0,required" errField:"id"`
}

type laundryDetailLaundryParam struct {
	DirtyPositionImg []string `form:"dirtyPositionImg" binding:"gt=0,lte=2,required,imgUrls" errField:"dirtyPositionImg"`
	Note             string   `form:"note" binding:"gt=0,required" errField:"note"`
}

func (l *LaundryParam) Bind(c *gin.Context) error {
	return c.ShouldBindJSON(l)
}

func (l *LaundryParam) Validate(err error) []*sysError.ValidateError {
	return validator.Validate(err)
}
