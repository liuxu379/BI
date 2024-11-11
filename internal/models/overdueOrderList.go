package models

import "time"

// OverdueOrderList 结构体定义逾期订单数据
type OverdueOrderList struct {
	ID                 uint64  `gorm:"primaryKey"`
	OrderID            string  `gorm:"size:35"`
	ZMUserID           string  `gorm:"size:18"`
	UserPhone          string  `gorm:"size:15"`
	UserIDNumber       string  `gorm:"size:20"`
	OrderAmount        float64 `gorm:"type:decimal(10,2)"`
	OverdueHistoryDays int     `gorm:"type:int"`
	OverdueNowDays     int     `gorm:"type:int"`
	OverdueNowAmount   float64 `gorm:"type:decimal(10,2)"`
	OrderCreatedAt     time.Time
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
