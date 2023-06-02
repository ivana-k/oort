package redis

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config interface {
	Address() string
	Eviction() time.Duration
}

func NewRedisConfig() Config {
	hostname := os.Getenv("REDIS_HOSTNAME")
	port := os.Getenv("REDIS_PORT")
	eviction, err := strconv.Atoi(os.Getenv("CACHE_EVICTION_MIN"))
	if err != nil {
		eviction = 10
	}
	return config{
		hostname: hostname,
		port:     port,
		eviction: time.Duration(eviction),
	}
}

type config struct {
	hostname string
	port     string
	eviction time.Duration
}

func (c config) Address() string {
	return fmt.Sprintf("%s:%s", c.hostname, c.port)
}

func (c config) Eviction() time.Duration {
	return c.eviction
}
