package neo4j

import (
	"errors"
	"github.com/c12s/oort/internal/domain"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type RHABACRepo struct {
	manager *TransactionManager
	factory CypherFactory
}

func NewRHABACRepo(manager *TransactionManager, factory CypherFactory) domain.RHABACRepo {
	return RHABACRepo{
		manager: manager,
		factory: factory,
	}
}

func (store RHABACRepo) CreateResource(req domain.CreateResourceReq) domain.AdministrationResp {
	cypher1, params1 := store.factory.createResource(req)
	idAttrId, err := domain.NewAttributeId("id")
	if err != nil {
		return domain.AdministrationResp{
			Error: err,
		}
	}
	idAttr, err := domain.NewAttribute(*idAttrId, domain.String, req.Resource.Id())
	if err != nil {
		return domain.AdministrationResp{
			Error: err,
		}
	}
	idAttrReq := domain.PutAttributeReq{Resource: req.Resource, Attribute: *idAttr}
	cypher2, params2 := store.factory.putAttribute(idAttrReq)
	kindAttrId, err := domain.NewAttributeId("kind")
	if err != nil {
		return domain.AdministrationResp{
			Error: err,
		}
	}
	kindAttr, err := domain.NewAttribute(*kindAttrId, domain.String, req.Resource.Kind())
	if err != nil {
		return domain.AdministrationResp{
			Error: err,
		}
	}
	kindAttrReq := domain.PutAttributeReq{Resource: req.Resource, Attribute: *kindAttr}
	cypher3, params3 := store.factory.putAttribute(kindAttrReq)
	cyphers := []string{cypher1, cypher2, cypher3}
	params := []map[string]interface{}{params1, params2, params3}
	err = store.manager.WriteTransactions(cyphers, params)
	return domain.AdministrationResp{Error: err}
}

func (store RHABACRepo) DeleteResource(req domain.DeleteResourceReq) domain.AdministrationResp {
	cypher, params := store.factory.deleteResource(req)
	err := store.manager.WriteTransaction(cypher, params)
	return domain.AdministrationResp{Error: err}
}

func (store RHABACRepo) GetResource(req domain.GetResourceReq) domain.GetResourceResp {
	cypher, params := store.factory.getResource(req)
	records, err := store.manager.ReadTransaction(cypher, params)
	if err != nil {
		return domain.GetResourceResp{Resource: nil, Error: err}
	}

	recordList, ok := records.([]*neo4j.Record)
	if !ok {
		return domain.GetResourceResp{Error: errors.New("invalid resp format")}
	}
	if len(recordList) == 0 {
		return domain.GetResourceResp{Error: errors.New("resource not found")}
	}
	return domain.GetResourceResp{Resource: getResource(records), Error: nil}
}

// todo: if the resource is being created, assign default attrs to it
func (store RHABACRepo) PutAttribute(req domain.PutAttributeReq) domain.AdministrationResp {
	cypher, params := store.factory.putAttribute(req)
	err := store.manager.WriteTransaction(cypher, params)
	return domain.AdministrationResp{Error: err}
}

func (store RHABACRepo) DeleteAttribute(req domain.DeleteAttributeReq) domain.AdministrationResp {
	cypher, params := store.factory.deleteAttribute(req)
	err := store.manager.WriteTransaction(cypher, params)
	return domain.AdministrationResp{Error: err}
}

// todo: if the resource is being created, assign default attrs to it
func (store RHABACRepo) CreateInheritanceRel(req domain.CreateInheritanceRelReq) domain.AdministrationResp {
	cypher, params := store.factory.createInheritanceRel(req)
	err := store.manager.WriteTransaction(cypher, params)
	return domain.AdministrationResp{Error: err}
}

func (store RHABACRepo) DeleteInheritanceRel(req domain.DeleteInheritanceRelReq) domain.AdministrationResp {
	cypher, params := store.factory.deleteInheritanceRel(req)
	err := store.manager.WriteTransaction(cypher, params)
	return domain.AdministrationResp{Error: err}
}

// todo: if the resource is being created, assign default attrs to it
func (store RHABACRepo) CreatePolicy(req domain.CreatePolicyReq) domain.AdministrationResp {
	cypher, params := store.factory.createPolicy(req)
	err := store.manager.WriteTransaction(cypher, params)
	return domain.AdministrationResp{Error: err}
}

func (store RHABACRepo) DeletePolicy(req domain.DeletePolicyReq) domain.AdministrationResp {
	cypher, params := store.factory.deletePolicy(req)
	err := store.manager.WriteTransaction(cypher, params)
	return domain.AdministrationResp{Error: err}
}

func (store RHABACRepo) GetPermissionHierarchy(req domain.GetPermissionHierarchyReq) domain.GetPermissionHierarchyResp {
	cypher, params := store.factory.getEffectivePermissionsWithPriority(req)
	records, err := store.manager.ReadTransaction(cypher, params)
	if err != nil {
		return domain.GetPermissionHierarchyResp{Hierarchy: nil, Error: err}
	}

	hierarchy, err := getHierarchy(records)
	return domain.GetPermissionHierarchyResp{Hierarchy: hierarchy, Error: err}
}
