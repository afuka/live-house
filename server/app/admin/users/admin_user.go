package users

import (
	"afkser/model"
	"afkser/model/response"
	"afkser/utils"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

// LoginStruct 用户登录的参数
type LoginStruct struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaID string `json:"captcha_id"`
}

// LoginToken 登录的token凭证
type LoginToken struct {
	Token  string `json:"token"`
	UserID int    `json:"user_id"`
}

// UserLogin 登录
func UserLogin(c *gin.Context) {
	var params LoginStruct
	_ = c.ShouldBindJSON(&params)

	// 验证参数
	rules := utils.VerifyRules{
		"Username":  {utils.NotEmpty()},
		"Password":  {utils.NotEmpty()},
		"CaptchaID": {utils.NotEmpty()},
		"Captcha":   {utils.NotEmpty()},
	}
	if err := utils.Verify(params, rules); err != nil {
		response.Result(c, response.ParamsErr, err.Error())
		return
	}

	if !store.Verify(params.CaptchaID, params.Captcha, true) {
		response.Result(c, response.ParamsErr, "验证码错误或超时")
		return
	}
	// 查询用户
	user := model.AdminUsers{}
	err := user.GetUserByUsername(params.Username)
	if err != nil {
		response.Result(c, response.ParamsErr, "用户名不存在或已删除~")
		return
	}
	if !user.CheckPassword(params.Password) {
		response.Result(c, response.ParamsErr, "登陆失败! 用户名不存在或者密码错误~")
		return
	}
	if user.Status != true {
		response.Result(c, response.ParamsErr, "用户已停用~")
		return
	}
	// 生成member_token
	err = user.GenerateToken()
	if err != nil {
		response.Result(c, response.ParamsErr, "管理与签证生成失败~")
		return
	}

	tokenData := LoginToken{
		Token:  user.RememberToken,
		UserID: user.ID,
	}

	response.ResultWithData(c, response.Success, "成功", tokenData)
}

var tokenKey = "X-Token"
var userIDKey = "X-UserId"

// LoginOut 退出登录
func LoginOut(c *gin.Context) {
	token := c.Request.Header.Get(tokenKey)
	userIDStr := c.Request.Header.Get(userIDKey)

	if token == "" || userIDStr == "" {
		response.Result(c, response.AuthErr, "不存在的认证信息~")
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		response.Result(c, response.ParamsErr, "认证信息数据异常~")
		return
	}

	// 查询到当前用户
	user := model.AdminUsers{}
	err = user.GetUserByID(userID)

	if err != nil {
		response.Result(c, response.ParamsErr, "用户名不存在或已删除~")
		return
	}

	if user.RememberToken != token {
		response.Result(c, response.PermissionErr, "无权做退出操作~")
		return
	}
	// 重新生成token
	err = user.GenerateToken()
	if err != nil {
		response.Result(c, response.ServiceErr, "退出失败~")
		return
	}

	response.Result(c, response.Success, "退出成功~")
}

// UserInfo 响应的用户信息
type UserInfo struct {
	Name   string      `json:"name"`
	Avatar string      `json:"avatar"`
	Roles  interface{} `json:"roles"`
}

// GetUserInfo 用户的信息
func GetUserInfo(c *gin.Context) {
	token := c.Request.Header.Get(tokenKey)
	userIDStr := c.Request.Header.Get(userIDKey)

	if token == "" || userIDStr == "" {
		response.Result(c, response.AuthErr, "不存在的认证信息~")
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		response.Result(c, response.ParamsErr, "认证信息数据异常~")
		return
	}

	// 查询到当前用户
	user := model.AdminUsers{}
	err = user.GetUserByID(userID)

	fmt.Println(user)

	if err != nil {
		response.Result(c, response.ParamsErr, "用户名不存在或已删除~")
		return
	}

	if user.RememberToken != token {
		response.Result(c, response.PermissionErr, "无权做退出操作~")
		return
	}

	userExtInfo, err := user.GetExtInfo()
	if err != nil {
		response.Result(c, response.ServiceErr, "拓展信息获取失败~")
		return
	}

	// 查询角色权限
	response.ResultWithData(c, response.Success, "成功", UserInfo{
		Name:   user.Name,
		Avatar: user.Avatar,
		Roles:  userExtInfo.Roles,
	})
}
