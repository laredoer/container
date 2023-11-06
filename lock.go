package container

import (
	"context"
	"time"
)

const (
	prefxi = "lock."
)

// Lock 锁定操作
func Lock(ctx context.Context, key string, duration time.Duration) (success bool, err error) {
	success, err = GetRedis(ctx).SetNX(ctx, prefxi+key, 1, duration).Result()
	return
}

// Unlock 解锁
func Unlock(ctx context.Context, key string) (err error) {
	return GetRedis(ctx).Del(ctx, prefxi+key).Err()
}
