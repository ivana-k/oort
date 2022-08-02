package checker

import "github.com/c12s/oort/domain/model"

type CheckPermissionReq struct {
	Principal,
	Resource model.Resource
	PermissionName string
	Env            map[string]interface{}
}
