package acl

import (
	"github.com/c12s/oort/domain/model"
)

type CreateResourceReq struct {
	Resource model.Resource
	Callback func(error) model.OutboxMessage
}

type DeleteResourceReq struct {
	Resource model.Resource
	Callback func(error) model.OutboxMessage
}

type GetAttributeReq struct {
	Resource model.Resource
}

type GetPermissionHierarchyReq struct {
	Subject,
	Object model.Resource
	PermissionName string
}

type GetResourceReq struct {
	Resource model.Resource
}

type ConnectResourcesReq struct {
	Parent,
	Child model.Resource
	Callback func(error) model.OutboxMessage
}

type DisconnectResourcesReq struct {
	Parent,
	Child model.Resource
	Callback func(error) model.OutboxMessage
}

type CreateAttributeReq struct {
	Resource  model.Resource
	Attribute model.Attribute
	Callback  func(error) model.OutboxMessage
}

type UpdateAttributeReq struct {
	Resource  model.Resource
	Attribute model.Attribute
	Callback  func(error) model.OutboxMessage
}

type DeleteAttributeReq struct {
	Resource    model.Resource
	AttributeId model.AttributeId
	Callback    func(error) model.OutboxMessage
}

type CreateAggregationRelReq struct {
	Parent,
	Child model.Resource
	Callback func(error) model.OutboxMessage
}

type DeleteAggregationRelReq struct {
	Parent,
	Child model.Resource
	Callback func(error) model.OutboxMessage
}

type CreateCompositionRelReq struct {
	Parent,
	Child model.Resource
	Callback func(error) model.OutboxMessage
}

type DeleteCompositionRelReq struct {
	Parent,
	Child model.Resource
	Callback func(error) model.OutboxMessage
}

type CreatePermissionReq struct {
	Subject,
	Object model.Resource
	Permission model.Permission
	Callback   func(error) model.OutboxMessage
}

type DeletePermissionReq struct {
	Subject,
	Object model.Resource
	Permission model.Permission
	Callback   func(error) model.OutboxMessage
}
