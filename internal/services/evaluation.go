package services

import (
	"fmt"
	"github.com/c12s/oort/internal/domain"
)

type EvaluationService struct {
	store domain.RHABACStore
	//attrMarshaller      domain.AttributesMarshaller
	//authZRespMarshaller domain.AuthorizationRespMarshaller
}

func NewEvaluationService(store domain.RHABACStore) (*EvaluationService, error) {
	return &EvaluationService{
		store: store,
	}, nil
}

func (h EvaluationService) Authorize(req domain.AuthorizationReq) domain.AuthorizationResp {
	//if value, err := h.cache.Get(checkRespCacheKey(req)); err == nil {
	//	resp, err := h.checkRespSerializer.Deserialize(value)
	//	if err != nil {
	//		return resp
	//	}
	//}

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
		Allowed: allowed(evalResult),
		Error:   nil,
	}

	//if value, err := h.checkRespSerializer.Serialize(checkResp); err == nil {
	//	_ = h.cache.Set(checkRespCacheKey(req), value, []string{})
	//}

	return checkResp
}

func (h EvaluationService) getAttributes(resource domain.Resource) ([]domain.Attribute, error) {
	//if value, err := h.cache.Get(attrCacheKey(resource)); err == nil {
	//	attrs, err := h.attrSerializer.Deserialize(value)
	//	if err == nil {
	//		return attrs, nil
	//	}
	//}

	res := h.store.GetResource(domain.GetResourceReq{Resource: resource})
	if res.Error != nil {
		return nil, res.Error
	}

	//if bytes, err := h.attrSerializer.Serialize(res.Resource.Attributes); err == nil {
	//	_ = h.cache.Set(attrCacheKey(resource), bytes, []string{})
	//}

	return res.Resource.Attributes, nil
}

func checkRespCacheKey(req domain.AuthorizationReq) string {
	return fmt.Sprintf("%s/%s/%s",
		req.Subject.Name(),
		req.Object.Name(),
		req.PermissionName)
}

func attrCacheKey(resource domain.Resource) string {
	return resource.Name()
}

func errorResponse(err error) domain.AuthorizationResp {
	return domain.AuthorizationResp{
		Allowed: false,
		Error:   err,
	}
}

func allowed(result domain.EvalResult) bool {
	if result == domain.EvalResultAllowed {
		return true
	}
	return false
}
