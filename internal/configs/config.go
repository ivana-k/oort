package configs

import (
	"github.com/c12s/oort/internal/configs/nats"
	"github.com/c12s/oort/internal/configs/neo4j"
	"github.com/c12s/oort/internal/configs/server"
)

type Config interface {
	Neo4j() neo4j.Config
	Nats() nats.Config
	Server() server.Config
}

type config struct {
	neo4j  neo4j.Config
	nats   nats.Config
	server server.Config
}

func NewConfig() (Config, error) {
	return &config{
		neo4j:  neo4j.NewConfig(),
		nats:   nats.NewConfig(),
		server: server.NewConfig(),
	}, nil
}

func (c config) Neo4j() neo4j.Config {
	return c.neo4j
}

func (c config) Nats() nats.Config {
	return c.nats
}

func (c config) Server() server.Config {
	return c.server
}
