package category

import (
	"WeddingDressManage/lib/helper/urlHelper"
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/lib/validator"
	"github.com/gin-gonic/gin"
)

// AddParam 添加新品类礼服接口请求参数
type AddParam struct {
	Kind *addKindParam         `form:"kind" binding:"required" errField:"category"`
	Category *addCategoryParam `form:"category" binding:"required" errField:"category"`
	Dress    *addDressParam    `form:"dress" binding:"required" errField:"category"`
}

// addKindParam 添加新品类礼服接口请求参数中 礼服大类信息部分
type addKindParam struct {
	// Id 礼服大类信息Id
	Id int `form:"id" binding:"gt=0,required" errField:"id"`
}

// addCategoryParam 添加新品类礼服接口请求参数中 品类信息部分
type addCategoryParam struct {
	// SequenceNumber 品类序号
	SequenceNumber string `form:"sequenceNumber" binding:"gt=0,required,numeric" errField:"sequenceNumber"`
	// CharterMoney 租金
	CharterMoney int `form:"charterMoney" binding:"gt=0,required" errField:"charterMoney"`
	// CashPledge 押金
	CashPledge int `form:"cashPledge" binding:"gt=0,required" errField:"cashPledge"`
	// CoverImg 封面图
	CoverImg string `form:"coverImg" binding:"gt=0,required,imgUrl" errField:"coverImg"`
	// SecondaryImg 副图
	SecondaryImg []string `form:"secondaryImg" binding:"gte=0,lte=1,imgUrls" errField:"secondaryImg"`
}

// addDressParam 添加新品类礼服接口请求参数中 礼服信息部分
type addDressParam struct {
	Number int `form:"number" binding:"gt=0,required" errField:"number"`
	Size string `form:"size" binding:"required,oneof=S M F L XL XXL D" errField:"size"`
}

func (a *AddParam) Bind(c *gin.Context) error {
	return validator.Bind(a, []interface{}{&AddParam{}, &addCategoryParam{}, &addDressParam{}}, c)
}

func (a *AddParam) Validate(err error) []*sysError.ValidateError {
	return validator.Validate(err)
}

// ExtractUri 将校验参数中的封面图和副图的网址转化为uri
func (a *AddParam) ExtractUri() {
	a.Category.CoverImg = urlHelper.GetUriFromWebsite(a.Category.CoverImg)
	a.Category.CoverImg = urlHelper.GetUniqueUriFromImgUri(a.Category.CoverImg)
	secondaryImgUris := make([]string, 0, len(a.Category.SecondaryImg))
	for _, secondaryImgWebsite := range a.Category.SecondaryImg {
		secondaryImgUri := urlHelper.GetUriFromWebsite(secondaryImgWebsite)
		secondaryImgUris = append(secondaryImgUris, secondaryImgUri)
	}
	a.Category.SecondaryImg = secondaryImgUris
	a.Category.SecondaryImg = urlHelper.GetUniqueUriFromImgUris(a.Category.SecondaryImg)
}