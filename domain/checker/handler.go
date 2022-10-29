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

	resp := h.store.GetPermissionByPrecedence(acl.GetPermissionReq{
		Principal:      req.Principal,
		Resource:       req.Resource,
		PermissionName: req.PermissionName,
	})
	if resp.Error != nil {
		return errorResponse(resp.Error)
	}

	principalAttrs := h.getAttributes(req.Principal)
	if principalAttrs.Error != nil {
		return errorResponse(principalAttrs.Error)
	}
	resourceAttrs := h.getAttributes(req.Resource)
	if resourceAttrs.Error != nil {
		return errorResponse(resourceAttrs.Error)
	}

	evalReq := model.PermissionEvalRequest{
		Principal: principalAttrs.Attributes,
		Resource:  resourceAttrs.Attributes,
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

func (h Handler) getAttributes(resource model.Resource) acl.GetAttributeResp {
	if value, err := h.cache.Get(attrCacheKey(resource)); err == nil {
		attrs, err := h.attrSerializer.Deserialize(value)
		if err == nil {
			return acl.GetAttributeResp{
				Attributes: attrs,
				Error:      nil,
			}
		}
	}

	attrs := h.store.GetAttributes(acl.GetAttributeReq{Resource: resource})
	if attrs.Error != nil {
		return attrs
	}

	if bytes, err := h.attrSerializer.Serialize(attrs.Attributes); err == nil {
		_ = h.cache.Set(attrCacheKey(resource), bytes, []string{})
	}

	return attrs
}

func checkRespCacheKey(req CheckPermissionReq) string {
	return fmt.Sprintf("%s/%s/%s",
		req.Principal.Id(),
		req.Resource.Id(),
		req.PermissionName)
}

func attrCacheKey(resource model.Resource) string {
	return resource.Id()
}

func errorResponse(err error) CheckPermissionResp {
	return CheckPermissionResp{
		Allowed: false,
		Error:   err,
	}
}

func allowed(result model.EvalResult) bool {
	if result == model.Allowed {
		return true
	}
	return false
}
