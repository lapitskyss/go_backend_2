package rediesstore

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/cache/v8"

	"github.com/lapitskyss/go_backend_2/internal/srv/securitysrv"
)

func (s *RedisStore) SaveCode(ctx context.Context, key string, value string) error {
	return s.cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: value,
		TTL:   time.Hour,
	})
}

func (s *RedisStore) GetCode(ctx context.Context, key string) (string, error) {
	var code string
	err := s.cache.Get(ctx, key, &code)
	if err != nil {
		if errors.Is(err, cache.ErrCacheMiss) {
			return "", securitysrv.ErrCodeNotFound
		}

		return "", err
	}

	return code, nil
}
