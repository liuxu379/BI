package controller

import (
	"github.com/gin-gonic/gin"
	"risk/database"
	"risk/internal/models"
	"risk/utils"
)

// GetStrategyList 获取风控策略列表
func GetStrategyList(c *gin.Context) {
	// 查询风险策略列表
	var riskStrategyList []models.RiskStrategyList
	if err := database.RiskDB.
		Where("status = ?", 1).
		Find(&riskStrategyList).Error; err != nil {
		// 使用统一的错误响应
		utils.ErrorResponse(c, 500, "风控策略查询失败: "+err.Error())
		return
	}

	// 返回查询结果
	utils.SuccessResponse(c, riskStrategyList)
}

func AddStrategy(c *gin.Context) {
	// 获取参数
	data := []*models.RiskStrategyList{
		{
			StrategyName: "风控策略1",
		},
		{
			StrategyName: "风控策略2",
		},
	}
	result := database.RiskDB.Create(&data)
	if result.Error != nil {
		utils.ErrorResponse(c, 500, "风控策略添加失败: "+result.Error.Error())
		return
	}
}
