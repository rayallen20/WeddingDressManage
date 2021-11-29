package category

import (
	"WeddingDressManage/lib/helper/urlHelper"
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/lib/validator"
	"github.com/gin-gonic/gin"
)

type UpdateParam struct {
	Category *updateCategoryParam `form:"category" binding:"required" errField:"required"`
}

type updateCategoryParam struct {
	Id int `form:"id" binding:"gt=0,required" errField:"id"`
	SerialNumber string `form:"serialNumber" binding:"gt=0,required" errField:"serialNumber"`
	CharterMoney int `form:"charterMoney" binding:"gt=0,required" errField:"charterMoney"`
	CashPledge int `form:"cashPledge" binding:"gt=0,required" errField:"charterMoney"`
	CoverImg string `form:"coverImg" binding:"gt=0,required,imgUrl" errField:"coverImg"`
	SecondaryImg []string `form:"secondaryImg" binding:"gte=0,lte=1,imgUrls" errField:"secondaryImg"`
}

func (u *UpdateParam) Bind(c *gin.Context) error {
	return c.ShouldBindJSON(u)
}

func (u *UpdateParam) Validate(err error) []*sysError.ValidateError {
	return validator.Validate(err)
}

func (u *UpdateParam) ExtractUri() {
	u.Category.CoverImg = urlHelper.GetUriFromWebsite(u.Category.CoverImg)
	u.Category.CoverImg = urlHelper.GetUniqueUriFromImgUri(u.Category.CoverImg)
	secondaryImgUris := make([]string, 0, len(u.Category.SecondaryImg))
	for _, secondaryImgWebsite := range u.Category.SecondaryImg {
		secondaryImgUri := urlHelper.GetUriFromWebsite(secondaryImgWebsite)
		secondaryImgUris = append(secondaryImgUris, secondaryImgUri)
	}
	u.Category.SecondaryImg = secondaryImgUris
	u.Category.SecondaryImg = urlHelper.GetUniqueUriFromImgUris(u.Category.SecondaryImg)
}