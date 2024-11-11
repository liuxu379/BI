package models

// RiskStrategyList 结构体定义 RiskStrategyList 表的结构
type RiskStrategyList struct {
	ID                 uint      `json:"id" gorm:"primaryKey;column:id"`
	StrategyName       string    `json:"strategy_name" gorm:"column:strategy_name"`
	StrategyBeforeCode string    `json:"strategy_before_code" gorm:"column:strategy_before_code"`
	StrategyAfterCode  string    `json:"strategy_after_code" gorm:"column:strategy_after_code"`
	Status             int       `json:"status" gorm:"column:status"`
	CreatedAt          LocalTime `json:"created_at" gorm:"column:created_at"`
	UpdatedAt          LocalTime `json:"updated_at" gorm:"column:updated_at"`
}
