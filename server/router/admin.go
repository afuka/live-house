package router

import (
	"afkser/app/admin/users"

	"github.com/gin-gonic/gin"
)

// InitAdminRouter 初始化 api 的路由
func InitAdminRouter(r *gin.RouterGroup) {
	// 获取图形验证码
	r.GET("user/generate-img-captcha", users.GenerateImgCaptcha)
	// 登录
	r.POST("user/login", users.UserLogin)
	// 退出登录
	r.POST("user/logout", users.LoginOut)
	// 获取用户信息
	r.GET("user/info", users.GetUserInfo)
}
