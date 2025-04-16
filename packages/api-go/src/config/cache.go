package config

import (
	"github.com/leandro-andrade-candido/api-go/src/libs/cache"
)

var redisCache cache.Cache

func init() {
	redisCache = cache.NewRedisCache(
		GetString("database.redis.host"),
		GetString("database.redis.password"),
		GetInt("database.redis.db"),
	)
}

func GetCache() cache.Cache {
	return redisCache
}
