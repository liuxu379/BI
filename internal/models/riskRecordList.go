package models

type RiskRecordList struct {
	ID                   int       `gorm:"primaryKey;autoIncrement;column:id" json:"id"`                                                       // ID
	PlatformID           int8      `gorm:"column:platform_id;default:1" json:"platform_id"`                                                    // 订单来源
	SerialNumber         string    `gorm:"column:serial_number;size:50;not null;default:''" json:"serial_number"`                              // 流水号
	OrderID              string    `gorm:"column:order_id;size:30;not null" json:"order_id"`                                                   // 订单 ID
	ZmUserID             string    `gorm:"column:zm_user_id;size:50;not null;default:''" json:"zm_user_id"`                                    // 支付宝 ID
	UserName             string    `gorm:"column:user_name;size:20;not null;default:''" json:"user_name"`                                      // 姓名
	UserIDNumber         string    `gorm:"column:user_id_number;size:25;not null;default:''" json:"user_id_number"`                            // 身份证号码
	UserPhone            string    `gorm:"column:user_phone;size:15;not null;default:''" json:"user_phone"`                                    // 手机号码
	UserAddress          string    `gorm:"column:user_address;size:255;not null;default:''" json:"user_address"`                               // 地址
	Type                 int8      `gorm:"column:type;not null;default:0" json:"type"`                                                         // 风控类型
	RequestParams        string    `gorm:"column:request_params;type:text" json:"request_params"`                                              // 请求参数
	RequestResponse      string    `gorm:"column:request_response;type:text" json:"request_response"`                                          // 响应结果
	CreatedAt            LocalTime `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`                             // 创建时间
	UpdatedAt            LocalTime `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"` // 更新时间
	RequestParseResponse string    `gorm:"column:request_parse_response;type:text" json:"request_parse_response"`                              // 解析后的响应结果
	RequestResult        *int8     `gorm:"column:request_result" json:"request_result"`                                                        // 请求结果
}

// TableName sets the insert table name for this struct type
func (RiskRecordList) TableName() string {
	return "risk_record_list"
}
