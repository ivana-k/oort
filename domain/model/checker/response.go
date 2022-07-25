package checker

import "github.com/c12s/oort/domain/model"

type GetAttributeResp struct {
	Attributes []model.Attribute
	Error      error
}

type GetPermissionByPrecedenceResp struct {
	Hierarchy model.PermissionHierarchy
	Error     error
}

type CheckPermissionResp struct {
	Allowed bool
	Error   error
}
