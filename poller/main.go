package main

import (
	"github.com/c12s/oort/poller/async/nats"
	"github.com/c12s/oort/poller/config"
	"github.com/c12s/oort/poller/domain/poller"
	"github.com/c12s/oort/poller/store/outbox/neo4j"
	"log"
)

func main() {
	conf := config.NewConfig()

	manager, err := neo4j.NewTransactionManager(
		conf.Neo4j().Uri(),
		conf.Neo4j().DbName())
	if err != nil {
		log.Fatal(err)
	}
	outboxStore := neo4j.NewOutboxStore(manager)

	conn, err := nats.NewConnection(conf.Nats().Uri())
	if err != nil {
		panic(err)
	}
	publisher, err := nats.NewPublisher(conn)
	if err != nil {
		panic(err)
	}

	p := poller.New(outboxStore, publisher)

	p.Start(conf.Poller().IntervalInMs())
}
