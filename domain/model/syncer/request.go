package syncer

import "github.com/c12s/oort/domain/model"

type Request interface{}

type ConnectResourcesReq struct {
	Parent,
	Child model.Resource
}

type DisconnectResourcesReq struct {
	Parent,
	Child model.Resource
}

type UpsertAttributeReq struct {
	Resource  model.Resource
	Attribute model.Attribute
}

type RemoveAttributeReq struct {
	Resource    model.Resource
	AttributeId model.AttributeId
}

type InsertPermissionReq struct {
	Principal,
	Resource model.Resource
	Permission model.Permission
}

type RemovePermissionReq struct {
	Principal,
	Resource model.Resource
	Permission model.Permission
}
