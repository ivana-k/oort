package handler

import (
	"github.com/c12s/oort/domain/model"
	"github.com/c12s/oort/domain/model/checker"
	"github.com/c12s/oort/domain/store"
)

type CheckerHandler struct {
	store store.AclStore
}

func NewCheckerHandler(store store.AclStore) CheckerHandler {
	return CheckerHandler{
		store: store,
	}
}

func (h CheckerHandler) CheckPermission(req checker.CheckPermissionReq) checker.CheckPermissionResp {
	resp := h.store.GetPermissionByPrecedence(checker.GetPermissionReq{
		Principal:      req.Principal,
		Resource:       req.Resource,
		PermissionName: req.PermissionName,
	})

	if resp.Error != nil {
		return checker.CheckPermissionResp{
			Allowed: false,
			Error:   resp.Error,
		}
	}

	principalAttrs := h.store.GetAttributes(checker.GetAttributeReq{
		Resource: req.Principal,
	})
	if principalAttrs.Error != nil {
		return checker.CheckPermissionResp{
			Allowed: false,
			Error:   principalAttrs.Error,
		}
	}
	resourceAttrs := h.store.GetAttributes(checker.GetAttributeReq{
		Resource: req.Resource,
	})
	if resourceAttrs.Error != nil {
		return checker.CheckPermissionResp{
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

	return checker.CheckPermissionResp{
		Allowed: allowed,
		Error:   nil,
	}
}
