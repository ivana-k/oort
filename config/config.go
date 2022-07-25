package config

import (
	"github.com/c12s/oort/config/jetstream"
	"github.com/c12s/oort/config/neo4j"
	"github.com/c12s/oort/config/redis"
	"github.com/c12s/oort/config/server"
)

type Config interface {
	Neo4j() neo4j.Config
	JetStream() jetstream.Config
	Redis() redis.Config
	Server() server.Config
}

type config struct {
	neo4j     neo4j.Config
	jetStream jetstream.Config
	redis     redis.Config
	server    server.Config
}

func NewDefaultConfig() Config {
	return &config{
		neo4j:     neo4j.NewDefaultNeo4jConfig(),
		jetStream: jetstream.NewDefaultJetStreamConfig(),
		redis:     redis.NewRedisConfig(),
		server:    server.NewDefaultServerConfig(),
	}
}

func (c config) Neo4j() neo4j.Config {
	return c.neo4j
}

func (c config) JetStream() jetstream.Config {
	return c.jetStream
}

func (c config) Redis() redis.Config {
	return c.redis
}

func (c config) Server() server.Config {
	return c.server
}
