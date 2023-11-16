package container

import (
	"time"

	"git.5th.im/lb-public/gear/log"
	gocache "github.com/patrickmn/go-cache"
	"github.com/samber/do"
)

const localCacheNode = "gocache"

func localCache() *gocache.Cache {
	v, err := do.InvokeNamed[*gocache.Cache](container.injector, localCacheNode)
	if err != nil {
		log.Panicf("LocalCache err:%v", err)
	}

	return v
}

func SetLocalCache[T any](key string, value T, d time.Duration) {
	localCache().Set(key, value, 0)
}

func GetLocalCache[T any](key string) (T, bool) {
	var t T
	v, ok := localCache().Get(key)
	if !ok {
		return t, false
	}
	return v.(T), true
}
