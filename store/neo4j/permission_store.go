package neo4j

import (
	"fmt"
	"github.com/c12s/oort/domain/model"
	"github.com/c12s/oort/domain/store"
	storemodel "github.com/c12s/oort/store/neo4j/model"
)

//TODO: error handling

type permissionStore struct {
	handler *TransactionHandler
}

func NewNeo4jPermissionStore(handler *TransactionHandler) store.PermissionStore {
	return &permissionStore{
		handler: handler,
	}
}

func (store *permissionStore) AddResource(resource model.Resource) error {
	r := storemodel.Resource{Resource: resource}
	return store.addNode(r, r.ResourcePattern("r"))
}

func (store *permissionStore) AddIdentity(resource model.Resource) error {
	r := storemodel.Resource{Resource: resource}
	return store.addNode(r, r.IdentityPattern("r"))
}

func (store *permissionStore) AddResourceToPath(resource model.Resource, path storemodel.Path) error {
	r := storemodel.Resource{Resource: resource}
	return store.addNodeToPath(r, path, r.ResourcePattern("r"))
}

func (store *permissionStore) AddIdentityToPath(resource model.Resource, path storemodel.Path) error {
	r := storemodel.Resource{Resource: resource}
	return store.addNodeToPath(r, path, r.IdentityPattern("r"))
}

func (store *permissionStore) Connect(parentPath storemodel.Path, childPath storemodel.Path) error {
	parentCypherPath, parentVar := parentPath.Path("p")
	childCypherPath, childVar := childPath.Path("c")
	cypher := fmt.Sprintf("MATCH %s MATCH %s MERGE ((%s)-[:%s]->(%s))",
		parentCypherPath, childCypherPath, parentVar, parentPath.ParentRelationship(), childVar)
	_, err := store.handler.Write(cypher, nil)
	return err
}

func (store *permissionStore) AddPermission(identityPath storemodel.Path, resourcePath storemodel.Path, permission model.Permission) error {
	identityCypherPath, identityVar := identityPath.IdentityPath("i")
	resourceCypherPath, resourceVar := resourcePath.Path("r")
	cypher := fmt.Sprintf("MATCH %s MATCH %s MERGE ((%s)-[:%s{name:$name}]->(%s))",
		identityCypherPath, resourceCypherPath, identityVar, identityPath.PermissionRelationship(), resourceVar)
	_, err := store.handler.Write(cypher, map[string]interface{}{"name": permission.GetName()})
	return err
}

func (store *permissionStore) addNode(resource storemodel.Resource, pattern string) error {
	properties, params := resource.Properties("r")
	cypher := fmt.Sprintf("CREATE %s SET %s", pattern, properties)
	_, err := store.handler.Write(cypher, params)
	return err
}

func (store *permissionStore) addNodeToPath(resource storemodel.Resource, path storemodel.Path, pattern string) error {
	properties, params := resource.Properties("r")
	cypherPath, lastResourceVar := path.Path("r")
	cypher := fmt.Sprintf("MATCH p=(%s) MERGE ((%s)-[:%s]->%s) SET %s",
		cypherPath, lastResourceVar, path.ParentRelationship(), pattern, properties)
	_, err := store.handler.Write(cypher, params)
	return err
}
