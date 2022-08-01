package neo4j

import (
	"github.com/c12s/oort/domain/model/checker"
	"github.com/c12s/oort/domain/model/syncer"
)

func connectResourcesCypher(req syncer.ConnectResourcesReq) (string, map[string]interface{}) {
	return "MERGE (parent:Resource{id: $parentId, kind: $parentKind}) " +
			"MERGE (child:Resource{id: $childId, kind: $childKind}) " +
			"MERGE (parent)-[:Includes]->(child) " +
			"MERGE (parent)-[:Includes]->(parent) " +
			"MERGE (child)-[:Includes]->(child)",
		map[string]interface{}{"parentId": req.Parent.Id(), "parentKind": req.Parent.Kind(),
			"childId": req.Child.Id(), "childKind": req.Child.Kind()}
}

func disconnectResourcesCypher(req syncer.DisconnectResourcesReq) (string, map[string]interface{}) {
	return "MATCH (parent:Resource{id: $parentId, kind: $parentKind}) " +
			"MATCH (child:Resource{id: $childId, kind: $childKind}) " +
			"MATCH (parent)-[conn:Includes]->(child)" +
			"DELETE conn " +
			"WITH child " +
			"CALL apoc.path.subgraphNodes(child, {relationshipFilter: \"Includes>\"}) YIELD node " +
			"WITH child, collect(node) AS descendants " +
			"MATCH (descendant) WHERE descendant IN descendants " +
			"MATCH (ancestor)-[:Includes*]->(descendant) " +
			"WITH child, descendant, descendants, collect(distinct ancestor) AS ancestors " +
			"WHERE all(ancestor IN ancestors WHERE ancestor IN descendants OR ancestor = child) " +
			"DETACH DELETE descendant " +
			"WITH child " +
			"OPTIONAL MATCH (ancestor)-[:Includes*]->(child) " +
			"WITH child, collect(ancestor) AS ancestors " +
			"WHERE SIZE(ancestors) = 0 " +
			"DELETE child",
		map[string]interface{}{"parentId": req.Parent.Id(), "parentKind": req.Parent.Kind(),
			"childId": req.Child.Id(), "childKind": req.Child.Kind()}
}

func upsertAttributeCypher(req syncer.UpsertAttributeReq) (string, map[string]interface{}) {
	return "MATCH (resource:Resource{id: $id, kind: $kind}) " +
			"MERGE (attribute:Attribute{name: $attrName})" +
			"<-[:Includes]-(resource) " +
			"SET attribute.value = $attrValue, attribute.kind = $attrKind",
		map[string]interface{}{"id": req.Resource.Id(), "kind": req.Resource.Kind(),
			"attrName": req.Attribute.Name(), "attrKind": req.Attribute.Kind(),
			"attrValue": req.Attribute.Value()}
}

func removeAttributeCypher(req syncer.RemoveAttributeReq) (string, map[string]interface{}) {
	return "MATCH (:Resource{id: $id, kind: $kind})" +
			"-[:Includes]->" +
			"(attribute:Attribute{name: $attrName, kind: $attrKind})" +
			"DETACH DELETE attribute",
		map[string]interface{}{"id": req.Resource.Id(), "kind": req.Resource.Kind(),
			"attrName": req.AttributeId.Name(), "attrKind": req.AttributeId.Kind()}
}

func getAttributeCypher(req checker.GetAttributeReq) (string, map[string]interface{}) {
	return "MATCH (resource:Resource{id: $id, kind: $kind}) " +
			"MATCH (attr:Attribute)<-[:Includes]-(resource) " +
			"RETURN properties(attr)",
		map[string]interface{}{"id": req.Resource.Id(), "kind": req.Resource.Kind()}
}

func insertPermissionCypher(req syncer.InsertPermissionReq) (string, map[string]interface{}) {
	return "MERGE (principal:Resource{id: $principalId, kind: $principalKind}) " +
			"MERGE (resource:Resource{id: $resourceId, kind: $resourceKind}) " +
			"MERGE (principal)-[p:Permission{name: $name}]->(resource) " +
			"SET p.condition = $condition, p.kind = $kind " +
			"MERGE (principal)-[:Includes]->(principal) " +
			"MERGE (resource)-[:Includes]->(resource)",
		map[string]interface{}{"principalId": req.Principal.Id(), "principalKind": req.Principal.Kind(),
			"resourceId": req.Resource.Id(), "resourceKind": req.Resource.Kind(), "name": req.Permission.Name(),
			"kind": req.Permission.Kind(), "condition": req.Permission.Condition().Expression()}
}

func removePermissionCypher(req syncer.RemovePermissionReq) (string, map[string]interface{}) {
	return "MATCH (principal:Resource{id: $principalId, kind: $principalKind}) " +
			"MATCH (resource:Resource{id: $resourceId, kind: $resourceKind}) " +
			"MATCH (principal)-[p:Permission{name: $name, kind: $kind}]->(resource) " +
			"DELETE p",
		map[string]interface{}{"principalId": req.Principal.Id(), "principalKind": req.Principal.Kind(),
			"resourceId": req.Resource.Id(), "resourceKind": req.Resource.Kind(), "name": req.Permission.Name(),
			"kind": req.Permission.Kind()}
}

func getPermissionAndDistanceToPrincipal(req checker.GetPermissionReq) (string, map[string]interface{}) {
	return "MATCH (principal:Resource{id: $principalId, kind: $principalKind}) " +
			"MATCH (resource:Resource{id: $resourceId, kind: $resourceKind}) " +
			"MATCH (principal)-[:Includes*]->(pParent:Resource)-[permission:Permission{name: $name}]" +
			"->(:Resource)-[:Includes*]->(resource)" +
			"MATCH path=((principal)-[:Includes*]->(pParent:Resource)) " +
			"OPTIONAL MATCH (principalAttr:Attribute)<-[:Includes]-(principal) " +
			"OPTIONAL MATCH (resourceAttr:Attribute)<-[:Includes]-(resource) " +
			"RETURN properties(permission), (size(collect(distinct (nodes(path))))-1) as distance " +
			"ORDER BY distance ASC",
		map[string]interface{}{"principalId": req.Principal.Id(), "principalKind": req.Principal.Kind(),
			"resourceId": req.Resource.Id(), "resourceKind": req.Resource.Kind(), "name": req.PermissionName}
}
