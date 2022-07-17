package config

import (
	"github.com/c12s/oort/config/kafka"
	"github.com/c12s/oort/config/neo4j"
	"github.com/c12s/oort/config/redis"
)

type Config interface {
	Neo4j() neo4j.Config
	Kafka() kafka.Config
	Redis() redis.Config
}

type config struct {
	neo4j neo4j.Config
	kafka  kafka.Config
	redis redis.Config
}

func NewDefaultConfig() Config {
	return &config{
		neo4j: neo4j.NewDefaultNeo4jConfig(),
		kafka:  kafka.NewKafkaConfig(),
		redis: redis.NewRedisConfig(),
	}
}

func (c config) Neo4j() neo4j.Config {
	return c.neo4j
}

func (c config) Kafka() kafka.Config {
	return c.kafka
}

func (c config) Redis() redis.Config {
	return c.redis
}
