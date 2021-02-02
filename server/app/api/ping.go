package api

import (
	"afkser/model/response"

	"github.com/gin-gonic/gin"
)

// Ping 测试连通性
func Ping(c *gin.Context) {
	c.JSON(200, response.Response{
		Code: 0,
		Msg:  "Ping",
	})
}
