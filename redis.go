package container

import (
	"context"
	"time"

	"git.5th.im/lb-public/gear/cache"
	"git.5th.im/lb-public/gear/log"
	"github.com/samber/do"
)

const (
	redisNode = "Redis"
	cacheNode = "Cache"
)

// GetRedis get redis
func GetRedis(ctx context.Context) *cache.RedisClient {
	v, err := do.InvokeNamed[*cache.RedisClient](container.injector, redisNode)
	if err != nil {
		log.Panicf("GetRedis err:%v", err)
	}
	return v.WithContext(ctx)
}

// GetCache get cache
func GetCache() *cache.Cache {
	v, err := do.InvokeNamed[*cache.Cache](container.injector, cacheNode)
	if err != nil {
		log.Panicf("GetCache err:%v", err)
	}
	return v
}

// Fetch fetches a value from cache, if not found, call blockCall to get the value and set it to cache
func Fetch[T any](ctx context.Context, key string, expiration time.Duration, blockCall func() (rawResult T, err error)) (v *T, err error) {

	var result T
	ok, err := GetCache().Fetch(ctx, key, &result, expiration, func() (rawResult any, err error) {
		return blockCall()
	})
	if err != nil {
		return
	}
	if !ok {
		return nil, nil
	}

	return &result, nil
}
