package checker

import "github.com/c12s/oort/domain/model"

type CheckPermissionReq struct {
	Subject,
	Object model.Resource
	PermissionName string
	Env            []model.Attribute
}
