package checker

import "github.com/c12s/oort/domain/model"

type GetAttributeReq struct {
	Resource model.Resource
}

type GetPermissionReq struct {
	Principal,
	Resource model.Resource
	PermissionName string
}

type CheckPermissionReq struct {
	Principal,
	Resource model.Resource
	PermissionName string
	Env            map[string]interface{}
}
