package conf

import (
	"log"

	"github.com/joho/godotenv"
)

// 初始化env文件
func Init() {
	// 获取环境变量
	err := godotenv.Load("api.env")
	if err != nil {
		log.Fatalf("Load env err: %v", err)
	}
}
