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
		fmt.Println(err)
		fmt.Println(resp.Error)
		fmt.Println(resp)
		return resp
	}

	resp := h.store.GetPermissionByPrecedence(acl.GetPermissionReq{
		Principal:      req.Principal,
		Resource:       req.Resource,
		PermissionName: req.PermissionName,
	})

	if resp.Error != nil {
		return CheckPermissionResp{
			Allowed: false,
			Error:   resp.Error,
		}
	}

	principalAttrs, err := h.getAttributes(req.Principal)
	if err != nil {
		return CheckPermissionResp{
			Allowed: false,
			Error:   err,
		}
	}
	resourceAttrs, err := h.getAttributes(req.Resource)
	if err != nil {
		return CheckPermissionResp{
			Allowed: false,
			Error:   err,
		}
	}

	result := resp.Hierarchy.Eval(principalAttrs.Attributes, resourceAttrs.Attributes, req.Env)
	var allowed bool
	if result == model.Allowed {
		allowed = true
	} else {
		allowed = false
	}

	checkResp := CheckPermissionResp{
		Allowed: allowed,
		Error:   nil,
	}
	if value, err := h.checkRespSerializer.Serialize(checkResp); err == nil {
		_ = h.cache.Set(checkRespCacheKey(req), value, []string{})
	}
	return checkResp
}

func checkRespCacheKey(req CheckPermissionReq) string {
	return fmt.Sprintf("%s/%s/%s", req.Principal.Id(), req.Resource.Id(), req.PermissionName)
}

func attrCacheKey(resource model.Resource) string {
	fmt.Println(resource.Id())
	return resource.Id()
}

func (h Handler) getAttributes(resource model.Resource) (acl.GetAttributeResp, error) {
	var resourceAttrs acl.GetAttributeResp
	if value, err := h.cache.Get(attrCacheKey(resource)); err == nil {
		resourceAttrs.Attributes, err = h.attrSerializer.Deserialize(value)
		if err != nil {
			resourceAttrs.Attributes = nil
		}
	}
	if resourceAttrs.Attributes == nil {
		resourceAttrs = h.store.GetAttributes(acl.GetAttributeReq{
			Resource: resource,
		})
		if resourceAttrs.Error != nil {
			return acl.GetAttributeResp{}, resourceAttrs.Error
		}
		if bytes, err := h.attrSerializer.Serialize(resourceAttrs.Attributes); err == nil {
			_ = h.cache.Set(attrCacheKey(resource), bytes, []string{})
		}
	}
	return resourceAttrs, nil
}
