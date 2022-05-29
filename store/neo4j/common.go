package neo4j

import (
	"fmt"
	"github.com/c12s/oort/domain/model"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"time"
)

func (store *permissionStore) write(cypher string, params map[string]interface{}) (neo4j.Result, error) {
	return store.transaction(neo4j.AccessModeWrite, cypher, params)
}

func (store *permissionStore) read(cypher string, params map[string]interface{}) (neo4j.Result, error) {
	return store.transaction(neo4j.AccessModeRead, cypher, params)
}

func (store *permissionStore) transaction(mode neo4j.AccessMode, cypher string, params map[string]interface{}) (neo4j.Result, error) {
	session, err := store.manager.GetSession(mode)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	result, err := session.WriteTransaction(store.handleTransaction(cypher, params))
	if err != nil {
		return nil, err
	}
	return result.(neo4j.Result), err
}

func (store *permissionStore) handleTransaction(cypher string, params map[string]interface{}) func(transaction neo4j.Transaction) (interface{}, error) {
	return func(transaction neo4j.Transaction) (interface{}, error) {
		fmt.Println(cypher)
		result, err := transaction.Run(cypher, params)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
}

func (store *permissionStore) buildResourceOrIdentityProperties(resource model.Resource) (cypher string, params map[string]interface{}) {
	params = map[string]interface{}{}
	for key, value := range resource.GetArgs() {
		param := "$" + key
		cypher += resourceVar + "." + key + " = " + param + ", "
		params[key] = value
	}
	cypher = cypher[:len(cypher)-2]
	return
}

func (store *permissionStore) buildResourceOrIdentityPath(path []model.Resource) (pattern string, lastResourceVar string) {
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

func (store *permissionStore) buildResourcePattern() string {
	return "(" + resourceVar + ":" + resourceLabel + ")"
}

func (store *permissionStore) buildIdentityPattern() string {
	return "(" + resourceVar + ":" + resourceLabel + ":" + identityLabel + ")"
}

//TODO: ovo sigurno moze drugacije, sta sa date i time tipovima
func (store *permissionStore) getValue(value interface{}) string {
	switch value.(type) {
	case time.Time:
		return fmt.Sprintf("datetime('%s')", value.(time.Time).Format("2006-01-02T15:04:05.000000000[MST]"))
	case string:
		return fmt.Sprintf("\"%s\"", value)
	default:
		return fmt.Sprint(value)
	}
}
