package config

import (
	"github.com/c12s/oort/config/nats"
	"github.com/c12s/oort/config/neo4j"
	"github.com/c12s/oort/config/redis"
	"github.com/c12s/oort/config/server"
)

type Config interface {
	Neo4j() neo4j.Config
	Nats() nats.Config
	Redis() redis.Config
	Server() server.Config
}

type config struct {
	neo4j  neo4j.Config
	nats   nats.Config
	redis  redis.Config
	server server.Config
}

func NewConfig() Config {
	return &config{
		neo4j:  neo4j.NewConfig(),
		nats:   nats.NewConfig(),
		redis:  redis.NewRedisConfig(),
		server: server.NewConfig(),
	}
}

func (c config) Neo4j() neo4j.Config {
	return c.neo4j
}

func (c config) Nats() nats.Config {
	return c.nats
}

func (c config) Redis() redis.Config {
	return c.redis
}

func (c config) Server() server.Config {
	return c.server
}
