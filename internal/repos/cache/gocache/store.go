package gocache

//import (
//	"context"
//	"errors"
//	"github.com/allegro/bigcache/v3"
//	"github.com/c12s/oort/internal/services"
//	gocache "github.com/eko/gocache/v2/cache"
//	"github.com/eko/gocache/v2/repos"
//	"github.com/go-redis/redis/v8"
//	"log"
//	"time"
//)
//
//type GoCache struct {
//	manager        *gocache.ChainCache
//	redisClient    *redis.Client
//	bigcacheClient *bigcache.BigCache
//}
//
//func NewGoCache(redisAddress string, localEviction time.Duration) (services.Cache, error) {
//	bigcacheClient, _ := bigcache.NewBigCache(bigcache.DefaultConfig(localEviction * time.Minute))
//	bigcacheStore := repos.NewBigcache(bigcacheClient, nil)
//	redisClient := redis.NewClient(&redis.Options{Addr: redisAddress})
//	redisStore := repos.NewRedis(redisClient, nil)
//
//	cacheManager := gocache.NewChain(
//		gocache.New(bigcacheStore),
//		gocache.New(redisStore),
//	)
//
//	if cacheManager == nil {
//		return nil, errors.New("cache could not be initialized")
//	}
//	return GoCache{
//		manager:        cacheManager,
//		redisClient:    redisClient,
//		bigcacheClient: bigcacheClient,
//	}, nil
//}
//
//func (g GoCache) Get(key string) ([]byte, error) {
//	value, err := g.manager.Get(context.TODO(), key)
//	if err != nil {
//		return nil, err
//	}
//	return value.([]byte), nil
//}
//
//func (g GoCache) Set(key string, value []byte, tags []string) error {
//	return g.manager.Set(context.TODO(), key, value, &repos.Options{
//		Tags: tags,
//	})
//}
//
//func (g GoCache) Invalidate(tags []string) error {
//	return g.manager.Invalidate(context.TODO(), repos.InvalidateOptions{
//		Tags: tags,
//	})
//}
//
//func (g GoCache) Stop() {
//	err := g.bigcacheClient.Close()
//	if err != nil {
//		log.Println(err)
//	}
//	err = g.redisClient.Close()
//	if err != nil {
//		log.Println(err)
//	}
//}
