package initialize

import (
	"afkser/utils"
	"os"

	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load()

	// 设置日志级别
	utils.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 连接数据库
	mysqlDsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_PATH") + ")/" + os.Getenv("DB_NAME") + "?" + os.Getenv("DB_CONFIG")
	Database(mysqlDsn)

	// redis
	Redis()
}
