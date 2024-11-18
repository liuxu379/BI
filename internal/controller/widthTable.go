package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"risk/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
)

var db *gorm.DB
var mongoClient *mongo.Client

type ModelInterface interface {
	UpdateOrInsert(orderID string, data map[string]interface{}) error
}

func flattenJson(data interface{}, dataType int, prefix string) map[string]interface{} {
	result := make(map[string]interface{})
	switch v := data.(type) {
	case map[string]interface{}:
		for key, value := range v {
			newKey := key
			if prefix != "" {
				newKey = fmt.Sprintf("%s_%s", prefix, key)
			}
			result[newKey] = flattenJson(value, dataType, newKey)
		}
	case []interface{}:
		for _, item := range v {
			result = flattenJson(item, dataType, prefix)
		}
	default:
		result[prefix] = v
	}
	return result
}

func processNewData() {
	types := [][4]interface{}{
		{18, "Bairong", "bairong_18", "request_response"},
		{19, "Bairong", "bairong_19", "request_response"},
		// 这里添加其他类型...
	}

	for _, item := range types {
		typeID := item[0].(int)
		modelName := item[1].(string)
		prefix := item[2].(string)
		field := item[3].(string)

		yesterdayStart := time.Now().Add(-24 * time.Hour).Truncate(24 * time.Hour)
		yesterdayEnd := yesterdayStart.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

		var records []models.RiskRecordList
		db.Where("type = ? AND created_at BETWEEN ? AND ?", typeID, yesterdayStart, yesterdayEnd).
			Find(&records)

		for _, record := range records {
			// 使用 'field' 获取数据库中存储的 JSON 字段
			data := make(map[string]interface{})
			err := json.Unmarshal([]byte(field), &data)
			if err != nil {
				log.Println("Error unmarshalling field:", err)
				continue
			}

			// 使用 flattenJson 处理数据
			processedData := flattenJson(data, typeID, prefix)
			processedData["order_id"] = record.OrderID
			processedData["user_name"] = record.UserName
			processedData["user_phone"] = record.UserPhone
			processedData["user_id_number"] = record.UserIDNumber
			processedData["created_at"] = record.CreatedAt

			// 使用 modelName 动态选择集合进行插入或更新
			err = updateOrInsertMongoDB(modelName, record.OrderID, processedData)
			if err != nil {
				log.Printf("插入/更新 MongoDB 时出错，订单ID %s: %v", record.OrderID, err)
			}
		}
	}
}

func updateOrInsertMongoDB(modelName string, orderID string, data map[string]interface{}) error {
	// 动态选择 MongoDB 集合
	collection := mongoClient.Database("your_database").Collection(modelName)
	filter := bson.M{"order_id": orderID}
	_, err := collection.UpdateOne(
		context.Background(),
		filter,
		bson.M{"$set": data},
		options.Update().SetUpsert(true),
	)
	return err
}
