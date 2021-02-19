package facades

import "github.com/go-redis/redis"

// Redis Redis缓存客户端单例
var Redis *redis.Client
