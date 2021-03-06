package dress

import (
	"WeddingDressManage/lib/helper/urlHelper"
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/lib/validator"
	"github.com/gin-gonic/gin"
)

type LaundryParam struct {
	Dress         *LaundryDressParam         `form:"dress" binding:"required" errField:"dress"`
	LaundryDetail *LaundryLaundryDetailParam `form:"laundryDetail" binding:"required" errField:"laundryDetail"`
}

type LaundryDressParam struct {
	Id int `form:"id" binding:"gt=0,required" errField:"id"`
}

type LaundryLaundryDetailParam struct {
	DirtyPositionImg []string `form:"dirtyPositionImg" binding:"gt=0,lte=2,required,imgUrls" errField:"dirtyPositionImg"`
	Note             string   `form:"note" errField:"note"`
}

func (l *LaundryParam) Bind(c *gin.Context) error {
	return validator.Bind(l, []interface{}{l, &LaundryDressParam{}, &LaundryLaundryDetailParam{}}, c)
}

func (l *LaundryParam) Validate(err error) []*sysError.ValidateError {
	return validator.Validate(err)
}

func (l *LaundryParam) ExtractUri() {
	dirtyPositionUris := make([]string, 0, len(l.LaundryDetail.DirtyPositionImg))
	for _, dirtyPositionImg := range l.LaundryDetail.DirtyPositionImg {
		dirtyPositionImgUri := urlHelper.GetUriFromWebsite(dirtyPositionImg)
		dirtyPositionUris = append(dirtyPositionUris, dirtyPositionImgUri)
	}
	l.LaundryDetail.DirtyPositionImg = urlHelper.GetUniqueUriFromImgUris(dirtyPositionUris)
}
