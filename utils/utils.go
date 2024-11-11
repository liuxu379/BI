package utils

import "github.com/gin-gonic/gin"

func ErrorResponse(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{"msg": msg, "code": code})
}

// SuccessResponse 统一的成功响应
func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{"msg": "操作成功", "data": data, "code": 200})
}
