package rediesstore

import (
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

type RedisStore struct {
	cache *cache.Cache
}

func InitRedisStore(url string) *RedisStore {
	rdb := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: "",
		DB:       0,
	})

	c := cache.New(&cache.Options{
		Redis:      rdb,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	return &RedisStore{
		cache: c,
	}
}
