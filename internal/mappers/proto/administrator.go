package proto

import (
	"github.com/c12s/oort/internal/domain"
	"github.com/c12s/oort/pkg/api"
)

func CreateResourceReqToDomain(req *api.CreateResourceReq) (*domain.CreateResourceReq, error) {
	resource, err := ResourceToDomain(req.Resource)
	if err != nil {
		return nil, err
	}
	return &domain.CreateResourceReq{
		Resource: *resource,
	}, nil
}

func DeleteResourceReqToDomain(req *api.DeleteResourceReq) (*domain.DeleteResourceReq, error) {
	resource, err := ResourceToDomain(req.Resource)
	if err != nil {
		return nil, err
	}
	return &domain.DeleteResourceReq{
		Resource: *resource,
	}, nil
}

func PutAttributeReqToDomain(req *api.PutAttributeReq) (*domain.PutAttributeReq, error) {
	resource, err := ResourceToDomain(req.Resource)
	if err != nil {
		return nil, err
	}
	attr, err := AttributeToDomain(req.Attribute)
	if err != nil {
		return nil, err
	}
	return &domain.PutAttributeReq{
		Resource:  *resource,
		Attribute: *attr,
	}, nil
}

func DeleteAttributeReqToDomain(req *api.DeleteAttributeReq) (*domain.DeleteAttributeReq, error) {
	resource, err := ResourceToDomain(req.Resource)
	if err != nil {
		return nil, err
	}
	attrId, err := AttributeIdToDomain(req.AttributeId)
	if err != nil {
		return nil, err
	}
	return &domain.DeleteAttributeReq{
		Resource:    *resource,
		AttributeId: *attrId,
	}, nil
}

func CreateInheritanceRelReqToDomain(req *api.CreateInheritanceRelReq) (*domain.CreateInheritanceRelReq, error) {
	from, err := ResourceToDomain(req.From)
	if err != nil {
		return nil, err
	}
	to, err := ResourceToDomain(req.To)
	if err != nil {
		return nil, err
	}
	return &domain.CreateInheritanceRelReq{
		From: *from,
		To:   *to,
	}, nil
}

func DeleteInheritanceRelReqToDomain(req *api.DeleteInheritanceRelReq) (*domain.DeleteInheritanceRelReq, error) {
	from, err := ResourceToDomain(req.From)
	if err != nil {
		return nil, err
	}
	to, err := ResourceToDomain(req.To)
	if err != nil {
		return nil, err
	}
	return &domain.DeleteInheritanceRelReq{
		From: *from,
		To:   *to,
	}, nil
}

func CreatePolicyReqToDomain(req *api.CreatePolicyReq) (*domain.CreatePolicyReq, error) {
	permission, err := PermissionToDomain(req.Permission)
	if err != nil {
		return nil, err
	}
	subScope, err := ResourceToDomain(req.SubjectScope)
	if err != nil {
		return nil, err
	}
	objScope, err := ResourceToDomain(req.ObjectScope)
	if err != nil {
		return nil, err
	}
	return &domain.CreatePolicyReq{
		SubjectScope: *subScope,
		ObjectScope:  *objScope,
		Permission:   *permission,
	}, nil
}

func DeletePolicyReqToDomain(req *api.DeletePolicyReq) (*domain.DeletePolicyReq, error) {
	permission, err := PermissionToDomain(req.Permission)
	if err != nil {
		return nil, err
	}
	subScope, err := ResourceToDomain(req.SubjectScope)
	if err != nil {
		return nil, err
	}
	objScope, err := ResourceToDomain(req.ObjectScope)
	if err != nil {
		return nil, err
	}
	return &domain.DeletePolicyReq{
		SubjectScope: *subScope,
		ObjectScope:  *objScope,
		Permission:   *permission,
	}, nil
}

func AdministrationAsyncRespFromDomain(resp domain.AdministrationResp) (*api.AdministrationAsyncResp, error) {
	err := ""
	if resp.Error != nil {
		err = resp.Error.Error()
	}
	return &api.AdministrationAsyncResp{
		Error: err,
	}, nil
}
