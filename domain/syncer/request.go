package syncer

import "github.com/c12s/oort/domain/model"

type Request interface {
	Id() string
}

type ConnectResourcesReq struct {
	ReqId string
	Parent,
	Child model.Resource
}

func (r ConnectResourcesReq) Id() string {
	return r.ReqId
}

type DisconnectResourcesReq struct {
	ReqId string
	Parent,
	Child model.Resource
}

func (r DisconnectResourcesReq) Id() string {
	return r.ReqId
}

type UpsertAttributeReq struct {
	ReqId     string
	Resource  model.Resource
	Attribute model.Attribute
}

func (r UpsertAttributeReq) Id() string {
	return r.ReqId
}

type RemoveAttributeReq struct {
	ReqId       string
	Resource    model.Resource
	AttributeId model.AttributeId
}

func (r RemoveAttributeReq) Id() string {
	return r.ReqId
}

type InsertPermissionReq struct {
	ReqId string
	Principal,
	Resource model.Resource
	Permission model.Permission
}

func (r InsertPermissionReq) Id() string {
	return r.ReqId
}

type RemovePermissionReq struct {
	ReqId string
	Principal,
	Resource model.Resource
	Permission model.Permission
}

func (r RemovePermissionReq) Id() string {
	return r.ReqId
}
