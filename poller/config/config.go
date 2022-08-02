package config

import (
	"github.com/c12s/oort/poller/config/nats"
	"github.com/c12s/oort/poller/config/neo4j"
)

type Config interface {
	Neo4j() neo4j.Config
	Nats() nats.Config
}

type config struct {
	neo4j neo4j.Config
	nats  nats.Config
}

func NewConfig() Config {
	return &config{
		neo4j: neo4j.NewConfig(),
		nats:  nats.NewConfig(),
	}
}

func (c config) Neo4j() neo4j.Config {
	return c.neo4j
}

func (c config) Nats() nats.Config {
	return c.nats
}
