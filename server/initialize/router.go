package initialize

import (
	"afkser/middleware"
	"afkser/router"
	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {

	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())

	// 路由
	api := r.Group("/api")
	{
		router.InitAPIRouter(api)
	}

	// 管理后台
	admin := r.Group("/admin")
	{
		router.InitAdminRouter(admin)
	}

	return r
}
