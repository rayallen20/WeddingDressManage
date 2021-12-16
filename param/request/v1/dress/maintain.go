package dress

import (
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/lib/validator"
	"github.com/gin-gonic/gin"
)

type MaintainParam struct {
	Dress          *dressMaintainParam          `form:"dress" binding:"required" errField:"dress"`
	MaintainDetail *maintainDetailMaintainParam `form:"maintainDetail" binding:"required" errField:"maintainDetail"`
}

type dressMaintainParam struct {
	Id int `form:"id" binding:"gt=0,required" errField:"id"`
}

type maintainDetailMaintainParam struct {
	MaintainPositionImg []string `form:"maintainPositionImg" binding:"gt=0,lte=2,required,imgUrls" errField:"maintainPositionImg"`
	Note                string   `form:"note" binding:"gt=0,required" errField:"note"`
}

func (m *MaintainParam) Bind(c *gin.Context) error {
	return c.ShouldBindJSON(m)
}

func (m *MaintainParam) Validate(err error) []*sysError.ValidateError {
	return validator.Validate(err)
}
