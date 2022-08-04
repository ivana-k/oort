package gocache

import (
	"context"
	"errors"
	"fmt"
	"github.com/allegro/bigcache/v3"
	"github.com/c12s/oort/domain/store/cache"
	gocache "github.com/eko/gocache/v2/cache"
	"github.com/eko/gocache/v2/store"
	"github.com/go-redis/redis/v8"
	"time"
)

type GoCache struct {
	manager *gocache.ChainCache
}

func NewGoCache(redisAddress string, localEviction time.Duration) (cache.Cache, error) {
	bigcacheClient, _ := bigcache.NewBigCache(bigcache.DefaultConfig(localEviction * time.Minute))
	bigcacheStore := store.NewBigcache(bigcacheClient, nil)
	redisStore := store.NewRedis(redis.NewClient(&redis.Options{
		Addr: redisAddress,
	}), nil)

	cacheManager := gocache.NewChain(
		gocache.New(bigcacheStore),
		gocache.New(redisStore),
	)

	if cacheManager == nil {
		return nil, errors.New("cache could not be initialized")
	}
	return GoCache{manager: cacheManager}, nil
}

func (g GoCache) Get(key string) ([]byte, error) {
	value, err := g.manager.Get(context.TODO(), key)
	if err != nil {
		return nil, err
	}
	fmt.Println("VALUE --------")
	fmt.Println(value)
	return value.([]byte), nil
}

func (g GoCache) Set(key string, value []byte, tags []string) error {
	return g.manager.Set(context.TODO(), key, value, &store.Options{
		Tags: tags,
	})
}

func (g GoCache) Invalidate(tags []string) error {
	return g.manager.Invalidate(context.TODO(), store.InvalidateOptions{
		Tags: tags,
	})
}
