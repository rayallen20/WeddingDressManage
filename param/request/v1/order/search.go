package order

import (
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/lib/validator"
	"WeddingDressManage/param"
	"WeddingDressManage/param/request/v1/pagination"
	"github.com/gin-gonic/gin"
)

type SearchParam struct {
	Pagination      *pagination.Pagination `from:"pagination" binding:"required" errField:"pagination"`
	SearchCondition *Condition             `from:"searchCondition" json:"searchCondition" errField:"searchCondition"`
}

type Condition struct {
	// WeddingDate time.Time `form:"weddingDate" time_format:"2006-01-02" errField:"weddingDate"`
	WeddingDate param.Date `form:"weddingDate" json:"weddingDate" errField:"weddingDate"`
}

func (s *SearchParam) Bind(c *gin.Context) error {
	return validator.Bind(s, []interface{}{&pagination.Pagination{}, &Condition{}}, c)
}

func (s *SearchParam) Validate(errs error) (validateErrors []*sysError.ValidateError) {
	return validator.Validate(errs)
}
