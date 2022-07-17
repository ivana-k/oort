package app

import (
	"github.com/c12s/oort/config"
	"github.com/c12s/oort/domain/model"
	"github.com/c12s/oort/store/acl/neo4j"
	"log"
)

func Run(config config.Config) {
	manager, err := neo4j.NewTransactionManager(
		config.Neo4j().Uri(),
		config.Neo4j().Username(),
		config.Neo4j().Password(),
		config.Neo4j().DbName())
	if err != nil {
		log.Fatal(err)
	}
	store := neo4j.NewAclStore(manager)
	parent := model.NewResource("abcd", "cluster")
	child := model.NewResource("efgh", "namespace")
	gchild := model.NewResource("aaa", "secrets")

	err = store.ConnectResources(parent, child)
	if err != nil {
		log.Fatal(err)
	}
	err = store.ConnectResources(child, gchild)
	attr := model.NewAttribute("owner", model.String, []byte("pera"))
	err = store.UpsertAttribute(gchild, attr)
	if err != nil {
		log.Fatal(err)
	}
	//err = store.RemoveAttribute(parent, attr)
	//if err != nil {
	//	log.Fatal(err)
	//}
	err = store.DisconnectResources(child, gchild)
	if err != nil {
		log.Fatal(err)
	}
}
