package api

import (
	"afkser/model/response"

	"github.com/gin-gonic/gin"
)

// Ping 测试连通性
func Ping(c *gin.Context) {
	response.Result(c, response.Success, "success")
}
