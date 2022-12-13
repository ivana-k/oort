package neo4j

import (
	"github.com/c12s/oort/domain/model"
	"github.com/c12s/oort/domain/store/acl"
)

type AclStore struct {
	manager *TransactionManager
}

func NewAclStore(manager *TransactionManager) acl.Store {
	return AclStore{
		manager: manager,
	}
}

func (store AclStore) CreateResource(req acl.CreateResourceReq) acl.SyncResp {
	cypher, params := createResourceCypher(req)
	err := store.manager.WriteTransaction(cypher, params, req.Callback)
	return acl.SyncResp{Error: err}
}

func (store AclStore) DeleteResource(req acl.DeleteResourceReq) acl.SyncResp {
	cypher, params := deleteResourceCypher(req)
	err := store.manager.WriteTransaction(cypher, params, req.Callback)
	return acl.SyncResp{Error: err}
}

func (store AclStore) GetResource(req acl.GetResourceReq) acl.GetResourceResp {
	cypher, params := getResourceCypher(req)
	records, err := store.manager.ReadTransaction(cypher, params)
	if err != nil {
		return acl.GetResourceResp{Resource: nil, Error: err}
	}

	results := records.([]interface{})[0].([]interface{})[0]

	return acl.GetResourceResp{Resource: getResource(results), Error: nil}
}

func (store AclStore) CreateAttribute(req acl.CreateAttributeReq) acl.SyncResp {
	cypher, params := createAttributeCypher(req)
	err := store.manager.WriteTransaction(cypher, params, req.Callback)
	return acl.SyncResp{Error: err}
}

func (store AclStore) UpdateAttribute(req acl.UpdateAttributeReq) acl.SyncResp {
	cypher, params := updateAttributeCypher(req)
	err := store.manager.WriteTransaction(cypher, params, req.Callback)
	return acl.SyncResp{Error: err}
}

func (store AclStore) DeleteAttribute(req acl.DeleteAttributeReq) acl.SyncResp {
	cypher, params := deleteAttributeCypher(req)
	err := store.manager.WriteTransaction(cypher, params, req.Callback)
	return acl.SyncResp{Error: err}
}

func (store AclStore) GetAttributes(req acl.GetAttributeReq) acl.GetAttributeResp {
	cypher, params := getAttributeCypher(req)
	records, err := store.manager.ReadTransaction(cypher, params)
	if err != nil {
		return acl.GetAttributeResp{Attributes: nil, Error: err}
	}

	results := make([]interface{}, 0)
	for _, record := range records.([]interface{}) {
		results = append(results, record.([]interface{})[0])
	}

	return acl.GetAttributeResp{Attributes: getAttributes(results), Error: nil}
}

func (store AclStore) CreateAggregationRel(req acl.CreateAggregationRelReq) acl.SyncResp {
	cypher, params := createAggregationRelCypher(req)
	err := store.manager.WriteTransaction(cypher, params, req.Callback)
	return acl.SyncResp{Error: err}
}

func (store AclStore) DeleteAggregationRel(req acl.DeleteAggregationRelReq) acl.SyncResp {
	cypher, params := deleteAggregationRelCypher(req)
	err := store.manager.WriteTransaction(cypher, params, req.Callback)
	return acl.SyncResp{Error: err}
}

func (store AclStore) CreateCompositionRel(req acl.CreateCompositionRelReq) acl.SyncResp {
	cypher, params := createCompositionRelCypher(req)
	err := store.manager.WriteTransaction(cypher, params, req.Callback)
	return acl.SyncResp{Error: err}
}

func (store AclStore) DeleteCompositionRel(req acl.DeleteCompositionRelReq) acl.SyncResp {
	cypher, params := deleteCompositionRelCypher(req)
	err := store.manager.WriteTransaction(cypher, params, req.Callback)
	return acl.SyncResp{Error: err}
}

func (store AclStore) CreatePermission(req acl.CreatePermissionReq) acl.SyncResp {
	cypher, params := createPermissionCypher(req)
	err := store.manager.WriteTransaction(cypher, params, req.Callback)
	return acl.SyncResp{Error: err}
}

func (store AclStore) DeletePermission(req acl.DeletePermissionReq) acl.SyncResp {
	cypher, params := deletePermissionCypher(req)
	err := store.manager.WriteTransaction(cypher, params, req.Callback)
	return acl.SyncResp{Error: err}
}

func (store AclStore) GetPermissionHierarchy(req acl.GetPermissionHierarchyReq) acl.GetPermissionHierarchyResp {
	cypher, params := getEffectivePermissionsWithPriorityCypher(req)
	records, err := store.manager.ReadTransaction(cypher, params)
	if err != nil {
		return acl.GetPermissionHierarchyResp{Hierarchy: nil, Error: err}
	}

	permsByPriority := make(map[int][]model.Permission)
	for _, result := range records.([]interface{}) {
		priority := int(result.([]interface{})[1].(int64))
		_, ok := permsByPriority[priority]
		if !ok {
			permsByPriority[priority] = make([]model.Permission, 0)
		}
		permission, err := getPermission(result.([]interface{})[0])
		if err != nil {
			return acl.GetPermissionHierarchyResp{Hierarchy: nil, Error: err}
		}
		permsByPriority[priority] = append(permsByPriority[priority], permission)
	}

	return acl.GetPermissionHierarchyResp{Hierarchy: getHierarchy(permsByPriority), Error: nil}
}
