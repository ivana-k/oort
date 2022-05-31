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

	//childPath := make([]model.Resource, 1)
	//childPath[0] = ns1
	//parentPath := make([]model.Resource, 1)
	//parentPath[0] = c
	//err := store.Connect(parentPath, childPath)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//path := make([]model.Resource, 2)
	//path[0] = c
	//path[1] = ns1
	cf1 := model.NewResource()
	cf1.AddArg("id", "config/cf1")
	//err := store.AddResourceToPath(cf1, path)
	//if err != nil {
	//	log.Fatal(err)
	//}

	u := model.NewResource()
	u.AddArg("id", "user/u1")
	//err := store.AddIdentity(u)
	//if err != nil {
	//	log.Fatal(err)
	//}
	ipath := make([]model.Resource, 1)
	rpath := make([]model.Resource, 2)
	ipath[0] = u
	rpath[0] = c
	rpath[1] = ns1
	//err := store.AddPermission(ipath, rpath, model.NewPermission("namespace.list"))
	//if err != nil {
	//	log.Fatal(err)
	//}
	permission, err := store.CheckPermission(ipath, rpath, model.NewPermission("namespace.list"))
	if err != nil {
		log.Fatal(err)
	}
	print(permission, "\n")
	permission, err = store.CheckPermission(ipath, rpath, model.NewPermission("namespace.create"))
	if err != nil {
		log.Fatal(err)
	}
	print(permission, "\n")

	rpath2 := make([]model.Resource, 3)
	rpath2[0] = c
	rpath2[1] = ns1
	rpath2[2] = cf1
	permission, err = store.CheckPermission(ipath, rpath2, model.NewPermission("namespace.list"))
	if err != nil {
		log.Fatal(err)
	}
	print(permission, "\n")
}
