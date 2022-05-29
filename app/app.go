package app

import (
	"github.com/c12s/oort/config"
	"github.com/c12s/oort/domain/model"
	"github.com/c12s/oort/store"
	"log"
)

func Run(config config.Config) {
	manager := store.NewManager(config.Neo4j())
	store := store.NewNeo4jPermissionStore(manager)

	resource := model.NewResource()
	//time, err := time.Parse("2006-01-02T15:04:05.000000000[MST]", "2022-05-29T16:01:35.291400226[UTC]")
	//if err != nil {
	//	log.Fatal(err)
	//}
	resource.AddArg("key", "AAAAAAAAAAAAAAAAAAAAAA")
	path := make([]model.Resource, 1)
	path[0] = resource
	//path[1] = resource
	err := store.AddIdentityToPath(resource, path)
	if err != nil {
		log.Fatal(err)
	}
	//err = store.AddIdentity(resource, path)
	//if err != nil {
	//	log.Fatal(err)
	//}
}
