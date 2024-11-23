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
	if err := database.RiskDB.Session(&gorm.Session{PrepareStmt: true}).Where("status = ?", 1).
		Find(&riskStrategyList).Error; err != nil {
		log.Default().Printf("[ERROR] 风控策略查询失败: %v", err)
		utils.ErrorResponse(c, 500, "风控策略查询失败: "+err.Error())
		return
	}
	utils.SuccessResponse(c, riskStrategyList)
}
