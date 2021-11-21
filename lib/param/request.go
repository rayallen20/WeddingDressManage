package param

import (
	"WeddingDressManage/lib/sysError"
	"github.com/gin-gonic/gin"
)

// RequestParam 定义请求参数结构体的行为
type RequestParam interface {
	// Bind 绑定请求参数至结构体
	Bind(c *gin.Context) error
	// Validate 对已经绑定参数的结构体实例做字段校验
	Validate(err error) []*sysError.ValidateError
}
