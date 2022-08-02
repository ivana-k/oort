package app

import (
	"github.com/c12s/oort/poller/async/nats"
	"github.com/c12s/oort/poller/config"
	"github.com/c12s/oort/poller/domain/poller"
	"github.com/c12s/oort/poller/store/outbox/neo4j"
	"log"
)

func Run(config config.Config) {
	manager, err := neo4j.NewTransactionManager(
		config.Neo4j().Uri(),
		config.Neo4j().DbName())
	if err != nil {
		log.Fatal(err)
	}
	outboxStore := neo4j.NewOutboxStore(manager)

	conn, err := nats.NewConnection(config.Nats().Uri())
	if err != nil {
		panic(err)
	}
	publisher, err := nats.NewPublisher(conn)
	if err != nil {
		panic(err)
	}

	p := poller.New(outboxStore, publisher)

	p.Start(config.Poller().IntervalInMs())
}
