package acl

import "github.com/c12s/oort/domain/model"

type SyncResp struct {
	Error error
}

type GetAttributeResp struct {
	Attributes []model.Attribute
	Error      error
}

type GetPermissionByPrecedenceResp struct {
	Hierarchy model.PermissionHierarchy
	Error     error
}
