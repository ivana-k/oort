package neo4j

import (
	"github.com/c12s/oort/domain/model"
	"github.com/c12s/oort/domain/store/acl"
)

func createResourceCypher(req acl.CreateResourceReq) (string, map[string]interface{}) {
	return "MERGE (r:Resource{id: $id, kind: $kind})\n" +
			"MERGE (root:Resource{id: $rootId, kind: $rootKind})\n" +
			"MERGE (root)-[:Includes{kind: $relKind}]->(r)",
		map[string]interface{}{"id": req.Resource.Id(),
			"kind":     req.Resource.Kind(),
			"relKind":  model.CompositionRelationship,
			"rootId":   model.RootResource.Id(),
			"rootKind": model.RootResource.Kind()}
}

func deleteResourceCypher(req acl.DeleteResourceReq) (string, map[string]interface{}) {
	return "MATCH (r:Resource{id: $id, kind: $kind})\n" +
			"OPTIONAL MATCH ((r)-[:Includes*{kind: $relKind}]->(d))\n" +
			"WITH collect(d) AS descendants, r\n" +
			"OPTIONAL MATCH ((p:Permission{})-[]-(r))\n" +
			"WITH r, descendants, collect(p) AS permissions\n" +
			"MATCH (n)\n" +
			"WHERE n IN descendants OR n IN permissions OR n = r\n" +
			"WITH collect(n) AS delNodes\n" +
			"MATCH (n)\n" +
			"WHERE n IN delNodes\n" +
			"DETACH DELETE n",
		map[string]interface{}{"id": req.Resource.Id(),
			"kind":    req.Resource.Kind(),
			"relKind": model.CompositionRelationship}
}

func getResourceCypher(req acl.GetResourceReq) (string, map[string]interface{}) {
	return "MATCH (resource:Resource{id: $id, kind: $kind}) " +
			"RETURN properties(resource)",
		map[string]interface{}{"id": req.Id,
			"kind": req.Kind}
}

func createAttributeCypher(req acl.CreateAttributeReq) (string, map[string]interface{}) {
	return "MATCH (r:Resource{id: $id, kind: $kind})" +
			"WHERE NOT ((r)-[:Includes{kind: $relKind}]->(:Attribute{name: $attrName}))" +
			"CREATE (r)-[:Includes{kind: $relKind}]->" +
			"(:Attribute{name: $attrName, kind: $attrKind, value: $attrValue})",
		map[string]interface{}{"id": req.Resource.Id(),
			"kind":      req.Resource.Kind(),
			"attrName":  req.Attribute.Name(),
			"attrKind":  req.Attribute.Kind(),
			"attrValue": req.Attribute.Value(),
			"relKind":   model.CompositionRelationship}
}

func updateAttributeCypher(req acl.UpdateAttributeReq) (string, map[string]interface{}) {
	return "MATCH ((:Resource{id: $id, kind: $kind})-[:Includes{kind: $relKind}]->" +
			"(a:Attribute{name: $attrName, kind: $attrKind}))" +
			"SET a.value = $attrValue",
		map[string]interface{}{"id": req.Resource.Id(),
			"kind":      req.Resource.Kind(),
			"attrName":  req.Attribute.Name(),
			"attrKind":  req.Attribute.Kind(),
			"attrValue": req.Attribute.Value(),
			"relKind":   model.CompositionRelationship}
}

func deleteAttributeCypher(req acl.DeleteAttributeReq) (string, map[string]interface{}) {
	return "MATCH ((:Resource{id: $id, kind: $kind})-[:Includes{kind: $relKind}]->" +
			"(a:Attribute{name: $attrName}))" +
			"DETACH DELETE a",
		map[string]interface{}{"id": req.Resource.Id(),
			"kind":     req.Resource.Kind(),
			"attrName": req.AttributeId,
			"relKind":  model.CompositionRelationship}
}

func getAttributeCypher(req acl.GetAttributeReq) (string, map[string]interface{}) {
	return "MATCH (resource:Resource{id: $id, kind: $kind})\n" +
			"MATCH (attr:Attribute)<-[:Includes{kind: $relKind}]-(resource)\n" +
			"RETURN properties(attr)",
		map[string]interface{}{"id": req.Resource.Id(),
			"kind":    req.Resource.Kind(),
			"relKind": model.CompositionRelationship}
}

func createAggregationRelCypher(req acl.CreateAggregationRelReq) (string, map[string]interface{}) {
	return "MATCH (parent:Resource{id: $parentId, kind: $parentKind})\n" +
			"MATCH (child:Resource{id: $childId, kind: $childKind})\n" +
			"MERGE (parent)-[:Includes{kind: $relKind}]->(child)",
		map[string]interface{}{"parentId": req.Parent.Id(),
			"parentKind": req.Parent.Kind(),
			"childId":    req.Child.Id(),
			"childKind":  req.Child.Kind(),
			"relKind":    model.AggregateRelationship}
}

func deleteAggregationRelCypher(req acl.DeleteAggregationRelReq) (string, map[string]interface{}) {
	return "MATCH (:Resource{id: $parentId, kind: $parentKind})-[i:Includes{kind: $relKind}]->" +
			"(:Resource{id: $childId, kind: $childKind})\n" +
			"DELETE i",
		map[string]interface{}{"parentId": req.Parent.Id(),
			"parentKind": req.Parent.Kind(),
			"childId":    req.Child.Id(),
			"childKind":  req.Child.Kind(),
			"relKind":    model.AggregateRelationship}
}

func createCompositionRelCypher(req acl.CreateCompositionRelReq) (string, map[string]interface{}) {
	return "MATCH (parent:Resource{id: $parentId, kind: $parentKind})\n" +
			"MATCH (child:Resource{id: $childId, kind: $childKind})\n" +
			"MERGE (parent)-[:Includes{kind: $relKind}]->(child)",
		map[string]interface{}{"parentId": req.Parent.Id(),
			"parentKind": req.Parent.Kind(),
			"childId":    req.Child.Id(),
			"childKind":  req.Child.Kind(),
			"relKind":    model.CompositionRelationship}
}

func deleteCompositionRelCypher(req acl.DeleteCompositionRelReq) (string, map[string]interface{}) {
	return "MATCH (:Resource{id: $parentId, kind: $parentKind})-[i:Includes{kind: $relKind}]->" +
			"(:Resource{id: $childId, kind: $childKind})\n" +
			"DELETE i",
		map[string]interface{}{"parentId": req.Parent.Id(),
			"parentKind": req.Parent.Kind(),
			"childId":    req.Child.Id(),
			"childKind":  req.Child.Kind(),
			"relKind":    model.CompositionRelationship}
}

func createPermissionCypher(req acl.CreatePermissionReq) (string, map[string]interface{}) {
	return "MATCH (sub:Resource{id: $subId, kind: $subKind})\n" +
			"MATCH (obj:Resource{id: $objId, kind: $objKind})\n" +
			"WHERE NOT ((sub)-[:Has]->(:Permission{name: $name, kind: $kind})-[:On]->(obj))\n" +
			"CREATE (sub)-[:Has]->(" +
			":Permission{name: $name, kind: $kind, condition: $condition})-[:On]->(obj)",
		map[string]interface{}{"subId": req.Subject.Id(),
			"subKind":   req.Subject.Kind(),
			"objId":     req.Object.Id(),
			"objKind":   req.Object.Kind(),
			"name":      req.Permission.Name(),
			"kind":      req.Permission.Kind(),
			"condition": req.Permission.Condition().Expression()}
}

func deletePermissionCypher(req acl.DeletePermissionReq) (string, map[string]interface{}) {
	return "MATCH (sub:Resource{id: $subId, kind: $subKind})\n" +
			"MATCH (obj:Resource{id: $objId, kind: $objKind})\n" +
			"MATCH ((sub)-[:Has]->(p:Permission{name: $name, kind: $kind})-[:On]->(obj))\n" +
			"DETACH DELETE p",
		map[string]interface{}{"subId": req.Subject.Id(),
			"subKind":   req.Subject.Kind(),
			"objId":     req.Object.Id(),
			"objKind":   req.Object.Kind(),
			"name":      req.Permission.Name(),
			"kind":      req.Permission.Kind(),
			"condition": req.Permission.Condition().Expression()}
}

func getOutboxMessageCypher(message model.OutboxMessage) (string, map[string]interface{}) {
	return "CREATE (:OutboxMessage{kind: $kind, payload: $payload, processing: $processing})",
		map[string]interface{}{"kind": message.Kind,
			"payload":    message.Payload,
			"processing": false}
}

func getEffectivePermissionsWithPriorityCypher(req acl.GetPermissionHierarchyReq) (string, map[string]interface{}) {
	return "OPTIONAL MATCH (sub:Resource{id: $subId, kind: $subKind})<-[:Includes*0..]-" +
			"(subParent:Resource)-[:Has]->(p:Permission{name: $name})-[:On]->" +
			"(objParent:Resource)-[:Includes*0..]->(obj:Resource{id: $objId, kind: $objKind})\n" +
			"WHERE sub IS NOT null AND obj IS NOT null AND p IS NOT null\n" +
			"MATCH path=shortestPath((sub)<-[:Includes*0..10]-(subParent))" +
			"RETURN properties(p), -length(path) as priority " +
			"ORDER BY distance ASC",
		map[string]interface{}{"subId": req.Subject.Id(),
			"subKind":      req.Subject.Kind(),
			"resourceId":   req.Object.Id(),
			"resourceKind": req.Object.Kind(),
			"name":         req.PermissionName}
}
