package config

import (
	"github.com/leandro-andrade-candido/api-go/src/libs/cache"
)

var redisCache cache.Cache

func GetCache() cache.Cache {
	return redisCache
}

func initializeCache() {
	redisAddress := GetString("database.redis.address")
	redisCache = cache.NewRedisCache(
		redisAddress,
		GetString("database.redis.password"),
		GetInt("database.redis.db"),
	)
}
