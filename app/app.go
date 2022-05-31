package app

import (
	"github.com/c12s/oort/config"
	"github.com/c12s/oort/domain/model"
	"github.com/c12s/oort/store/neo4j"
	"log"
)

func Run(config config.Config) {
	manager := neo4j.NewTransactionHandler(config.Neo4j())
	store := neo4j.NewNeo4jPermissionStore(manager)

	//time, err := time.Parse("2006-01-02T15:04:05.000000000[MST]", "2022-05-29T16:01:35.291400226[UTC]")
	//if err != nil {
	//	log.Fatal(err)
	//}

	c := model.NewResource()
	c.AddArg("id", "cluster/c1")
	//err := store.AddResource(c)
	//if err != nil {
	//	log.Fatal(err)
	//}

	ns1 := model.NewResource()
	ns1.AddArg("id", "namespace/ns1")
	//err = store.AddResource(ns1)
	//if err != nil {
	//	log.Fatal(err)
	//}

	childPath := make([]model.Resource, 1)
	childPath[0] = ns1
	parentPath := make([]model.Resource, 1)
	parentPath[0] = c
	err := store.Connect(parentPath, childPath)
	if err != nil {
		log.Fatal(err)
	}
}
