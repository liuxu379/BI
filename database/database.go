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
