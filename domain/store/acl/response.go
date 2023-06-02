package acl

import (
	"errors"
	"github.com/c12s/oort/domain/model"
)

var (
	ErrNotFound = errors.New("resource not found")
)

type SyncResp struct {
	Error error
}

type GetAttributeResp struct {
	Attributes []model.Attribute
	Error      error
}

type GetResourceResp struct {
	Resource *model.Resource
	Error    error
}

type GetPermissionHierarchyResp struct {
	Hierarchy model.PermissionHierarchy
	Error     error
}
