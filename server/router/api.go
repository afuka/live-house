package router

import (
	"afkser/app/api"

	"github.com/gin-gonic/gin"
)

// InitAPIRouter 初始化 api 的路由
func InitAPIRouter(r *gin.RouterGroup) {
	r.GET("hello", api.Ping)
}
