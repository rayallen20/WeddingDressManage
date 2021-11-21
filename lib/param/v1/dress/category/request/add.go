package request

import (
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/lib/validator"
	"github.com/gin-gonic/gin"
)

// AddParam 添加新品类礼服接口请求参数
type AddParam struct {
	Category Category `form:"category" binding:"required" errField:"category"`
	Dress    Dress    `form:"dress" binding:"required" errField:"category"`
}

// Category 添加新品类礼服接口请求参数中 品类信息部分
type Category struct {
	// Kind 品类编码前缀
	Kind string `form:"kind" binding:"gt=0,required" errField:"kind"`
	// SequenceNumber 品类序号
	// SequenceNumber string `form:"sequenceNumber" binding:"gt=0,required" errField:"sequenceNumber" numeric:"true"`
	SequenceNumber string `form:"sequenceNumber" binding:"gt=0,required,numeric" errField:"sequenceNumber"`
	// CharterMoney 租金
	CharterMoney int `form:"charterMoney" binding:"gt=0,required" errField:"charterMoney"`
	// CashPledge 押金
	CashPledge int `form:"cashPledge" binding:"gt=0,required" errField:"cashPledge"`
	// CoverImg 封面图
	CoverImg string `form:"coverImg" binding:"gt=0,required" errField:"coverImg"`
	// SecondaryImg 副图
	SecondaryImg []string `form:"secondaryImg" binding:"gt=0,lte=1" errField:"secondaryImg"`
}

// Dress 添加新品类礼服接口请求参数中 礼服信息部分
type Dress struct {
	DressNumber int `form:"dressNumber" binding:"gt=0,required" errField:"dressNumber"`
	Size string `form:"size" binding:"required,oneof=S M F L XL XXL D" errField:"size"`
}

func (a *AddParam) Bind(c *gin.Context) error {
	return validator.Bind(a, []interface{}{&AddParam{}, &Category{}, &Dress{}}, c)
}

func (a AddParam) Validate(err error) []*sysError.ValidateError {
	return validator.Validate(err)
}