package dress

import (
	"WeddingDressManage/lib/helper/urlHelper"
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/lib/validator"
	"github.com/gin-gonic/gin"
)

type AddParam struct {
	Category *Category `form:"category" errField:"category"`
	Dress *Dress `form:"dress" errField:"dress"`
}

// Category 品类信息参数
type Category struct {
	// Id 品类ID
	Id int `form:"id" binding:"gt=0,required" errField:"id"`
}

// Dress 礼服信息参数
type Dress struct {
	// Size 尺码
	Size string `form:"size" binding:"required,oneof=S M F L XL XXL D" errField:"size"`
	// Number 礼服数量
	Number int `form:"number" binding:"gt=0,required" errField:"number"`
	// CoverImg 封面图
	CoverImg string `form:"coverImg" binding:"gt=0,required,imgUrl" errField:"coverImg"`
	// SecondaryImg 副图
	SecondaryImg []string `form:"secondaryImg" binding:"gt=0,lte=1,imgUrls" errField:"secondaryImg"`
}

func (a *AddParam) Bind(c *gin.Context) error {
	return validator.Bind(a, []interface{}{a, &Category{}, &Dress{}}, c)
}

func (a *AddParam) Validate(err error) []*sysError.ValidateError {
	return validator.Validate(err)
}

// ExtractUri 将校验参数中的封面图和副图的网址转化为uri
func (a *AddParam) ExtractUri() {
	a.Dress.CoverImg = urlHelper.GetUriFromWebsite(a.Dress.CoverImg)
	secondaryImgUris := make([]string, 0, len(a.Dress.SecondaryImg))
	for _, secondaryImgWebsite := range a.Dress.SecondaryImg {
		secondaryImgUri := urlHelper.GetUriFromWebsite(secondaryImgWebsite)
		secondaryImgUris = append(secondaryImgUris, secondaryImgUri)
	}
	a.Dress.SecondaryImg = secondaryImgUris
}