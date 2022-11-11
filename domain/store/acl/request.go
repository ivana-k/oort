package acl

import (
	"github.com/c12s/oort/domain/model"
)

type GetAttributeReq struct {
	Resource model.Resource
}

type GetPermissionReq struct {
	Principal,
	Resource model.Resource
	PermissionName string
}

type GetResourceReq struct {
	Id   string
	Kind string
}

type ConnectResourcesReq struct {
	Parent,
	Child model.Resource
	Callback func(error) *model.OutboxMessage
}

type DisconnectResourcesReq struct {
	Parent,
	Child model.Resource
	Callback func(error) *model.OutboxMessage
}

type UpsertAttributeReq struct {
	Resource  model.Resource
	Attribute model.Attribute
	Callback  func(error) *model.OutboxMessage
}

type RemoveAttributeReq struct {
	Resource    model.Resource
	AttributeId model.AttributeId
	Callback    func(error) *model.OutboxMessage
}

type InsertPermissionReq struct {
	Principal,
	Resource model.Resource
	Permission model.Permission
	Callback   func(error) *model.OutboxMessage
}

type RemovePermissionReq struct {
	Principal,
	Resource model.Resource
	Permission model.Permission
	Callback   func(error) *model.OutboxMessage
}
