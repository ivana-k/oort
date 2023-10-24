package neo4j

import (
	"github.com/c12s/oort/internal/domain"
)

type CypherFactory interface {
	createResource(req domain.CreateResourceReq) (string, map[string]interface{})
	deleteResource(req domain.DeleteResourceReq) (string, map[string]interface{})
	getResource(req domain.GetResourceReq) (string, map[string]interface{})
	putAttribute(req domain.PutAttributeReq) (string, map[string]interface{})
	deleteAttribute(req domain.DeleteAttributeReq) (string, map[string]interface{})
	createInheritanceRel(req domain.CreateInheritanceRelReq) (string, map[string]interface{})
	deleteInheritanceRel(req domain.DeleteInheritanceRelReq) (string, map[string]interface{})
	createPolicy(req domain.CreatePolicyReq) (string, map[string]interface{})
	deletePolicy(req domain.DeletePolicyReq) (string, map[string]interface{})
	getEffectivePermissionsWithPriority(req domain.GetPermissionHierarchyReq) (string, map[string]interface{})
}

type simpleCypherFactory struct {
}

func NewSimpleCypherFactory() CypherFactory {
	return &simpleCypherFactory{}
}

const ncCreateResourceCypher = `
MERGE (r:Resource{name: $name})
MERGE (root:Resource{name: $rootName})
MERGE (r)-[:INHERITS_FROM]->(root)
`

func (f simpleCypherFactory) createResource(req domain.CreateResourceReq) (string, map[string]interface{}) {
	return ncCreateResourceCypher,
		map[string]interface{}{
			"name":     req.Resource.Name(),
			"rootName": domain.RootResource.Name()}
}

const ncDeleteResourceCypher = `
MATCH (r:Resource{name: $name})
WITH r
// delete all attributes of r
CALL {
    WITH r
    MATCH (r)-[:HAS]->(a:Attribute)
    DETACH DELETE a
}
// delete all directly assigned permissions of r
CALL {
    WITH r
    MATCH (r)-[:HAS|ON]-(p:Permission)
    DETACH DELETE p
}
// delete r
CALL {
    WITH r
    DETACH DELETE r
}
`

func (f simpleCypherFactory) deleteResource(req domain.DeleteResourceReq) (string, map[string]interface{}) {
	return ncDeleteResourceCypher,
		map[string]interface{}{
			"name": req.Resource.Name()}
}

const ncGetResourceCypher = `
MATCH (resource:Resource{name: $name})
OPTIONAL MATCH (attr:Attribute)<-[:HAS]-(resource)
RETURN resource.name, collect(properties(attr)) as attrs
`

func (f simpleCypherFactory) getResource(req domain.GetResourceReq) (string, map[string]interface{}) {
	return ncGetResourceCypher,
		map[string]interface{}{
			"name": req.Resource.Name()}
}

const ncPutAttributeCypher = `
MERGE (r:Resource{name: $name})
MERGE (root:Resource{name: $rootName})
MERGE (r)-[:INHERITS_FROM]->(root)
MERGE ((r)-[:HAS]->(a:Attribute{name: $attrName}))
SET a += {kind: $attrKind, value: $ attrValue}
`

func (f simpleCypherFactory) putAttribute(req domain.PutAttributeReq) (string, map[string]interface{}) {
	return ncPutAttributeCypher,
		map[string]interface{}{
			"name":      req.Resource.Name(),
			"rootName":  domain.RootResource.Name(),
			"attrName":  req.Attribute.Name(),
			"attrKind":  req.Attribute.Kind(),
			"attrValue": req.Attribute.Value()}
}

const ncDeleteAttributeCypher = `
MATCH ((:Resource{name: $name})-[:HAS]->(a:Attribute{name: $attrName}))
DETACH DELETE a
`

func (f simpleCypherFactory) deleteAttribute(req domain.DeleteAttributeReq) (string, map[string]interface{}) {
	return ncDeleteAttributeCypher,
		map[string]interface{}{
			"name":     req.Resource.Name(),
			"attrName": req.AttributeId.Name()}
}

const ncCreateInheritanceRelCypher = `
MERGE (from:Resource{name: $fromName})
MERGE (to:Resource{name: $toName})
MERGE (root:Resource{name: $rootName})
MERGE (from)-[:INHERITS_FROM]->(root)
MERGE (to)-[:INHERITS_FROM]->(root)
WITH from, to
MATCH (f) WHERE ID(f) = ID(from)
MATCH (t) WHERE ID(t) = ID(to)
AND NOT (t)-[:INHERITS_FROM]->(f) AND NOT (f)-[:INHERITS_FROM*]->(t)
CREATE (t)-[:INHERITS_FROM]->(f)
`

func (f simpleCypherFactory) createInheritanceRel(req domain.CreateInheritanceRelReq) (string, map[string]interface{}) {
	return ncCreateInheritanceRelCypher,
		map[string]interface{}{
			"fromName": req.From.Name(),
			"toName":   req.To.Name(),
			"rootName": domain.RootResource.Name()}
}

const ncDeleteInheritanceRelCypher = `
MATCH (:Resource{name: $toName})-[rel:INHERITS_FROM]->(:Resource{name: $fromName})
DELETE rel
`

func (f simpleCypherFactory) deleteInheritanceRel(req domain.DeleteInheritanceRelReq) (string, map[string]interface{}) {
	return ncDeleteInheritanceRelCypher,
		map[string]interface{}{
			"fromName": req.From.Name(),
			"toName":   req.To.Name()}
}

const ncCreatePermissionCypher = `
MERGE (sub:Resource{name: $subName})
MERGE (obj:Resource{name: $objName})
MERGE (root:Resource{name: $rootName})
MERGE (sub)-[:INHERITS_FROM]->(root)
MERGE (obj)-[:INHERITS_FROM]->(root)
MERGE ((sub)-[:HAS]->(p:Permission{name: $permName, kind: $permKind})-[:ON]->(obj))
SET p.condition = $permCond
`

func (f simpleCypherFactory) createPolicy(req domain.CreatePolicyReq) (string, map[string]interface{}) {
	return ncCreatePermissionCypher,
		map[string]interface{}{
			"subName":  req.SubjectScope.Name(),
			"objName":  req.ObjectScope.Name(),
			"rootName": domain.RootResource.Name(),
			"permName": req.Permission.Name(),
			"permKind": req.Permission.Kind(),
			"permCond": req.Permission.Condition().Expression()}
}

const ncDeletePermissionCypher = `
MATCH (sub:Resource{name: $subName})
MATCH (obj:Resource{name: $objName})
MATCH ((sub)-[:HAS]->(p:Permission{name: $permName, kind: $permKind})-[:ON]->(obj))
DETACH DELETE p
`

func (f simpleCypherFactory) deletePolicy(req domain.DeletePolicyReq) (string, map[string]interface{}) {
	return ncDeletePermissionCypher,
		map[string]interface{}{
			"subName":  req.SubjectScope.Name(),
			"objName":  req.ObjectScope.Name(),
			"permName": req.Permission.Name(),
			"permKind": req.Permission.Kind()}
}

const ncGetPermissionsCypher = `
MATCH (sub:Resource{name: $subName})-[:INHERITS_FROM*0..]->(subParent:Resource)-[:HAS]->
(p:Permission{name: $permName})-[:ON]->(objParent:Resource)<-[:INHERITS_FROM*0..]-(obj:Resource{name: $objName})
WITH p, sub, subParent, obj, objParent
CALL {
	WITH sub, subParent
	MATCH path=(sub)-[:INHERITS_FROM*0..100]->(subParent)
	RETURN -length(path) AS subPriority
	ORDER BY subPriority ASC
	LIMIT 1
}
CALL {
	WITH obj, objParent
	MATCH path=(obj)-[:INHERITS_FROM*0..100]->(objParent)
	RETURN -length(path) AS objPriority
	ORDER BY objPriority ASC
	LIMIT 1
}
RETURN p.name, p.kind, p.condition, subPriority, objPriority
`

func (f simpleCypherFactory) getEffectivePermissionsWithPriority(req domain.GetPermissionHierarchyReq) (string, map[string]interface{}) {
	return ncGetPermissionsCypher,
		map[string]interface{}{
			"subName":  req.Subject.Name(),
			"objName":  req.Object.Name(),
			"permName": req.PermissionName}
}

// todo: sredi ovo
//type cachedPermsCypherFactory struct {
//}
//
//func NewCachedPermsCypherFactory() cypherFactory {
//	return &cachedPermsCypherFactory{}
//}
//
//const cCreateResourceCypher = `
//MERGE (r:Resource{name: $name})
//MERGE (root:Resource{name: $rootName})
//MERGE (r)-[:INHERITS_FROM]->(root)
//WITH r, root
//// repos perms that r inherits from root
//CALL {
//	WITH r, root
//	MATCH (root)-[srel:HAS]->(p:Permission)
//	MERGE (r)-[:HAS{priority: srel.priority - 1}]->(p)
//}
//CALL {
//	WITH r, root
//	MATCH (root)<-[orel:ON]-(p:Permission)
//	MERGE (r)<-[:On{priority: orel.priority - 1}]-(p)
//}
//`
//
//func (f cachedPermsCypherFactory) createResourceCypher(req rhabac.CreateResourceReq) (string, map[string]interface{}) {
//	return cCreateResourceCypher,
//		map[string]interface{}{
//			"name":     req.Resource.Name(),
//			"rootName": model.RootResource.Name()}
//}
//
//const cDeleteResourceCypher = `
//MATCH (r:Resource{name: $name})
//WITH r
//// obrisi sve atribute resursa za brisanje
//CALL {
//   WITH r
//   MATCH (r)-[:HAS]->(a:Attribute)
//   DETACH DELETE a
//}
//// obrisi sve direktno dodeljene dozvole resursa za brisanje
//CALL {
//   WITH r
//   MATCH (r)-[:HAS|ON{priority: 0}]-(p:Permission{})
//   DETACH DELETE p
//}
//// nadji sve resurse kojima treba ukloniti dozvole
//CALL {
//   WITH r
//   MATCH (d:Resource)-[:INHERITS_FROM*]->(r)
//   RETURN collect(d) AS descendants
//}
//// ukloni nasledjene dozvole iz potomaka
//CALL {
//   WITH r, descendants
//   UNWIND descendants AS d
//   // nadji i obrisi sve putanje resursa do direktno dodeljene dozvole - subjekat
//   CALL {
//     WITH r, d
//     MATCH path=(d)-[:INHERITS_FROM*]->(:Resource)-[:HAS{priority: 0}]->(perm:Permission)
//     WHERE ANY(resOnPath IN NODES(path) WHERE resOnPath = r)
//     MATCH ((d)-[srel:HAS{priority: -(length(path)-1)}]->(perm))
//     WITH collect(srel) AS dels
//     UNWIND dels[..1] AS del
//     DELETE del
//   }
//   // nadji i obrisi sve putanje resursa do direktno dodeljene dozvole - objekat
//   CALL {
//     WITH r, d
//     MATCH path=(d)-[:INHERITS_FROM*]->(:Resource)<-[:ON{priority: 0}]-(perm:Permission)
//     WHERE ANY(resOnPath IN NODES(path) WHERE resOnPath = r)
//     MATCH ((d)<-[orel:ON{priority: -(length(path)-1)}]-(perm))
//     WITH collect(orel) AS dels
//     UNWIND dels[..1] AS del
//     DELETE del
//   }
//}
//// obrisi resurs
//CALL {
//   WITH r
//   DETACH DELETE r
//}
//`
//
//func (f cachedPermsCypherFactory) deleteResourceCypher(req rhabac.DeleteResourceReq) (string, map[string]interface{}) {
//	return cDeleteResourceCypher,
//		map[string]interface{}{
//			"name": req.Resource.Name()}
//}
//
//const cGetResourceCypher = ncGetResourceCypher
//
//func (f cachedPermsCypherFactory) getResource(req rhabac.GetResourceReq) (string, map[string]interface{}) {
//	return cGetResourceCypher,
//		map[string]interface{}{
//			"name": req.Resource.Name()}
//}
//
//const cPutAttributeCypher = ncPutAttributeCypher
//
//func (f cachedPermsCypherFactory) putAttribute(req rhabac.PutAttributeReq) (string, map[string]interface{}) {
//	return cPutAttributeCypher,
//		map[string]interface{}{
//			"name":      req.Resource.Name(),
//			"rootName":  model.RootResource.Name(),
//			"attrName":  req.Attribute.Name(),
//			"attrKind":  req.Attribute.Kind(),
//			"attrValue": req.Attribute.Value()}
//}
//
//const cDeleteAttributeCypher = ncDeleteAttributeCypher
//
//func (f cachedPermsCypherFactory) deleteAttribute(req rhabac.DeleteAttributeReq) (string, map[string]interface{}) {
//	return cDeleteAttributeCypher,
//		map[string]interface{}{
//			"name":     req.Resource.Name(),
//			"attrName": req.AttributeId}
//}
//
//const cCreateInheritanceRelCypher = `
//MERGE (from:Resource{name: $fromName})
//MERGE (to:Resource{name: $toName})
//MERGE (root:Resource{name: $rootName})
//MERGE (from)-[:INHERITS_FROM]->(root)
//MERGE (to)-[:INHERITS_FROM]->(root)
//WHERE NOT (from)-[:INHERITS_FROM]->(to) AND NOT (to)-[:INHERITS_FROM*]->(from)
//// create relationship
//CREATE (from)-[newRel:INHERITS_FROM]->(to)
//// find nwwly inherited permissions
//WITH from, newRel
//CALL {
//   WITH from
//   MATCH (from)-[srel:HAS|ON]-(p:Permission)
//   RETURN collect({priority: srel.priority, type: type(srel), permission: p}) AS rels
//}
//// find new paths from 'from' resource
//CALL {
//   WITH from, newRel
//   MATCH path=(from)-[:INHERITS_FROM*]->(d:Resource)
//   WHERE newRel in RELATIONSHIPS(path)
//   RETURN collect(path) AS newPaths
//}
////  assign permissions to descendants
//CALL {
//   WITH newPaths, rels
//   UNWIND newPaths AS newPath
//   WITH last(nodes(newPath)) AS res, length(newPath) AS dist, rels AS policies
//   UNWIND policies AS policy
//   WITH policy.priority AS priority, policy.type AS type, policy.permission AS policy, res, dist
//   FOREACH (i in CASE WHEN type = "HAS" THEN [1] ELSE [] END |
//       CREATE (res)-[:Has{priority: priority - dist}]->(policy)
//   )
//   FOREACH (i in CASE WHEN type = "ON" THEN [1] ELSE [] END |
//       CREATE (res)<-[:On{priority: priority - dist}]-(policy)
//   )
//}
//`
//
//func (f cachedPermsCypherFactory) createInheritanceRelCypher(req rhabac.CreateInheritanceRelReq) (string, map[string]interface{}) {
//	return cCreateInheritanceRelCypher,
//		map[string]interface{}{
//			"fromName": req.From.Name(),
//			"toName":   req.To.Name(),
//			"rootName": model.RootResource.Name()}
//}
//
//const cDeleteINheritanceRelCypher = `
//MATCH (from:Resource{name: $fromName})-[includes:INHERITS_FROM]->(:Resource{name: $toName})
//WITH from, includes
//// nadji sve dozvole koje ima roditelj (kao subjekat)
//CALL {
//   WITH parent
//   MATCH (parent)-[rel:Has]->(p:Permission)
//   RETURN collect({permission: p, priority: rel.priority}) AS subPermissions
//}
//// nadji sve dozvole koje ima roditelj (kao objekat)
//CALL {
//   WITH parent
//   MATCH (parent)<-[rel:On]-(p:Permission)
//   RETURN collect({permission: p, priority: rel.priority}) AS objPermissions
//}
//// pronadji sve putanje od roditelja koje ce biti prekinute
//CALL {
//   WITH parent, includes
//   MATCH path=(parent)-[:Includes*]->(d:Resource)
//   WHERE includes IN RELATIONSHIPS(path)
//   RETURN collect(path) AS oldPaths
//}
//// obrisi sve nasledjene dozvole
//CALL {
//   WITH parent, subPermissions, objPermissions, oldPaths
//   CALL {
//       WITH parent, subPermissions, oldPaths
//       UNWIND oldPaths AS oldPath
//       WITH last(nodes(oldPath)) AS res, length(oldPath) AS dist, subPermissions
//       UNWIND subPermissions AS policy
//       WITH policy.priority AS priority, policy.permission AS permission, res, dist
//       MATCH (res)-[rrel:Has{priority: priority - dist}]->(permission)
//       WITH permission, rrel.priority as priority, collect(rrel) AS del
//       UNWIND del[..1] AS d
//       DELETE d
//   }
//   CALL {
//       WITH parent, objPermissions, oldPaths
//       UNWIND oldPaths AS oldPath
//       WITH last(nodes(oldPath)) AS res, length(oldPath) AS dist, objPermissions
//       UNWIND objPermissions AS policy
//       WITH policy.priority AS priority, policy.permission AS permission, res, dist
//       MATCH (res)<-[rrel:On{priority: priority - dist}]-(permission)
//       WITH permission, rrel.priority as priority, collect(rrel) AS del
//       UNWIND del[..1] AS d
//       DELETE d
//   }
//}
//// obrisi vezu izmedju roditelja i deteta
//CALL {
//   WITH includes
//   DELETE includes
//}
//`
//
//func (f cachedPermissionsCypherFactory) deleteAggregationRelCypher(req rhabac.DeleteAggregationRelReq) (string, map[string]interface{}) {
//	return cDeleteRelCypher,
//		map[string]interface{}{
//			"parentName":  req.Parent.Name(),
//			"childName":   req.Child.Name(),
//			"relKind":     model.AggregateRelationship,
//			"composition": model.CompositionRelationship,
//			"rootName":    model.RootResource.Name()}
//}

//const cCreatePermissionCypher = `
//MATCH (sub:Resource{name: $subName})
//MATCH (obj:Resource{name: $objName})
//WHERE NOT (sub)-[:Has{priority: 0}]->(:Permission{name: $permName, kind: $permKind})-[:On{priority: 0}]->(obj)
//CREATE (sub)-[srel:Has{priority: 0}]->(p:Permission{name: $permName, kind: $permKind, condition: $permCond})-[orel:On{priority: 0}]->(obj)
//WITH sub, obj, p
//CALL {
//    WITH sub, p
//    MATCH path=((sub)-[:Includes*]->(descendant:Resource))
//    CREATE (descendant)-[:Has{priority: -length(path)}]->(p)
//}
//CALL {
//    WITH obj, p
//    MATCH path=((obj)-[:Includes*]->(descendant:Resource))
//    CREATE (descendant)<-[:On{priority: -length(path)}]-(p)
//}
//`
//
//func (f cachedPermissionsCypherFactory) createPermissionCypher(req rhabac.CreatePermissionReq) (string, map[string]interface{}) {
//	return cCreatePermissionCypher,
//		map[string]interface{}{
//			"subName":  req.Subject.Name(),
//			"objName":  req.Object.Name(),
//			"permName": req.Permission.Name(),
//			"permKind": req.Permission.Kind(),
//			"permCond": req.Permission.Condition().Expression()}
//}
//
//const cDeletePermissionCypher = `
//MATCH (sub:Resource{name: $subName})
//MATCH (obj:Resource{name: $objName})
//MATCH (sub)-[:Has{priority: 0}]->(p:Permission{name: $permName, kind: $permKind})-[:On{priority: 0}]->(obj)
//DETACH DELETE p
//`
//
//func (f cachedPermissionsCypherFactory) deletePermissionCypher(req rhabac.DeletePermissionReq) (string, map[string]interface{}) {
//	return cDeletePermissionCypher,
//		map[string]interface{}{
//			"subName":  req.Subject.Name(),
//			"objName":  req.Object.Name(),
//			"permName": req.Permission.Name(),
//			"permKind": req.Permission.Kind()}
//}
//
//const cGetPermissionsCypher = `
//MATCH (sub:Resource{name: $subName})
//MATCH (obj:Resource{name: $objName})
//MATCH (sub)-[srel:HAS]->(p:Permission{name: $permName})-[orel:ON]->(obj)
//WITH p, min(srel.priority) AS spriority, min(orel.priority) AS opriority
//RETURN p.name, p.kind, p.condition, spriority, opriority
//`
//
//func (f cachedPermsCypherFactory) getEffectivePermissionsWithPriorityCypher(req rhabac.GetPermissionHierarchyReq) (string, map[string]interface{}) {
//	return cGetPermissionsCypher,
//		map[string]interface{}{
//			"subName":  req.Subject.Name(),
//			"objName":  req.Object.Name(),
//			"permName": req.PermissionName}
//}
