package initialize

import (
	"afkser/utils"
	"os"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// DB db
var DB *gorm.DB

// Redis redis
var Redis *redis.Client

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load()

	// 设置日志级别
	utils.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 连接数据库
	DB = NewDatabase(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PATH"), os.Getenv("DB_NAME"), os.Getenv("DB_CONFIG"))

	// redis
	Redis = NewRedis(os.Getenv("REDIS_ADDR"), os.Getenv("REDIS_PW"), os.Getenv("REDIS_DB"))
}
