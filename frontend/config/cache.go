package config

import (
	"os"

	"github.com/tsrnd/goweb5/frontend/services/cache"
	"github.com/tsrnd/goweb5/frontend/services/cache/redis"
)

// Cache func
func Cache() cache.Cache {
	return redis.Connect(
		os.Getenv("REDIS_ADDR"),
		os.Getenv("REDIS_PASSWORD"),
		0,
	)
}
