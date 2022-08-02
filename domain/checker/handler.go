package checker

import (
	"github.com/c12s/oort/domain/model"
	"github.com/c12s/oort/domain/store/acl"
)

type Handler struct {
	store acl.Store
}

func NewHandler(store acl.Store) Handler {
	return Handler{
		store: store,
	}
}

func (h Handler) CheckPermission(req CheckPermissionReq) CheckPermissionResp {
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

	principalAttrs := h.store.GetAttributes(acl.GetAttributeReq{
		Resource: req.Principal,
	})
	if principalAttrs.Error != nil {
		return CheckPermissionResp{
			Allowed: false,
			Error:   principalAttrs.Error,
		}
	}
	resourceAttrs := h.store.GetAttributes(acl.GetAttributeReq{
		Resource: req.Resource,
	})
	if resourceAttrs.Error != nil {
		return CheckPermissionResp{
			Allowed: false,
			Error:   resourceAttrs.Error,
		}
	}

	result := resp.Hierarchy.Eval(principalAttrs.Attributes, resourceAttrs.Attributes, req.Env)
	var allowed bool
	if result == model.Allowed {
		allowed = true
	} else {
		allowed = false
	}

	return CheckPermissionResp{
		Allowed: allowed,
		Error:   nil,
	}
}
