package initialize

import (
	"afkser/utils"
	"strconv"

	"github.com/go-redis/redis"
)

// NewRedis 在中间件中初始化redis链接
func NewRedis(addr, password, dbselect string) *redis.Client {
	db, _ := strconv.ParseUint(dbselect, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:       addr,
		Password:   password,
		DB:         int(db),
		MaxRetries: 1,
	})

	_, err := client.Ping().Result()

	if err != nil {
		utils.Log().Panic("连接Redis不成功", err)
	}

	return client
}
