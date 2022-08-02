package config

import (
	"github.com/c12s/oort/poller/config/nats"
	"github.com/c12s/oort/poller/config/neo4j"
	"github.com/c12s/oort/poller/config/poller"
)

type Config interface {
	Neo4j() neo4j.Config
	Nats() nats.Config
	Poller() poller.Config
}

type config struct {
	neo4j  neo4j.Config
	nats   nats.Config
	poller poller.Config
}

func NewConfig() Config {
	return &config{
		neo4j:  neo4j.NewConfig(),
		nats:   nats.NewConfig(),
		poller: poller.NewConfig(),
	}
}

func (c config) Neo4j() neo4j.Config {
	return c.neo4j
}

func (c config) Nats() nats.Config {
	return c.nats
}

func (c config) Poller() poller.Config {
	return c.poller
}
