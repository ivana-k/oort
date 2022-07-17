package neo4j

import (
	"github.com/c12s/oort/domain/model"
	"github.com/c12s/oort/domain/store"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type AclStore struct {
	manager *TransactionManager
}

func NewAclStore(manager *TransactionManager) store.AclStore {
	return AclStore{
		manager: manager,
	}
}

func (store AclStore) ConnectResources(parent, child model.Resource) error {
	_, err := store.manager.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MERGE (parent:Resource{id: $parentId, kind: $parentKind}) "+
				"MERGE (child:Resource{id: $childId, kind: $childKind}) "+
				"MERGE (parent)-[:Includes]->(child)",
			map[string]interface{}{"parentId": parent.Id(), "parentKind": parent.Kind(),
				"childId": child.Id(), "childKind": child.Kind()})
		if err != nil {
			return nil, err
		}

		return nil, result.Err()
	})
	return err
}

func (store AclStore) DisconnectResources(parent, child model.Resource) error {
	_, err := store.manager.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (parent:Resource{id: $parentId, kind: $parentKind}) "+
				"MATCH (child:Resource{id: $childId, kind: $childKind}) "+
				"MATCH (parent)-[conn:Includes]->(child)"+
				"DELETE conn "+
				"WITH child "+
				"CALL apoc.path.spanningTree(child, {relationshipFilter: \"Includes>\"}) YIELD path "+
				"WITH child, nodes(path) AS descendants "+
				"MATCH (descendant) WHERE descendant IN descendants "+
				"MATCH (ancestor)-[:Includes*]->(descendant) "+
				"WITH child, descendant, descendants, collect(distinct ancestor) AS ancestors "+
				"WHERE all(ancestor IN ancestors WHERE ancestor IN descendants OR ancestor = child) "+
				"DETACH DELETE descendant "+
				"WITH child "+
				"OPTIONAL MATCH (ancestor)-[:Includes*]->(child) "+
				"WITH child, collect(ancestor) AS ancestors "+
				"WHERE SIZE(ancestors) = 0 "+
				"DELETE child",
			map[string]interface{}{"parentId": parent.Id(), "parentKind": parent.Kind(),
				"childId": child.Id(), "childKind": child.Kind()})
		if err != nil {
			return nil, err
		}

		return nil, result.Err()
	})
	return err
}

func (store AclStore) UpsertAttribute(resource model.Resource, attribute model.Attribute) error {
	_, err := store.manager.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (resource:Resource{id: $id, kind: $kind}) "+
				"MERGE (attribute:Attribute{name: $attrName, kind: $attrKind})"+
				"<-[:Includes]-(resource) "+
				"SET attribute.value = $attrValue",
			map[string]interface{}{"id": resource.Id(), "kind": resource.Kind(),
				"attrName": attribute.Name(), "attrKind": attribute.Kind(),
				"attrValue": attribute.Value()})
		if err != nil {
			return nil, err
		}

		return nil, result.Err()
	})
	return err
}

func (store AclStore) RemoveAttribute(resource model.Resource, attribute model.Attribute) error {
	_, err := store.manager.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (:Resource{id: $id, kind: $kind})"+
				"-[:Includes]->"+
				"(attribute:Attribute{name: $attrName, kind: $attrKind})"+
				"DETACH DELETE attribute",
			map[string]interface{}{"id": resource.Id(), "kind": resource.Kind(),
				"attrName": attribute.Name(), "attrKind": attribute.Kind()})
		if err != nil {
			return nil, err
		}

		return nil, result.Err()
	})
	return err
}

func (store AclStore) InsertPermission(principal, resource model.Resource, permission model.Permission) error {
	return nil
}

func (store AclStore) RemovePermission(principal, resource model.Resource, permission model.Permission) error {
	return nil
}

func (store AclStore) CheckPermission(principal, resource model.Resource, permissionName string) error {
	return nil
}
