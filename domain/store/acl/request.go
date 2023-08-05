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

type GetResourceReq struct {
	Resource model.Resource
}

type PutAttributeReq struct {
	Resource  model.Resource
	Attribute model.Attribute
	Callback  func(error) model.OutboxMessage
}

type DeleteAttributeReq struct {
	Resource    model.Resource
	AttributeId model.AttributeId
	Callback    func(error) model.OutboxMessage
}

type GetAttributeReq struct {
	Resource model.Resource
}

type CreateInheritanceRelReq struct {
	From     model.Resource
	To       model.Resource
	Callback func(err error) model.OutboxMessage
}

type DeleteInheritanceRelReq struct {
	From     model.Resource
	To       model.Resource
	Callback func(err error) model.OutboxMessage
}

type CreatePolicyReq struct {
	SubjectScope,
	ObjectScope model.Resource
	Permission model.Permission
	Callback   func(error) model.OutboxMessage
}

type DeletePolicyReq struct {
	SubjectScope,
	ObjectScope model.Resource
	Permission model.Permission
	Callback   func(error) model.OutboxMessage
}

type GetPermissionHierarchyReq struct {
	Subject,
	Object model.Resource
	PermissionName string
}
