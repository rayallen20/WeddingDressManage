package dress

import (
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/lib/validator"
	"github.com/gin-gonic/gin"
)

type MaintainGiveBackParam struct {
	Maintain *givebackMaintainParam
}

type givebackMaintainParam struct {
	Id int `json:"id" binding:"gt=0,required" errField:"id"`
}

func (m *MaintainGiveBackParam) Bind(c *gin.Context) error {
	return validator.Bind(m, []interface{}{m}, c)
}

func (m *MaintainGiveBackParam) Validate(err error) []*sysError.ValidateError {
	return validator.Validate(err)
}
