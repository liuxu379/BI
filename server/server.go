package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"risk/config"
	"risk/database"
	"risk/middleware"
	"risk/routes"
)

// StartServer 启动 Web 服务
func StartServer() {
	// 加载配置文件
	if err := config.LoadEnv(); err != nil {
		log.Fatalf("配置加载失败: %v", err)
	}

	// 初始化数据库连接
	err := database.ConnectApiDB()
	if err != nil {
		log.Fatalf("API 数据库连接失败: %v", err)
	}

	err = database.ConnectRiskDB()
	if err != nil {
		log.Fatalf("Risk 数据库连接失败: %v", err)
	}

	// 创建 Gin 引擎
	r := gin.Default()

	// 注册日志中间件
	r.Use(middleware.LoggerMiddleware())

	// 设置路由
	routes.Routes(r)

	// 启动服务器
	err = r.Run(":8080")
	if err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
