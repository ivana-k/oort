package checker

import (
	"fmt"
	"github.com/c12s/oort/domain/model"
	"github.com/c12s/oort/domain/store/acl"
	"github.com/c12s/oort/domain/store/cache"
)

type Handler struct {
	store               acl.Store
	cache               cache.Cache
	attrSerializer      AttributeSerializer
	checkRespSerializer CheckPermissionResponseSerializer
}

func NewHandler(store acl.Store, cache cache.Cache, attrSerializer AttributeSerializer, checkRespSerializer CheckPermissionResponseSerializer) Handler {
	return Handler{
		store:               store,
		cache:               cache,
		attrSerializer:      attrSerializer,
		checkRespSerializer: checkRespSerializer,
	}
}

func (h Handler) CheckPermission(req CheckPermissionReq) CheckPermissionResp {
	if value, err := h.cache.Get(checkRespCacheKey(req)); err == nil {
		resp, err := h.checkRespSerializer.Deserialize(value)
		if err != nil {
			return resp
		}
	}

	resp := h.store.GetPermissionHierarchy(acl.GetPermissionHierarchyReq{
		Subject:        req.Subject,
		Object:         req.Object,
		PermissionName: req.PermissionName,
	})
	if resp.Error != nil {
		return errorResponse(resp.Error)
	}

	principalAttrs, err := h.getAttributes(req.Subject)
	if err != nil {
		return errorResponse(err)
	}
	resourceAttrs, err := h.getAttributes(req.Object)
	if err != nil {
		return errorResponse(err)
	}

	evalReq := model.PermissionEvalRequest{
		Principal: principalAttrs,
		Resource:  resourceAttrs,
		Env:       req.Env,
	}
	evalResult := resp.Hierarchy.Eval(evalReq)

	checkResp := CheckPermissionResp{
		Allowed: allowed(evalResult),
		Error:   nil,
	}

	if value, err := h.checkRespSerializer.Serialize(checkResp); err == nil {
		_ = h.cache.Set(checkRespCacheKey(req), value, []string{})
	}

	return checkResp
}

func (h Handler) getAttributes(resource model.Resource) ([]model.Attribute, error) {
	if value, err := h.cache.Get(attrCacheKey(resource)); err == nil {
		attrs, err := h.attrSerializer.Deserialize(value)
		if err == nil {
			return attrs, nil
		}
	}

	res := h.store.GetResource(acl.GetResourceReq{Resource: resource})
	if res.Error != nil {
		return nil, res.Error
	}

	if bytes, err := h.attrSerializer.Serialize(res.Resource.Attributes); err == nil {
		_ = h.cache.Set(attrCacheKey(resource), bytes, []string{})
	}

	return res.Resource.Attributes, nil
}

func checkRespCacheKey(req CheckPermissionReq) string {
	return fmt.Sprintf("%s/%s/%s",
		req.Subject.Name(),
		req.Object.Name(),
		req.PermissionName)
}

func attrCacheKey(resource model.Resource) string {
	return resource.Name()
}

func errorResponse(err error) CheckPermissionResp {
	return CheckPermissionResp{
		Allowed: false,
		Error:   err,
	}
}

func allowed(result model.EvalResult) bool {
	if result == model.EvalResultAllowed {
		return true
	}
	return false
}
