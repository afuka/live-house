package router

import (
	"afkser/app/api"

	"github.com/gin-gonic/gin"
)

// InitAdminRouter 初始化 api 的路由
func InitAdminRouter(r *gin.RouterGroup) {
	r.GET("hello", api.Ping)
}
