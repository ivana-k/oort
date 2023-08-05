package neo4j

import (
	"errors"
	"github.com/c12s/oort/domain/model"
	"github.com/c12s/oort/domain/store/acl"
)

type AclStore struct {
	manager *TransactionManager
	factory CypherFactory
}

func NewAclStore(manager *TransactionManager, factory CypherFactory) acl.Store {
	return AclStore{
		manager: manager,
		factory: factory,
	}
}

func (store AclStore) CreateResource(req acl.CreateResourceReq) acl.SyncResp {
	cypher1, params1 := store.factory.createResource(req)
	idAttr := model.NewAttribute(model.NewAttributeId("id"), model.String, req.Resource.Id())
	idAttrReq := acl.PutAttributeReq{Resource: req.Resource, Attribute: idAttr}
	cypher2, params2 := store.factory.putAttribute(idAttrReq)
	kindAttr := model.NewAttribute(model.NewAttributeId("kind"), model.String, req.Resource.Kind())
	kindAttrReq := acl.PutAttributeReq{Resource: req.Resource, Attribute: kindAttr}
	cypher3, params3 := store.factory.putAttribute(kindAttrReq)
	cyphers := []string{cypher1, cypher2, cypher3}
	params := []map[string]interface{}{params1, params2, params3}
	err := store.manager.WriteTransactions(cyphers, params, req.Callback)
	return acl.SyncResp{Error: err}
}

func (store AclStore) DeleteResource(req acl.DeleteResourceReq) acl.SyncResp {
	cypher, params := store.factory.deleteResource(req)
	err := store.manager.WriteTransaction(cypher, params, req.Callback)
	return acl.SyncResp{Error: err}
}

func (store AclStore) GetResource(req acl.GetResourceReq) acl.GetResourceResp {
	cypher, params := store.factory.getResource(req)
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

// todo: if the resource is being created, assign default attrs to it
func (store AclStore) PutAttribute(req acl.PutAttributeReq) acl.SyncResp {
	cypher, params := store.factory.putAttribute(req)
	err := store.manager.WriteTransaction(cypher, params, req.Callback)
	return acl.SyncResp{Error: err}
}

func (store AclStore) DeleteAttribute(req acl.DeleteAttributeReq) acl.SyncResp {
	cypher, params := store.factory.deleteAttribute(req)
	err := store.manager.WriteTransaction(cypher, params, req.Callback)
	return acl.SyncResp{Error: err}
}

// todo: if the resource is being created, assign default attrs to it
func (store AclStore) CreateInheritanceRel(req acl.CreateInheritanceRelReq) acl.SyncResp {
	cypher, params := store.factory.createInheritanceRel(req)
	err := store.manager.WriteTransaction(cypher, params, req.Callback)
	return acl.SyncResp{Error: err}
}

func (store AclStore) DeleteInheritanceRel(req acl.DeleteInheritanceRelReq) acl.SyncResp {
	cypher, params := store.factory.deleteInheritanceRel(req)
	err := store.manager.WriteTransaction(cypher, params, req.Callback)
	return acl.SyncResp{Error: err}
}

// todo: if the resource is being created, assign default attrs to it
func (store AclStore) CreatePolicy(req acl.CreatePolicyReq) acl.SyncResp {
	cypher, params := store.factory.createPolicy(req)
	err := store.manager.WriteTransaction(cypher, params, req.Callback)
	return acl.SyncResp{Error: err}
}

func (store AclStore) DeletePolicy(req acl.DeletePolicyReq) acl.SyncResp {
	cypher, params := store.factory.deletePolicy(req)
	err := store.manager.WriteTransaction(cypher, params, req.Callback)
	return acl.SyncResp{Error: err}
}

func (store AclStore) GetPermissionHierarchy(req acl.GetPermissionHierarchyReq) acl.GetPermissionHierarchyResp {
	cypher, params := store.factory.getEffectivePermissionsWithPriority(req)
	records, err := store.manager.ReadTransaction(cypher, params)
	if err != nil {
		return acl.GetPermissionHierarchyResp{Hierarchy: nil, Error: err}
	}

	hierarchy, err := getHierarchy(records)
	return acl.GetPermissionHierarchyResp{Hierarchy: hierarchy, Error: err}
}
