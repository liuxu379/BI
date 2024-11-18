package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"risk/database"
	"risk/internal/models"
	"risk/utils"
)

var riskStrategyList []models.RiskStrategyList

// GetStrategyList 获取风控策略列表
func GetStrategyList(c *gin.Context) {
	// 查询风险策略列表
	// if err := database.RiskDB.Session(&gorm.Session{PrepareStmt: true}).
	// 	Where("status = @status", map[string]interface{}{"status": 1}).
	// 	Find(&riskStrategyList).Error; err != nil {
	// 	//记录日志
	// 	log.Default().Printf("[ERROR] 风控策略查询失败: %v", err)
	// 	// 使用统一的错误响应
	// 	utils.ErrorResponse(c, 500, "风控策略查询失败: "+err.Error())
	// 	return
	// }

	// database.RiskDB.Commit()
	// // 返回查询结果
	// utils.SuccessResponse(c, riskStrategyList)

	if err := database.RiskDB.Session(&gorm.Session{PrepareStmt: true}).Where("status = ?", 1).
		Find(&riskStrategyList).Error; err != nil {
		log.Default().Printf("[ERROR] 风控策略查询失败: %v", err)
		utils.ErrorResponse(c, 500, "风控策略查询失败: "+err.Error())
		return
	}
}
