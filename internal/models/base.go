package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// LocalTime 自定义时间类型
type LocalTime struct {
	time.Time
}

// MarshalJSON 重写 JSON 序列化格式
func (t LocalTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

// Value 实现 driver.Valuer 接口，用于数据库存储
func (t LocalTime) Value() (driver.Value, error) {
	return t.Time, nil
}

// Scan 实现 sql.Scanner 接口，用于从数据库读取
func (t *LocalTime) Scan(value interface{}) error {
	if v, ok := value.(time.Time); ok {
		*t = LocalTime{Time: v}
		return nil
	}
	return fmt.Errorf("cannot convert %v to LocalTime", value)
}
