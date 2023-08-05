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

type PutAttributeReq struct {
	ReqId     string
	Resource  model.Resource
	Attribute model.Attribute
}

func (r PutAttributeReq) Id() string {
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

type CreateInheritanceRelReq struct {
	ReqId string
	From,
	To model.Resource
}

func (r CreateInheritanceRelReq) Id() string {
	return r.ReqId
}

type DeleteInheritanceRelReq struct {
	ReqId string
	From,
	To model.Resource
}

func (r DeleteInheritanceRelReq) Id() string {
	return r.ReqId
}

type CreatePolicyReq struct {
	ReqId string
	SubjectScope,
	ObjectScope *model.Resource
	Permission model.Permission
}

func (r CreatePolicyReq) Id() string {
	return r.ReqId
}

type DeletePolicyReq struct {
	ReqId string
	SubjectScope,
	ObjectScope *model.Resource
	Permission model.Permission
}

func (r DeletePolicyReq) Id() string {
	return r.ReqId
}
