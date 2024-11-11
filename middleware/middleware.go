package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// LoggerMiddleware 记录请求日志
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 记录日志
		log.Printf("请求路径: %s | 状态码: %d | 耗时: %v", c.Request.URL.Path, c.Writer.Status(), time.Since(startTime))
	}
}
