package services

import (
	"github.com/c12s/oort/internal/domain"
)

type EvaluationService struct {
	store domain.RHABACStore
}

func NewEvaluationService(store domain.RHABACStore) (*EvaluationService, error) {
	return &EvaluationService{
		store: store,
	}, nil
}

func (h EvaluationService) Authorize(req domain.AuthorizationReq) domain.AuthorizationResp {
	resp := h.store.GetPermissionHierarchy(domain.GetPermissionHierarchyReq{
		Subject:        req.Subject,
		Object:         req.Object,
		PermissionName: req.PermissionName,
	})
	if resp.Error != nil {
		return errorResponse(resp.Error)
	}

	subAttrs, err := h.getAttributes(req.Subject)
	if err != nil {
		return errorResponse(err)
	}
	objAttrs, err := h.getAttributes(req.Object)
	if err != nil {
		return errorResponse(err)
	}

	evalReq := domain.PermissionEvalRequest{
		Subject: subAttrs,
		Object:  objAttrs,
		Env:     req.Env,
	}
	evalResult := resp.Hierarchy.Eval(evalReq)

	checkResp := domain.AuthorizationResp{
		Authorized: authorized(evalResult),
		Error:      nil,
	}

	return checkResp
}

func (h EvaluationService) getAttributes(resource domain.Resource) ([]domain.Attribute, error) {
	res := h.store.GetResource(domain.GetResourceReq{Resource: resource})
	if res.Error != nil {
		return nil, res.Error
	}
	return res.Resource.Attributes, nil
}

func errorResponse(err error) domain.AuthorizationResp {
	return domain.AuthorizationResp{
		Authorized: false,
		Error:      err,
	}
}

func authorized(result domain.EvalResult) bool {
	if result == domain.EvalResultAllowed {
		return true
	}
	return false
}
