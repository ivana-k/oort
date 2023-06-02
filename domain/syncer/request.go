package syncer

import "github.com/c12s/oort/domain/model"

type Request interface {
	Id() string
}

type CreateResourceReq struct {
	ReqId    string
	Resource model.Resource
}

func (r CreateResourceReq) Id() string {
	return r.ReqId
}

type DeleteResourceReq struct {
	ReqId    string
	Resource model.Resource
}

func (r DeleteResourceReq) Id() string {
	return r.ReqId
}

type CreateAttributeReq struct {
	ReqId     string
	Resource  model.Resource
	Attribute model.Attribute
}

func (r CreateAttributeReq) Id() string {
	return r.ReqId
}

type UpdateAttributeReq struct {
	ReqId     string
	Resource  model.Resource
	Attribute model.Attribute
}

func (r UpdateAttributeReq) Id() string {
	return r.ReqId
}

type DeleteAttributeReq struct {
	ReqId       string
	Resource    model.Resource
	AttributeId model.AttributeId
}

func (r DeleteAttributeReq) Id() string {
	return r.ReqId
}

type CreateAggregationRelReq struct {
	ReqId string
	Parent,
	Child model.Resource
}

func (r CreateAggregationRelReq) Id() string {
	return r.ReqId
}

type DeleteAggregationRelReq struct {
	ReqId string
	Parent,
	Child model.Resource
}

func (r DeleteAggregationRelReq) Id() string {
	return r.ReqId
}

type CreateCompositionRelReq struct {
	ReqId string
	Parent,
	Child model.Resource
}

func (r CreateCompositionRelReq) Id() string {
	return r.ReqId
}

type DeleteCompositionRelReq struct {
	ReqId string
	Parent,
	Child model.Resource
}

func (r DeleteCompositionRelReq) Id() string {
	return r.ReqId
}

type CreatePermissionReq struct {
	ReqId string
	Subject,
	Object *model.Resource
	Permission model.Permission
}

func (r CreatePermissionReq) Id() string {
	return r.ReqId
}

type DeletePermissionReq struct {
	ReqId string
	Subject,
	Object *model.Resource
	Permission model.Permission
}

func (r DeletePermissionReq) Id() string {
	return r.ReqId
}
