package logInterface

import "github.com/gin-gonic/gin"

// SysLog 系统日志接口
type SysLog interface {
	// GetData 获取请求参数
	GetData(c *gin.Context)
	// Logger 记录日志
	Logger()
}
