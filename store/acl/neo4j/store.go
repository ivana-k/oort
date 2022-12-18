package neo4j

import (
	"errors"
	"github.com/c12s/oort/domain/model"
	"github.com/c12s/oort/domain/store/acl"
)

type AclStore struct {
	manager *TransactionManager
	factory cypherFactory
}

func NewAclStore(manager *TransactionManager, factory cypherFactory) acl.Store {
	return AclStore{
		manager: manager,
		factory: factory,
	}
}

func (store AclStore) CreateResource(req acl.CreateResourceReq) acl.SyncResp {
	cypher1, params1 := store.factory.createResourceCypher(req)
	idAttr := model.NewAttribute(model.NewAttributeId("id"), model.String, req.Resource.Id())
	idAttrReq := acl.CreateAttributeReq{Resource: req.Resource, Attribute: idAttr}
	cypher2, params2 := store.factory.createAttributeCypher(idAttrReq)
	kindAttr := model.NewAttribute(model.NewAttributeId("kind"), model.String, req.Resource.Kind())
	kindAttrReq := acl.CreateAttributeReq{Resource: req.Resource, Attribute: kindAttr}
	cypher3, params3 := store.factory.createAttributeCypher(kindAttrReq)
	cyphers := []string{cypher1, cypher2, cypher3}
	params := []map[string]interface{}{params1, params2, params3}
	err := store.manager.WriteTransactions(cyphers, params, req.Callback)
	return acl.SyncResp{Error: err}
}

func (store AclStore) DeleteResource(req acl.DeleteResourceReq) acl.SyncResp {
	cypher, params := store.factory.deleteResourceCypher(req)
	err := store.manager.WriteTransaction(cypher, params, req.Callback)
	return acl.SyncResp{Error: err}
}

func (store AclStore) GetResource(req acl.GetResourceReq) acl.GetResourceResp {
	cypher, params := store.factory.getResourceCypher(req)
	records, err := store.manager.ReadTransaction(cypher, params)
	if err != nil {
		return acl.GetResourceResp{Resource: nil, Error: err}
	}

	recordList, ok := records.([]interface{})
	if !ok {
		return acl.GetResourceResp{Error: errors.New("invalid resp format")}
	}
	if len(recordList) == 0 {
		return acl.GetResourceResp{Error: errors.New("resource not found")}
	}
	resourceRecord := records.([]interface{})[0]
	resourceAttrs, ok := resourceRecord.([]interface{})
	if !ok {
		return acl.GetResourceResp{Error: errors.New("invalid resp format")}
	}

	return acl.GetResourceResp{Resource: getResource(resourceAttrs), Error: nil}
}

func (store AclStore) CreateAttribute(req acl.CreateAttributeReq) acl.SyncResp {
	cypher, params := store.factory.createAttributeCypher(req)
	err := store.manager.WriteTransaction(cypher, params, req.Callback)
	return acl.SyncResp{Error: err}
}

func (store AclStore) UpdateAttribute(req acl.UpdateAttributeReq) acl.SyncResp {
	cypher, params := store.factory.updateAttributeCypher(req)
	err := store.manager.WriteTransaction(cypher, params, req.Callback)
	return acl.SyncResp{Error: err}
}

func (store AclStore) DeleteAttribute(req acl.DeleteAttributeReq) acl.SyncResp {
	cypher, params := store.factory.deleteAttributeCypher(req)
	err := store.manager.WriteTransaction(cypher, params, req.Callback)
	return acl.SyncResp{Error: err}
}

func (store AclStore) CreateAggregationRel(req acl.CreateAggregationRelReq) acl.SyncResp {
	cypher, params := store.factory.createAggregationRelCypher(req)
	err := store.manager.WriteTransaction(cypher, params, req.Callback)
	return acl.SyncResp{Error: err}
}

func (store AclStore) DeleteAggregationRel(req acl.DeleteAggregationRelReq) acl.SyncResp {
	cypher, params := store.factory.deleteAggregationRelCypher(req)
	err := store.manager.WriteTransaction(cypher, params, req.Callback)
	return acl.SyncResp{Error: err}
}

func (store AclStore) CreateCompositionRel(req acl.CreateCompositionRelReq) acl.SyncResp {
	cypher, params := store.factory.createCompositionRelCypher(req)
	err := store.manager.WriteTransaction(cypher, params, req.Callback)
	return acl.SyncResp{Error: err}
}

func (store AclStore) DeleteCompositionRel(req acl.DeleteCompositionRelReq) acl.SyncResp {
	cypher, params := store.factory.deleteCompositionRelCypher(req)
	err := store.manager.WriteTransaction(cypher, params, req.Callback)
	return acl.SyncResp{Error: err}
}

func (store AclStore) CreatePermission(req acl.CreatePermissionReq) acl.SyncResp {
	cypher, params := store.factory.createPermissionCypher(req)
	err := store.manager.WriteTransaction(cypher, params, req.Callback)
	return acl.SyncResp{Error: err}
}

func (store AclStore) DeletePermission(req acl.DeletePermissionReq) acl.SyncResp {
	cypher, params := store.factory.deletePermissionCypher(req)
	err := store.manager.WriteTransaction(cypher, params, req.Callback)
	return acl.SyncResp{Error: err}
}

func (store AclStore) GetPermissionHierarchy(req acl.GetPermissionHierarchyReq) acl.GetPermissionHierarchyResp {
	cypher, params := store.factory.getEffectivePermissionsWithPriorityCypher(req)
	records, err := store.manager.ReadTransaction(cypher, params)
	if err != nil {
		return acl.GetPermissionHierarchyResp{Hierarchy: nil, Error: err}
	}

	hierarchy, err := getHierarchy(records)
	return acl.GetPermissionHierarchyResp{Hierarchy: hierarchy, Error: err}
}
