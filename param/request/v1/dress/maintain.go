package dress

import (
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/lib/validator"
	"github.com/gin-gonic/gin"
)

type MaintainParam struct {
	Dress          *MaintainDressParam          `form:"dress" binding:"required" errField:"dress"`
	MaintainDetail *MaintainDetailMaintainParam `form:"maintainDetail" binding:"required" errField:"maintainDetail"`
}

type MaintainDressParam struct {
	Id int `form:"id" binding:"gt=0,required" errField:"id"`
}

type MaintainDetailMaintainParam struct {
	MaintainPositionImg []string `form:"maintainPositionImg" binding:"gt=0,lte=2,required,imgUrls" errField:"maintainPositionImg"`
	Note                string   `form:"note" binding:"gt=0,required" errField:"note"`
}

func (m *MaintainParam) Bind(c *gin.Context) error {
	return validator.Bind(m, []interface{}{m, &MaintainDressParam{}, &MaintainDetailMaintainParam{}}, c)
}

func (m *MaintainParam) Validate(err error) []*sysError.ValidateError {
	return validator.Validate(err)
}
