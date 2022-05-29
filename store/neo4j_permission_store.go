package store

import (
	"fmt"
	"github.com/c12s/oort/domain/model"
	"github.com/c12s/oort/domain/store"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"time"
)

const (
	resourceVar        = "r"
	resourceLabel      = "Resource"
	identityLabel      = "Identity"
	parentRelationship = "PARENT"
)

//TODO: error handling da bude lepsi
//TODO: previse koda ce biti, nekako podeliti

type neo4jPermissionStore struct {
	manager *Manager
}

func NewNeo4jPermissionStore(manager *Manager) store.PermissionStore {
	return &neo4jPermissionStore{
		manager: manager,
	}
}

func (store *neo4jPermissionStore) AddResource(resource model.Resource) error {
	return store.addNode(resource, store.buildResourcePattern())
}

func (store *neo4jPermissionStore) AddIdentity(resource model.Resource) error {
	return store.addNode(resource, store.buildIdentityPattern())
}

func (store *neo4jPermissionStore) AddResourceToPath(resource model.Resource, path []model.Resource) error {
	return store.addNodeToPath(resource, path, store.buildResourcePattern())
}

func (store *neo4jPermissionStore) AddIdentityToPath(resource model.Resource, path []model.Resource) error {
	return store.addNodeToPath(resource, path, store.buildIdentityPattern())
}

func (store *neo4jPermissionStore) addNode(resource model.Resource, pattern string) error {
	properties, params := store.buildResourceOrIdentityProperties(resource)
	cypher := fmt.Sprintf("CREATE %s SET %s", pattern, properties)
	_, err := store.write(cypher, params)
	return err
}

func (store *neo4jPermissionStore) addNodeToPath(resource model.Resource, path []model.Resource, pattern string) error {
	properties, params := store.buildResourceOrIdentityProperties(resource)
	cypherPath, lastResourceVar := store.buildResourceOrIdentityPath(path)
	cypher := fmt.Sprintf("MATCH p=(%s) MERGE ((%s)-[:%s]->%s) SET %s", cypherPath, lastResourceVar, parentRelationship, pattern, properties)
	_, err := store.write(cypher, params)
	return err
}

func (store *neo4jPermissionStore) write(cypher string, params map[string]interface{}) (neo4j.Result, error) {
	return store.transaction(neo4j.AccessModeWrite, cypher, params)
}

func (store *neo4jPermissionStore) read(cypher string, params map[string]interface{}) (neo4j.Result, error) {
	return store.transaction(neo4j.AccessModeRead, cypher, params)
}

func (store *neo4jPermissionStore) transaction(mode neo4j.AccessMode, cypher string, params map[string]interface{}) (neo4j.Result, error) {
	session, err := store.manager.GetSession(mode)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	result, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		fmt.Println(cypher)
		result, err := transaction.Run(cypher, params)

		if err != nil {
			return nil, err
		}

		return result, nil
	})
	if err != nil {
		return nil, err
	}
	return result.(neo4j.Result), err
}

func (store *neo4jPermissionStore) buildResourceOrIdentityProperties(resource model.Resource) (cypher string, params map[string]interface{}) {
	params = map[string]interface{}{}
	for key, value := range resource.GetArgs() {
		param := "$" + key
		cypher += resourceVar + "." + key + " = " + param + ", "
		params[key] = value
	}
	cypher = cypher[:len(cypher)-2]
	return
}

func (store *neo4jPermissionStore) buildResourceOrIdentityPath(path []model.Resource) (pattern string, lastResourceVar string) {
	for i, resource := range path {
		pattern += fmt.Sprintf("(%s%d: %s{ ", resourceVar, i, resourceLabel)
		for key, value := range resource.GetArgs() {
			pattern += fmt.Sprintf("%s: %s,", key, store.getValue(value))
		}
		pattern = pattern[:len(pattern)-1]
		pattern += "})"
		if i == len(path)-1 {
			lastResourceVar = fmt.Sprintf("%s%d", resourceVar, i)
			break
		}
		pattern += fmt.Sprintf("-[:%s]->", parentRelationship)
	}
	return
}

func (store *neo4jPermissionStore) buildResourcePattern() string {
	return "(" + resourceVar + ":" + resourceLabel + ")"
}

func (store *neo4jPermissionStore) buildIdentityPattern() string {
	return "(" + resourceVar + ":" + resourceLabel + ":" + identityLabel + ")"
}

func (store *neo4jPermissionStore) getValue(value interface{}) string {
	switch value.(type) {
	case time.Time:
		return fmt.Sprintf("datetime('%s')", value.(time.Time).Format("2006-01-02T15:04:05.000000000[MST]"))
	case string:
		return fmt.Sprintf("\"%s\"", value)
	default:
		return fmt.Sprint(value)
	}
}
