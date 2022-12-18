package neo4j

import (
	"github.com/c12s/oort/domain/model"
	"github.com/c12s/oort/domain/store/acl"
)

type cypherFactory interface {
	createResourceCypher(req acl.CreateResourceReq) (string, map[string]interface{})
	deleteResourceCypher(req acl.DeleteResourceReq) (string, map[string]interface{})
	getResourceCypher(req acl.GetResourceReq) (string, map[string]interface{})
	createAttributeCypher(req acl.CreateAttributeReq) (string, map[string]interface{})
	updateAttributeCypher(req acl.UpdateAttributeReq) (string, map[string]interface{})
	deleteAttributeCypher(req acl.DeleteAttributeReq) (string, map[string]interface{})
	createAggregationRelCypher(req acl.CreateAggregationRelReq) (string, map[string]interface{})
	deleteAggregationRelCypher(req acl.DeleteAggregationRelReq) (string, map[string]interface{})
	createCompositionRelCypher(req acl.CreateCompositionRelReq) (string, map[string]interface{})
	deleteCompositionRelCypher(req acl.DeleteCompositionRelReq) (string, map[string]interface{})
	createPermissionCypher(req acl.CreatePermissionReq) (string, map[string]interface{})
	deletePermissionCypher(req acl.DeletePermissionReq) (string, map[string]interface{})
	getEffectivePermissionsWithPriorityCypher(req acl.GetPermissionHierarchyReq) (string, map[string]interface{})
}

type nonCachedPermissionsCypherFactory struct {
}

func NewNonCachedPermissionsCypherFactory() cypherFactory {
	return &nonCachedPermissionsCypherFactory{}
}

const ncCreateResourceQuery = `
MERGE (r:Resource{name: $name})
MERGE (root:Resource{name: $rootName})
MERGE (root)-[:Includes{kind: $composition}]->(r)
`

func (f nonCachedPermissionsCypherFactory) createResourceCypher(req acl.CreateResourceReq) (string, map[string]interface{}) {
	return ncCreateResourceQuery,
		map[string]interface{}{
			"name":        req.Resource.Name(),
			"rootName":    model.RootResource.Name(),
			"composition": model.CompositionRelationship}
}

const ncDeleteResourceQuery = `
MATCH (r:Resource{name: $name})
WITH r
// nadji sve resurse za brisanje
CALL {
    WITH r
    OPTIONAL MATCH (r)-[:Includes*{kind: $composition}]->(d:Resource)
    RETURN collect(d) + collect(r) AS delRes
}
// obrisi sve atribute resursa za brisanje
CALL {
    WITH delRes
    UNWIND delRes AS r
    MATCH (r)-[:Includes{kind: $composition}]->(a:Attribute)
    DETACH DELETE a
}
// obrisi sve direktno dodeljene dozvole resursa za brisanje
CALL {
    WITH delRes
    UNWIND delRes AS r
    MATCH (r)-[:Has|On]-(p:Permission)
    DETACH DELETE p
}
// nadji sve resurse koje treba povezati sa root-om (prvi nivo)
CALL {
    WITH delRes
    MATCH (p:Resource)-[:Includes]->(c:Resource)
    WHERE p IN delRes AND NOT c IN delRes AND NOT EXISTS
        {((c)<-[:Includes]-(op:Resource)) WHERE NOT (op IN delRes)}
    RETURN collect(DISTINCT c) AS rootRes
}
// obrisi resurse
CALL {
    WITH delRes
    UNWIND delRes AS r
    DETACH DELETE r
}
// povezati resurse sa root-om
CALL {
    WITH rootRes
    MATCH (root:Resource{name: $rootName})
    UNWIND rootRes AS rr
    CREATE (root)-[:Includes{kind: $composition}]->(rr)
}
`

func (f nonCachedPermissionsCypherFactory) deleteResourceCypher(req acl.DeleteResourceReq) (string, map[string]interface{}) {
	return ncDeleteResourceQuery,
		map[string]interface{}{
			"name":        req.Resource.Name(),
			"composition": model.CompositionRelationship,
			"rootName":    model.RootResource.Name()}
}

const ncGetResourceQuery = `
MATCH (resource:Resource{name: $name})
OPTIONAL MATCH (attr:Attribute)<-[:Includes{kind: $composition}]-(resource)
RETURN resource.name, collect(properties(attr)) as attrs
`

func (f nonCachedPermissionsCypherFactory) getResourceCypher(req acl.GetResourceReq) (string, map[string]interface{}) {
	return ncGetResourceQuery,
		map[string]interface{}{
			"name":        req.Resource.Name(),
			"composition": model.CompositionRelationship}
}

const ncCreateAttributeQuery = `
MATCH (r:Resource{name: $name})
WHERE NOT ((r)-[:Includes{kind: $composition}]->(:Attribute{name: $attrName}))
CREATE (r)-[:Includes{kind: $composition}]->(:Attribute{name: $attrName, kind: $attrKind, value: $attrValue})
`

func (f nonCachedPermissionsCypherFactory) createAttributeCypher(req acl.CreateAttributeReq) (string, map[string]interface{}) {
	return ncCreateAttributeQuery,
		map[string]interface{}{
			"name":        req.Resource.Name(),
			"attrName":    req.Attribute.Name(),
			"attrKind":    req.Attribute.Kind(),
			"attrValue":   req.Attribute.Value(),
			"composition": model.CompositionRelationship}
}

const ncUpdateAttributeQuery = `
MATCH ((:Resource{name: $name})-[:Includes{kind: $composition}]->(a:Attribute{name: $attrName, kind: $attrKind}))
SET a.value = $attrValue
`

func (f nonCachedPermissionsCypherFactory) updateAttributeCypher(req acl.UpdateAttributeReq) (string, map[string]interface{}) {
	return ncUpdateAttributeQuery,
		map[string]interface{}{
			"name":        req.Resource.Name(),
			"attrName":    req.Attribute.Name(),
			"attrKind":    req.Attribute.Kind(),
			"attrValue":   req.Attribute.Value(),
			"composition": model.CompositionRelationship}
}

const ncDeleteAttributeQuery = `
MATCH ((:Resource{name: $name})-[:Includes{kind: $composition}]->(a:Attribute{name: $attrName}))
DETACH DELETE a
`

func (f nonCachedPermissionsCypherFactory) deleteAttributeCypher(req acl.DeleteAttributeReq) (string, map[string]interface{}) {
	return ncDeleteAttributeQuery,
		map[string]interface{}{
			"name":        req.Resource.Name(),
			"attrName":    req.AttributeId,
			"composition": model.CompositionRelationship}
}

const ncCreateRelQuery = `
MATCH (parent:Resource{name: $parentName})
MATCH (child:Resource{name: $childName})
WHERE NOT (parent)-[:Includes]->(child) AND NOT (child)-[:Includes*]->(parent)
CREATE (parent)-[:Includes{kind: $relKind}]->(child)
// RASKIDANJE VEZE SA ROOT-OM
CALL {
    WITH child
    MATCH (child)<-[rootRel:Includes{kind: $composition}]-(root:Resource{name: $rootName})
    DELETE rootRel
}
`

func (f nonCachedPermissionsCypherFactory) createAggregationRelCypher(req acl.CreateAggregationRelReq) (string, map[string]interface{}) {
	return ncCreateRelQuery,
		map[string]interface{}{
			"parentName":  req.Parent.Name(),
			"childName":   req.Child.Name(),
			"relKind":     model.AggregateRelationship,
			"composition": model.CompositionRelationship,
			"rootName":    model.RootResource.Name()}
}

const ncDeleteRelQuery = `
MATCH (parent:Resource{name: $parentName})-[includes:Includes{kind: $relKind}]->(child:Resource{name: $childName})
WITH parent, child, includes
// obrisi vezu izmedju roditelja i deteta
CALL {
    WITH includes
    DELETE includes
}
// povezi child sa root-om ako nema drugih roditelja
CALL {
    WITH child
    MATCH (c)
    WHERE c = child AND  NOT (child)<-[:Includes]-()
    MATCH (root:Resource{name: $rootName})
    CREATE (root)-[:Includes{kind: $composition}]->(child)
}
`

func (f nonCachedPermissionsCypherFactory) deleteAggregationRelCypher(req acl.DeleteAggregationRelReq) (string, map[string]interface{}) {
	return ncDeleteRelQuery,
		map[string]interface{}{
			"parentName":  req.Parent.Name(),
			"childName":   req.Child.Name(),
			"relKind":     model.AggregateRelationship,
			"composition": model.CompositionRelationship,
			"rootName":    model.RootResource.Name()}
}

func (f nonCachedPermissionsCypherFactory) createCompositionRelCypher(req acl.CreateCompositionRelReq) (string, map[string]interface{}) {
	return ncCreateRelQuery,
		map[string]interface{}{
			"parentName":  req.Parent.Name(),
			"childName":   req.Child.Name(),
			"relKind":     model.CompositionRelationship,
			"composition": model.CompositionRelationship,
			"rootName":    model.RootResource.Name()}
}

func (f nonCachedPermissionsCypherFactory) deleteCompositionRelCypher(req acl.DeleteCompositionRelReq) (string, map[string]interface{}) {
	return ncDeleteRelQuery,
		map[string]interface{}{
			"parentName":  req.Parent.Name(),
			"childName":   req.Child.Name(),
			"relKind":     model.CompositionRelationship,
			"composition": model.CompositionRelationship,
			"rootName":    model.RootResource.Name()}
}

const ncCreatePermissionQuery = `
MATCH (sub:Resource{name: $subName})
MATCH (obj:Resource{name: $objName})
WHERE NOT ((sub)-[:Has]->(:Permission{name: $permName, kind: $permKind})-[:On]->(obj))
CREATE (sub)-[:Has]->(:Permission{name: $permName, kind: $permKind, condition: $permCond})-[:On]->(obj)
`

func (f nonCachedPermissionsCypherFactory) createPermissionCypher(req acl.CreatePermissionReq) (string, map[string]interface{}) {
	return ncCreatePermissionQuery,
		map[string]interface{}{
			"subName":  req.Subject.Name(),
			"objName":  req.Object.Name(),
			"permName": req.Permission.Name(),
			"permKind": req.Permission.Kind(),
			"permCond": req.Permission.Condition().Expression()}
}

const ncDeletePermissionQuery = `
MATCH (sub:Resource{name: $subName})
MATCH (obj:Resource{name: $objName})
MATCH ((sub)-[:Has]->(p:Permission{name: $permName, kind: $permKind})-[:On]->(obj))
DETACH DELETE p
`

func (f nonCachedPermissionsCypherFactory) deletePermissionCypher(req acl.DeletePermissionReq) (string, map[string]interface{}) {
	return ncDeletePermissionQuery,
		map[string]interface{}{
			"subName":  req.Subject.Name(),
			"objName":  req.Object.Name(),
			"permName": req.Permission.Name(),
			"permKind": req.Permission.Kind()}
}

const ncGetPermissionsQuery = `
OPTIONAL MATCH (sub:Resource{name: $subName})<-[:Includes*0..]-(subParent:Resource)-[:Has]->
(p:Permission{name: $permName})-[:On]->(objParent:Resource)-[:Includes*0..]->(obj:Resource{name: $objName})
WHERE sub IS NOT null AND obj IS NOT null AND p IS NOT null
MATCH subPath=shortestPath((sub)<-[:Includes*0..10]-(subParent))
MATCH objPath=shortestPath((obj)<-[:Includes*0..10]-(objParent))
RETURN p.name, p.kind, p.condition, -length(subPath), -length(objPath)
`

func (f nonCachedPermissionsCypherFactory) getEffectivePermissionsWithPriorityCypher(req acl.GetPermissionHierarchyReq) (string, map[string]interface{}) {
	return ncGetPermissionsQuery,
		map[string]interface{}{
			"subName":  req.Subject.Name(),
			"objName":  req.Object.Name(),
			"permName": req.PermissionName}
}

type cachedPermissionsCypherFactory struct {
}

func NewCachedPermissionsCypherFactory() cypherFactory {
	return &cachedPermissionsCypherFactory{}
}

const cCreateResourceQuery = `
MERGE (r:Resource{name: $name})
MERGE (root:Resource{name: $rootName})
MERGE (root)-[:Includes{kind: $composition}]->(r)
WITH r, root
CALL {
	WITH r, root
	MATCH (root)-[srel:Has]->(p:Permission)
	MERGE (r)-[:Has{priority: srel.prioriy - 1}]->(p)
}
CALL {
	WITH r, root
	MATCH (root)<-[orel:On]-(p:Permission)
	MERGE (r)<-[:On{priority: orel.priority - 1}]-(p)
}
`

func (f cachedPermissionsCypherFactory) createResourceCypher(req acl.CreateResourceReq) (string, map[string]interface{}) {
	return cCreateResourceQuery,
		map[string]interface{}{
			"name":        req.Resource.Name(),
			"rootName":    model.RootResource.Name(),
			"composition": model.CompositionRelationship}
}

const cDeleteResourceCypher = `
MATCH (r:Resource{name: $name})
WITH r
// nadji sve resurse za brisanje
CALL {
    WITH r
    OPTIONAL MATCH (r)-[:Includes*{kind: $composition}]->(d:Resource)
    RETURN collect(d) + collect(r) AS delRes
}
// obrisi sve atribute resursa za brisanje
CALL {
    WITH delRes
    UNWIND delRes AS r
    MATCH (r)-[:Includes{kind: $composition}]->(a:Attribute)
    DETACH DELETE a
}
// obrisi sve direktno dodeljene dozvole resursa za brisanje
CALL {
    WITH delRes
    UNWIND delRes AS r
    MATCH (r)-[:Has|On{priority: 0}]-(p:Permission{})
    DETACH DELETE p
}
// nadji sve resurse kojima treba obrisati dozvole (prvi nivo)
CALL {
    WITH delRes
    MATCH (p:Resource)-[:Includes]->(c:Resource)
    WHERE p IN delRes AND NOT c IN delRes
    RETURN collect({parent: p, child: c}) AS permRes
}
// nadji sve resurse koje treba povezati sa root-om (prvi nivo)
CALL {
    WITH delRes
    MATCH (p:Resource)-[:Includes]->(c:Resource)
    WHERE p IN delRes AND NOT c IN delRes AND NOT EXISTS
        {((c)<-[:Includes]-(op:Resource)) WHERE NOT (op IN delRes)}
    RETURN collect(DISTINCT c) AS rootRes
}
MATCH (root:Resource{name: $rootName})
WITH *
// // nadji sve root dozvole (subjekat)
CALL {
    MATCH (root)-[rel:Has]->(p:Permission)
    RETURN collect({priority: rel.priority, permission: p}) AS sRootPerms
}
// // nadji sve root dozvole (objekat)
CALL {
    MATCH (root)<-[rel:On]-(p:Permission)
    RETURN collect({priority: rel.priority, permission: p}) AS oRootPerms
}
// ukloni nasledjene dozvole iz potomaka
CALL {
    WITH permRes
    UNWIND permRes AS pr
    WITH pr.parent AS parent, pr.child AS child
    // nadji potomke trenutnog resursa (ukljucen i sam resurs)
    CALL {
        WITH child
        OPTIONAL MATCH path=(child)-[:Includes*]->(d:Resource)
        RETURN collect({resource: d, distance: length(path) + 1}) AS descendants
        UNION
        WITH child
        MATCH (c)
        WHERE c = child
        RETURN collect({resource: c, distance: 1}) AS descendants
    }
    // obrisi dozvole iz potomaka (subjekat)
    CALL {
        WITH descendants, parent
        UNWIND descendants AS d
        WITH d.resource AS r, d.distance AS dist, parent
        MATCH (parent)-[psrel:Has]->(p:Permission)
        MATCH (r)-[dsrel:Has{priority: psrel.priority - dist}]->(p)
        WITH dsrel.priority as priority, collect(dsrel) AS del
        UNWIND del[..1] AS d
        DELETE d
    }
    // obrisi dozvole iz potomaka (objekat)
    CALL {
        WITH descendants, parent
        UNWIND descendants AS d
        WITH d.resource AS r, d.distance AS dist, parent
        MATCH (parent)<-[psrel:On]-(p:Permission)
        MATCH (r)<-[dsrel:On{priority: psrel.priority - dist}]-(p)
        WITH dsrel.priority as priority, collect(dsrel) AS del
        UNWIND del[..1] AS d
        DELETE d
    }
}
// obrisi resurse
CALL {
    WITH delRes
    UNWIND delRes AS r
    DETACH DELETE r
}
// povezati resurse sa root-om
// i njima i svim njihovim potomcima dodeliti dozvole
CALL {
    WITH rootRes, sRootPerms, oRootPerms, root
    UNWIND rootRes AS rr
    CREATE (root)-[:Includes{kind: $composition}]->(rr)
    WITH rr, sRootPerms, oRootPerms, root
    // nadji potomke trenutnog resursa (ukljucen i sam resurs)
    CALL {
        WITH rr
        WITH rr AS child
        MATCH path=(child)-[:Includes*]->(d:Resource)
        RETURN collect({resource: d, distance: length(path) + 1}) AS descendants
        UNION
        MATCH (child)
        RETURN collect({resource: child, distance: 1}) AS descendants
    }
   // dodeli potomcima root dozvole (sa strane subjekta)
    CALL {
        WITH descendants, sRootPerms
        UNWIND descendants AS d
        WITH d.resource AS r, d.distance AS dist, sRootPerms
        UNWIND sRootPerms AS rp
        WITH rp.priority AS priority, rp.permission AS p, r, dist
        CREATE (r)-[:Has{priority: priority - dist}]->(p)
    }
   // dodeli potomcima root dozvole (sa strane objekta)
    CALL {
        WITH descendants, oRootPerms
        UNWIND descendants AS d
        WITH d.resource AS r, d.distance AS dist, oRootPerms
        UNWIND oRootPerms AS rp
        WITH rp.priority AS priority, rp.permission AS p, r, dist
        CREATE (r)<-[:On{priority: priority - dist}]-(p)
    }
}
`

func (f cachedPermissionsCypherFactory) deleteResourceCypher(req acl.DeleteResourceReq) (string, map[string]interface{}) {
	return cDeleteResourceCypher,
		map[string]interface{}{
			"name":        req.Resource.Name(),
			"composition": model.CompositionRelationship,
			"rootName":    model.RootResource.Name()}
}

const cGetResourceQuery = `
MATCH (resource:Resource{name: $name})
OPTIONAL MATCH (attr:Attribute)<-[:Includes{kind: $composition}]-(resource)
RETURN resource.name, collect(properties(attr)) as attrs
`

func (f cachedPermissionsCypherFactory) getResourceCypher(req acl.GetResourceReq) (string, map[string]interface{}) {
	return cGetResourceQuery,
		map[string]interface{}{
			"name":        req.Resource.Name(),
			"composition": model.CompositionRelationship}
}

const cCreateAttributeQuery = `
MATCH (r:Resource{name: $name})
WHERE NOT ((r)-[:Includes{kind: $composition}]->(:Attribute{name: $attrName}))
CREATE (r)-[:Includes{kind: $composition}]->(:Attribute{name: $attrName, kind: $attrKind, value: $attrValue})
`

func (f cachedPermissionsCypherFactory) createAttributeCypher(req acl.CreateAttributeReq) (string, map[string]interface{}) {
	return cCreateAttributeQuery,
		map[string]interface{}{
			"name":        req.Resource.Name(),
			"attrName":    req.Attribute.Name(),
			"attrKind":    req.Attribute.Kind(),
			"attrValue":   req.Attribute.Value(),
			"composition": model.CompositionRelationship}
}

const cUpdateAttributeQuery = `
MATCH ((:Resource{name: $name})-[:Includes{kind: $composition}]->(a:Attribute{name: $attrName, kind: $attrKind}))
SET a.value = $attrValue
`

func (f cachedPermissionsCypherFactory) updateAttributeCypher(req acl.UpdateAttributeReq) (string, map[string]interface{}) {
	return cUpdateAttributeQuery,
		map[string]interface{}{
			"name":        req.Resource.Name(),
			"attrName":    req.Attribute.Name(),
			"attrKind":    req.Attribute.Kind(),
			"attrValue":   req.Attribute.Value(),
			"composition": model.CompositionRelationship}
}

const cDeleteAttributeQuery = `
MATCH ((:Resource{name: $name})-[:Includes{kind: $composition}]->(a:Attribute{name: $attrName}))
DETACH DELETE a
`

func (f cachedPermissionsCypherFactory) deleteAttributeCypher(req acl.DeleteAttributeReq) (string, map[string]interface{}) {
	return cDeleteAttributeQuery,
		map[string]interface{}{
			"name":        req.Resource.Name(),
			"attrName":    req.AttributeId,
			"composition": model.CompositionRelationship}
}

const cCreateRelQuery = `
MATCH (parent:Resource{name: $parentName})
MATCH (child:Resource{name: $childName})
WHERE NOT (parent)-[:Includes]->(child) AND NOT (child)-[:Includes*]->(parent)
CREATE (parent)-[:Includes{kind: $relKind}]->(child)
// nadji sve dozvole koje treba da se naslede
WITH parent, child
CALL {
    WITH parent
    MATCH (parent)-[srel:Has|On]-(p:Permission)
    RETURN collect({priority: srel.priority, type: type(srel), permission: p}) AS rels
}
// nadji potomke
CALL {
    WITH child
    MATCH path=(child)-[:Includes*]->(d:Resource)
    RETURN collect({resource: d, distance: length(path) + 1}) + collect({resource: child, distance: 1}) AS descendants
}
//  dodeli dozvole potomcima
CALL {
    WITH descendants, rels
    UNWIND descendants AS d
    WITH d.resource AS r, d.distance AS dist, rels
    UNWIND rels AS rel
    WITH rel.priority AS priority, rel.type AS type, r.permission AS p, r, dist
    FOREACH (i in CASE WHEN type = "Has" THEN [1] ELSE [] END |
        CREATE (r)-[:Has{priority: priority - dist}]->(p)
    )
    FOREACH (i in CASE WHEN type = "On" THEN [1] ELSE [] END |
        CREATE (r)<-[:On{priority: priority - dist}]-(p)
    )
}
// RASKIDANJE VEZE SA ROOT-OM
CALL {
    WITH *
    MATCH (child)<-[rootRel:Includes{kind: $composition}]-(root:Resource{name: $rootName})
    DELETE rootRel
    WITH *
    // nadji sve root dozvole koje treba da se uklone
    CALL {
        WITH root
        MATCH (root)-[srel:Has|On]-(p:Permission)
        RETURN collect({priority: srel.priority, type: type(srel), permission: p}) AS rootRels
    }
    CALL {
        WITH descendants, rootRels
        WITH descendants, rootRels AS rels
        UNWIND descendants AS d
        WITH d.resource AS r, d.distance AS dist, rels
        UNWIND rels AS rel
        WITH rel.priority AS priority, rel.type AS type, r.permission AS p, r, dist
        CALL {
            WITH r, priority, dist, p, type
            MATCH (r)-[dsrel:Has{priority: priority - dist}]->(p)
            WHERE type(dsrel) = type
            WITH dsrel.priority as priority, collect(dsrel) AS del
            UNWIND del[..1] AS d
            DELETE d
        }
        CALL {
            WITH r, priority, dist, p, type
            MATCH (r)<-[dsrel:On{priority: priority - dist}]-(p)
            WHERE type(dsrel) = type
            WITH dsrel.priority as priority, collect(dsrel) AS del
            UNWIND del[..1] AS d
            DELETE d
        }
    }
}
`

func (f cachedPermissionsCypherFactory) createAggregationRelCypher(req acl.CreateAggregationRelReq) (string, map[string]interface{}) {
	return cCreateRelQuery,
		map[string]interface{}{
			"parentName":  req.Parent.Name(),
			"childName":   req.Child.Name(),
			"relKind":     model.AggregateRelationship,
			"composition": model.CompositionRelationship,
			"rootName":    model.RootResource.Name()}
}

const cDeleteRelQuery = `
MATCH (parent:Resource{name: $parentName})-[includes:Includes{kind: $relKind}]->(child:Resource{name: $childName})
WITH parent, child, includes
// nadji sve dozvole koje ima roditelj (kao subjekat)
CALL {
    WITH parent
    MATCH (parent)-[rel:Has]->(p:Permission)
    RETURN collect({permission: p, priority: rel.priority}) AS subPermissions
}
// nadji sve dozvole koje ima roditelj (kao objekat)
CALL {
    WITH parent
    MATCH (parent)<-[rel:On]-(p:Permission)
    RETURN collect({permission: p, priority: rel.priority}) AS objPermissions
}
// pronadji potomke deteta i njihovu udaljenost od roditelja
// i spoji sa detetom i njegovom udaljenoscu od roditelja
CALL {
    WITH child
    MATCH path=(child)-[:Includes*]->(d:Resource)
    RETURN collect({resource: d, distance: length(path) + 1}) + collect({resource: child, distance: 1}) AS descendants
}
// obrisi sve nasledjene dozvole
// iz deteta i njegovih potomaka
CALL {
    WITH parent, child, subPermissions, objPermissions, descendants
    // obrisi iz potomaka (sa strane subjekta)
    CALL {
        WITH parent, subPermissions, descendants
        UNWIND descendants AS d
        WITH d.resource AS child, d.distance AS dist, subPermissions
        UNWIND subPermissions AS prel
        WITH prel.priority AS priority, prel.permission AS permission, child, dist
        MATCH (child)-[crel:Has{priority: priority - dist}]->(permission)
        WITH crel.priority as priority, collect(crel) AS del
        UNWIND del[..1] AS d
        DELETE d
    }
    // obrisi iz potomaka (sa strane objekta)
    CALL {
        WITH parent, objPermissions, descendants
        UNWIND descendants AS d
        WITH d.resource AS child, d.distance AS dist, objPermissions
        UNWIND objPermissions AS prel
        WITH prel.priority AS priority, prel.permission AS permission, child, dist
        MATCH (child)<-[crel:On{priority: priority - dist}]-(permission)
        WITH crel.priority as priority, collect(crel) AS del
        UNWIND del[..1] AS d
        DELETE d
    }
}
// obrisi vezu izmedju roditelja i deteta
CALL {
    WITH includes
    DELETE includes
}
// povezi child sa root-om ako nema drugih roditelja
// i dodeli njemu i njegovim potomcima root dozvole
CALL {
    WITH child, descendants
    MATCH (child)
    WHERE NOT (child)<-[:Includes]-()
    MATCH (root:Resource{name: $rootName})
    CREATE (root)-[:Includes{kind: $composition}]->(child)
    WITH child, root, descendants
    // dodeli potomcima root dozvole (sa strane subjekta)
    CALL {
        WITH descendants, root
        UNWIND descendants AS d
        WITH d.resource AS r, d.distance AS dist, root
        MATCH (root)-[srel:Has]->(p:Permission)
        MERGE (r)-[:Has{priority: srel.prioriy - dist}]->(p)
    }
    // dodeli potomcima root dozvole (sa strane objekta)
    CALL {
        WITH descendants, root
        UNWIND descendants AS d
        WITH d.resource AS r, d.distance AS dist, root
        MATCH (root)<-[srel:On]-(p:Permission)
        MERGE (r)<-[:On{priority: srel.prioriy - dist}]->(p)
    }
}
`

func (f cachedPermissionsCypherFactory) deleteAggregationRelCypher(req acl.DeleteAggregationRelReq) (string, map[string]interface{}) {
	return cDeleteRelQuery,
		map[string]interface{}{
			"parentName":  req.Parent.Name(),
			"childName":   req.Child.Name(),
			"relKind":     model.AggregateRelationship,
			"composition": model.CompositionRelationship,
			"rootName":    model.RootResource.Name()}
}

func (f cachedPermissionsCypherFactory) createCompositionRelCypher(req acl.CreateCompositionRelReq) (string, map[string]interface{}) {
	return cCreateRelQuery,
		map[string]interface{}{
			"parentName":  req.Parent.Name(),
			"childName":   req.Child.Name(),
			"relKind":     model.CompositionRelationship,
			"composition": model.CompositionRelationship,
			"rootName":    model.RootResource.Name()}
}

func (f cachedPermissionsCypherFactory) deleteCompositionRelCypher(req acl.DeleteCompositionRelReq) (string, map[string]interface{}) {
	return cDeleteRelQuery,
		map[string]interface{}{
			"parentName":  req.Parent.Name(),
			"childName":   req.Child.Name(),
			"relKind":     model.CompositionRelationship,
			"composition": model.CompositionRelationship,
			"rootName":    model.RootResource.Name()}
}

const cCreatePermissionQuery = `
MATCH (sub:Resource{name: $subName})
MATCH (obj:Resource{name: $objName})
WHERE NOT (sub)-[:Has{priority: 0}]->(:Permission{name: $permName, kind: $permKind})-[:On{priority: 0}]->(obj)
CREATE (sub)-[srel:Has{priority: 0}]->(p:Permission{name: $permName, kind: $permKind, condition: $permCond})-[orel:On{priority: 0}]->(obj)
WITH sub, obj, p
CALL {
    WITH sub, p
    MATCH path=((sub)-[:Includes*]->(descendant:Resource))
    CREATE (descendant)-[:Has{priority: -length(path)}]->(p)
}
CALL {
    WITH obj, p
    MATCH path=((obj)-[:Includes*]->(descendant:Resource))
    CREATE (descendant)<-[:On{priority: -length(path)}]-(p)
}
`

func (f cachedPermissionsCypherFactory) createPermissionCypher(req acl.CreatePermissionReq) (string, map[string]interface{}) {
	return cCreatePermissionQuery,
		map[string]interface{}{
			"subName":  req.Subject.Name(),
			"objName":  req.Object.Name(),
			"permName": req.Permission.Name(),
			"permKind": req.Permission.Kind(),
			"permCond": req.Permission.Condition().Expression()}
}

const cDeletePermissionQuery = `
MATCH (sub:Resource{name: $subName})
MATCH (obj:Resource{name: $objName})
MATCH (sub)-[:Has{priority: 0}]->(p:Permission{name: $permName, kind: $permKind})-[:On{priority: 0}]->(obj)
DETACH DELETE p
`

func (f cachedPermissionsCypherFactory) deletePermissionCypher(req acl.DeletePermissionReq) (string, map[string]interface{}) {
	return cDeletePermissionQuery,
		map[string]interface{}{
			"subName":  req.Subject.Name(),
			"objName":  req.Object.Name(),
			"permName": req.Permission.Name(),
			"permKind": req.Permission.Kind()}
}

const cGetPermissionsQuery = `
MATCH (sub:Resource{name: $subName})
MATCH (obj:Resource{name: $objName})
MATCH (sub)-[srel:Has]->(p:Permission{name: $permName})-[orel:On]->(obj)
RETURN p.name, p.kind, p.condition, srel.priority, orel.priority
`

func (f cachedPermissionsCypherFactory) getEffectivePermissionsWithPriorityCypher(req acl.GetPermissionHierarchyReq) (string, map[string]interface{}) {
	return cGetPermissionsQuery,
		map[string]interface{}{
			"subName":  req.Subject.Name(),
			"objName":  req.Object.Name(),
			"permName": req.PermissionName}
}
