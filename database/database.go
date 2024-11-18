package database

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// 定义两个数据库的全局连接变量
var (
	ApiDB  *gorm.DB
	RiskDB *gorm.DB
)

// ConnectApiDB 连接 API 数据库
func ConnectApiDB() error {
	dsn := os.Getenv("API_DSN")
	if dsn == "" {
		log.Println("API_DSN 环境变量未设置")
		return fmt.Errorf("API_DSN 环境变量未设置")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		PrepareStmt: true,
	})
	if err != nil {
		log.Printf("API 数据库连接失败: %v", err)
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("获取 API 数据库连接失败: %v", err)
		return err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	ApiDB = db
	log.Println("API 数据库连接成功")
	return nil
}

// ConnectRiskDB 连接 Risk 数据库
func ConnectRiskDB() error {
	dsn := os.Getenv("RISK_DSN")
	if dsn == "" {
		log.Println("RISK_DSN 环境变量未设置")
		return fmt.Errorf("RISK_DSN 环境变量未设置")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		PrepareStmt: true,
	})
	if err != nil {
		log.Printf("Risk 数据库连接失败: %v", err)
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("获取 Risk 数据库连接失败: %v", err)
		return err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	RiskDB = db
	log.Println("Risk 数据库连接成功")
	return nil
}

func ConnectSmartDB() error {
	dsn := os.Getenv("SMART_DSN")
	if dsn == "" {
		log.Println("SMART_DSN 环境变量未设置")
		return fmt.Errorf("SMART_DSN 环境变量未设置")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Printf("Smart 数据库连接失败: %v", err)
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("获取 Smart 数据库连接失败: %v", err)
		return err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	RiskDB = db
	log.Println("Smart 数据库连接成功")
	return nil
}

//func ConnectMongoDB() {
//	var err error
//	// 设置MongoDB连接，需要提供用户名、密码、数据库和认证机制
//	clientOptions := options.Client().
//		ApplyURI("mongodb://risk:fub9ehk9yne2ZCY*kgc@8.134.56.190:27017/") // 替换为实际的用户名和密码
//	// 也可以添加认证数据库，如果与目标数据库不同
//	clientOptions.SetAuth(options.Credential{
//		Username:   "risk",                // MongoDB 用户名
//		Password:   "fub9ehk9yne2ZCY*kgc", // MongoDB 密码
//		AuthSource: "admin",
//	})
//	mongoClient, err = mongo.Connect(context.Background(), clientOptions)
//	if err != nil {
//		log.Fatal("连接MongoDB失败:", err)
//	}
//
//	// 检查MongoDB连接
//	err = mongoClient.Ping(context.Background(), nil)
//	if err != nil {
//		log.Fatal("MongoDB ping失败:", err)
//	}
//}
