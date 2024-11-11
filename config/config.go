package config

import (
	"github.com/joho/godotenv"
	"log"
)

// LoadEnv 加载环境变量
func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("加载 .env 文件失败")
		return err
	}
	return nil
}
