package services

import (
	"github.com/c12s/oort/internal/domain"
)

type AdministrationService struct {
	store domain.RHABACStore
}

func NewAdministrationService(store domain.RHABACStore) (*AdministrationService, error) {
	return &AdministrationService{
		store: store,
	}, nil
}

func (h AdministrationService) CreateResource(req domain.CreateResourceReq) domain.AdministrationResp {
	return h.store.CreateResource(req, h.outboxMsgGenerator(req.Id))
}

func (h AdministrationService) DeleteResource(req domain.DeleteResourceReq) domain.AdministrationResp {
	return h.store.DeleteResource(req, h.outboxMsgGenerator(req.Id))
}

func (h AdministrationService) PutAttribute(req domain.PutAttributeReq) domain.AdministrationResp {
	return h.store.PutAttribute(req, h.outboxMsgGenerator(req.Id))
}

func (h AdministrationService) DeleteAttribute(req domain.DeleteAttributeReq) domain.AdministrationResp {
	return h.store.DeleteAttribute(req, h.outboxMsgGenerator(req.Id))
}

func (h AdministrationService) CreateInheritanceRel(req domain.CreateInheritanceRelReq) domain.AdministrationResp {
	return h.store.CreateInheritanceRel(req, h.outboxMsgGenerator(req.Id))
}

func (h AdministrationService) DeleteInheritanceRel(req domain.DeleteInheritanceRelReq) domain.AdministrationResp {
	return h.store.DeleteInheritanceRel(req, h.outboxMsgGenerator(req.Id))
}

func (h AdministrationService) CreatePolicy(req domain.CreatePolicyReq) domain.AdministrationResp {
	if req.SubjectScope.Name() == "" {
		req.SubjectScope = domain.RootResource
	}
	if req.ObjectScope.Name() == "" {
		req.ObjectScope = domain.RootResource
	}
	return h.store.CreatePolicy(req, h.outboxMsgGenerator(req.Id))
}

func (h AdministrationService) DeletePolicy(req domain.DeletePolicyReq) domain.AdministrationResp {
	if req.SubjectScope.Name() == "" {
		req.SubjectScope = domain.RootResource
	}
	if req.ObjectScope.Name() == "" {
		req.ObjectScope = domain.RootResource
	}
	return h.store.DeletePolicy(req, h.outboxMsgGenerator(req.Id))
}

func (h AdministrationService) outboxMsgGenerator(reqId string) domain.OutboxMsgGenerator {
	return func(err error) *domain.OutboxMessage {
		//if err != nil {
		//	return h.respFactory(reqId, err.Error(), false)
		//}
		//return h.respFactory(reqId, "", true)
		return nil
	}
}
