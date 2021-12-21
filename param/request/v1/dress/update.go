package dress

import (
	"WeddingDressManage/lib/helper/urlHelper"
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/lib/validator"
	"github.com/gin-gonic/gin"
)

type UpdateParam struct {
	Dress *updateDressParam `form:"dress" binding:"required"`
}

type updateDressParam struct {
	Id           int      `form:"id" binding:"gt=0,required" errField:"id"`
	Size         string   `form:"size" binding:"required,oneof=S M F L XL XXL D" errField:"size"`
	CoverImg     string   `form:"coverImg" binding:"gt=0,required,imgUrl" errField:"coverImg"`
	SecondaryImg []string `form:"secondaryImg" binding:"gte=0,lte=1,imgUrls" errField:"secondaryImg"`
}

func (u *UpdateParam) Bind(c *gin.Context) error {
	return validator.Bind(u, []interface{}{u, &updateDressParam{}}, c)
}

func (u UpdateParam) Validate(err error) []*sysError.ValidateError {
	return validator.Validate(err)
}

func (u *UpdateParam) ExtractUri() {
	u.Dress.CoverImg = urlHelper.GetUriFromWebsite(u.Dress.CoverImg)
	u.Dress.CoverImg = urlHelper.GetUniqueUriFromImgUri(u.Dress.CoverImg)

	secondaryImgUris := make([]string, 0, len(u.Dress.SecondaryImg))
	for _, secondaryImgWebsite := range u.Dress.SecondaryImg {
		secondaryImgUri := urlHelper.GetUriFromWebsite(secondaryImgWebsite)
		secondaryImgUris = append(secondaryImgUris, secondaryImgUri)
	}
	u.Dress.SecondaryImg = secondaryImgUris
	u.Dress.SecondaryImg = urlHelper.GetUniqueUriFromImgUris(u.Dress.SecondaryImg)
}
