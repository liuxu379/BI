package routes

import (
	"github.com/gin-gonic/gin"
	"risk/internal/controller"
)

func Routes(r *gin.Engine) {
	// 注册路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	r.GET("/risk/getStrategyList", controller.GetStrategyList)
}
