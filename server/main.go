package main

import (
	"afkser/initialize"
)

func main() {
	// 初始化几个基础服务
	initialize.Init()

	// gin
	r := initialize.NewRouter()
	r.Run(":3000")
}
