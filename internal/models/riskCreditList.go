package models

import "time"

type RiskCreditList struct {
	ID              int64     `gorm:"primaryKey"`
	OrderID         string    `gorm:"column:order_id"`
	ZMUserID        string    `gorm:"column:zm_user_id"`
	UserPhone       string    `gorm:"column:user_phone"`
	UserIDNumber    string    `gorm:"column:user_id_number"`
	Type            int8      `gorm:"column:type"`
	Status          int8      `gorm:"column:status"`
	Params          string    `gorm:"column:params;type:json"`
	Response        string    `gorm:"column:response;type:json"`
	Remark          string    `gorm:"column:remark"`
	AlipayAuthInfo  string    `gorm:"column:alipay_auth_info;type:json"`
	AlipayAuthLevel *int8     `gorm:"column:alipay_auth_level"`
	AlipayAuthScore *string   `gorm:"column:alipay_auth_score"`
	RiskLevel       *string   `gorm:"column:risk_level"`
	RiskScore       *string   `gorm:"column:risk_score"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
}
