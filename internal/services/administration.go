package services

import (
	"github.com/c12s/oort/internal/domain"
)

type AdministrationService struct {
	repo domain.RHABACRepo
}

func NewAdministrationService(repo domain.RHABACRepo) (*AdministrationService, error) {
	return &AdministrationService{
		repo: repo,
	}, nil
}

func (h AdministrationService) CreateResource(req domain.CreateResourceReq) domain.AdministrationResp {
	return h.repo.CreateResource(req)
}

func (h AdministrationService) DeleteResource(req domain.DeleteResourceReq) domain.AdministrationResp {
	return h.repo.DeleteResource(req)
}

func (h AdministrationService) PutAttribute(req domain.PutAttributeReq) domain.AdministrationResp {
	return h.repo.PutAttribute(req)
}

func (h AdministrationService) DeleteAttribute(req domain.DeleteAttributeReq) domain.AdministrationResp {
	return h.repo.DeleteAttribute(req)
}

func (h AdministrationService) CreateInheritanceRel(req domain.CreateInheritanceRelReq) domain.AdministrationResp {
	return h.repo.CreateInheritanceRel(req)
}

func (h AdministrationService) DeleteInheritanceRel(req domain.DeleteInheritanceRelReq) domain.AdministrationResp {
	return h.repo.DeleteInheritanceRel(req)
}

func (h AdministrationService) CreatePolicy(req domain.CreatePolicyReq) domain.AdministrationResp {
	if req.SubjectScope.Name() == "" {
		req.SubjectScope = domain.RootResource
	}
	if req.ObjectScope.Name() == "" {
		req.ObjectScope = domain.RootResource
	}
	return h.repo.CreatePolicy(req)
}

func (h AdministrationService) DeletePolicy(req domain.DeletePolicyReq) domain.AdministrationResp {
	if req.SubjectScope.Name() == "" {
		req.SubjectScope = domain.RootResource
	}
	if req.ObjectScope.Name() == "" {
		req.ObjectScope = domain.RootResource
	}
	return h.repo.DeletePolicy(req)
}
