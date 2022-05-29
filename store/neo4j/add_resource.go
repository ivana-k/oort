package neo4j

import (
	"fmt"
	"github.com/c12s/oort/domain/model"
)

func (store *permissionStore) AddResource(resource model.Resource) error {
	return store.addNode(resource, store.buildResourcePattern())
}

func (store *permissionStore) AddIdentity(resource model.Resource) error {
	return store.addNode(resource, store.buildIdentityPattern())
}

func (store *permissionStore) AddResourceToPath(resource model.Resource, path []model.Resource) error {
	return store.addNodeToPath(resource, path, store.buildResourcePattern())
}

func (store *permissionStore) AddIdentityToPath(resource model.Resource, path []model.Resource) error {
	return store.addNodeToPath(resource, path, store.buildIdentityPattern())
}

func (store *permissionStore) addNode(resource model.Resource, pattern string) error {
	properties, params := store.buildResourceOrIdentityProperties(resource)
	cypher := fmt.Sprintf("CREATE %s SET %s", pattern, properties)
	_, err := store.write(cypher, params)
	return err
}

func (store *permissionStore) addNodeToPath(resource model.Resource, path []model.Resource, pattern string) error {
	properties, params := store.buildResourceOrIdentityProperties(resource)
	cypherPath, lastResourceVar := store.buildResourceOrIdentityPath(path)
	cypher := fmt.Sprintf("MATCH p=(%s) MERGE ((%s)-[:%s]->%s) SET %s",
		cypherPath, lastResourceVar, parentRelationship, pattern, properties)
	_, err := store.write(cypher, params)
	return err
}
