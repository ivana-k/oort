package proto

import (
	"errors"
	"github.com/c12s/oort/internal/domain"
	"github.com/golang/protobuf/proto"
)

func (x *CreateResourceReq) ToDomain() (*domain.CreateResourceReq, error) {
	resource, err := x.Resource.ToDomain()
	if err != nil {
		return nil, err
	}
	return &domain.CreateResourceReq{
		Id:       x.Id,
		Resource: *resource,
	}, nil
}

func (x *CreateResourceReq) FromDomain(req domain.CreateResourceReq) (*CreateResourceReq, error) {
	protoRes := &Resource{}
	protoRes, err := protoRes.FromDomain(req.Resource)
	if err != nil {
		return nil, err
	}
	return &CreateResourceReq{
		Id:       req.Id,
		Resource: protoRes,
	}, nil
}

func (x *DeleteResourceReq) ToDomain() (*domain.DeleteResourceReq, error) {
	resource, err := x.Resource.ToDomain()
	if err != nil {
		return nil, err
	}
	return &domain.DeleteResourceReq{
		Id:       x.Id,
		Resource: *resource,
	}, nil
}

func (x *DeleteResourceReq) FromDomain(req domain.DeleteResourceReq) (*DeleteResourceReq, error) {
	protoRes := &Resource{}
	protoRes, err := protoRes.FromDomain(req.Resource)
	if err != nil {
		return nil, err
	}
	return &DeleteResourceReq{
		Id:       req.Id,
		Resource: protoRes,
	}, nil
}

func (x *PutAttributeReq) ToDomain() (*domain.PutAttributeReq, error) {
	resource, err := x.Resource.ToDomain()
	if err != nil {
		return nil, err
	}
	attr, err := x.Attribute.ToDomain()
	if err != nil {
		return nil, err
	}
	return &domain.PutAttributeReq{
		Id:        x.Id,
		Resource:  *resource,
		Attribute: *attr,
	}, nil
}

func (x *PutAttributeReq) FromDomain(req domain.PutAttributeReq) (*PutAttributeReq, error) {
	resource := &Resource{}
	resource, err := resource.FromDomain(req.Resource)
	if err != nil {
		return nil, err
	}
	attr := &Attribute{}
	attr, err = attr.FromDomain(req.Attribute)
	if err != nil {
		return nil, err
	}
	return &PutAttributeReq{
		Id:        req.Id,
		Resource:  resource,
		Attribute: attr,
	}, nil
}

func (x *DeleteAttributeReq) ToDomain() (*domain.DeleteAttributeReq, error) {
	resource, err := x.Resource.ToDomain()
	if err != nil {
		return nil, err
	}
	attrId, err := x.AttributeId.ToDomain()
	if err != nil {
		return nil, err
	}
	return &domain.DeleteAttributeReq{
		Id:          x.Id,
		Resource:    *resource,
		AttributeId: *attrId,
	}, nil
}

func (x *DeleteAttributeReq) FromDomain(req domain.DeleteAttributeReq) (*DeleteAttributeReq, error) {
	resource := &Resource{}
	resource, err := resource.FromDomain(req.Resource)
	if err != nil {
		return nil, err
	}
	attrId := &AttributeId{}
	attrId, err = attrId.FromDomain(req.AttributeId)
	if err != nil {
		return nil, err
	}
	return &DeleteAttributeReq{
		Id:          req.Id,
		Resource:    resource,
		AttributeId: attrId,
	}, nil
}

func (x *CreateInheritanceRelReq) ToDomain() (*domain.CreateInheritanceRelReq, error) {
	from, err := x.From.ToDomain()
	if err != nil {
		return nil, err
	}
	to, err := x.To.ToDomain()
	if err != nil {
		return nil, err
	}
	return &domain.CreateInheritanceRelReq{
		Id:   x.Id,
		From: *from,
		To:   *to,
	}, nil
}

func (x *CreateInheritanceRelReq) FromDomain(req domain.CreateInheritanceRelReq) (*CreateInheritanceRelReq, error) {
	from := &Resource{}
	from, err := from.FromDomain(req.From)
	if err != nil {
		return nil, err
	}
	to := &Resource{}
	to, err = to.FromDomain(req.To)
	if err != nil {
		return nil, err
	}
	return &CreateInheritanceRelReq{
		Id:   req.Id,
		From: from,
		To:   to,
	}, nil
}

func (x *DeleteInheritanceRelReq) ToDomain() (*domain.DeleteInheritanceRelReq, error) {
	from, err := x.From.ToDomain()
	if err != nil {
		return nil, err
	}
	to, err := x.To.ToDomain()
	if err != nil {
		return nil, err
	}
	return &domain.DeleteInheritanceRelReq{
		Id:   x.Id,
		From: *from,
		To:   *to,
	}, nil
}

func (x *DeleteInheritanceRelReq) FromDomain(req domain.DeleteInheritanceRelReq) (*DeleteInheritanceRelReq, error) {
	from := &Resource{}
	from, err := from.FromDomain(req.From)
	if err != nil {
		return nil, err
	}
	to := &Resource{}
	to, err = to.FromDomain(req.To)
	if err != nil {
		return nil, err
	}
	return &DeleteInheritanceRelReq{
		Id:   req.Id,
		From: from,
		To:   to,
	}, nil
}

func (x *CreatePolicyReq) ToDomain() (*domain.CreatePolicyReq, error) {
	permission, err := x.Permission.ToDomain()
	if err != nil {
		return nil, err
	}
	subScope, err := x.SubjectScope.ToDomain()
	if err != nil {
		return nil, err
	}
	objScope, err := x.ObjectScope.ToDomain()
	if err != nil {
		return nil, err
	}
	return &domain.CreatePolicyReq{
		Id:           x.Id,
		SubjectScope: *subScope,
		ObjectScope:  *objScope,
		Permission:   *permission,
	}, nil
}

func (x *CreatePolicyReq) FromDomain(req domain.CreatePolicyReq) (*CreatePolicyReq, error) {
	permission := &Permission{}
	permission, err := permission.FromDomain(req.Permission)
	if err != nil {
		return nil, err
	}
	subScope := &Resource{}
	subScope, err = subScope.FromDomain(req.SubjectScope)
	if err != nil {
		return nil, err
	}
	objScope := &Resource{}
	objScope, err = objScope.FromDomain(req.ObjectScope)
	if err != nil {
		return nil, err
	}
	return &CreatePolicyReq{
		Id:           req.Id,
		SubjectScope: subScope,
		ObjectScope:  objScope,
		Permission:   permission,
	}, nil
}

func (x *DeletePolicyReq) ToDomain() (*domain.DeletePolicyReq, error) {
	permission, err := x.Permission.ToDomain()
	if err != nil {
		return nil, err
	}
	subScope, err := x.SubjectScope.ToDomain()
	if err != nil {
		return nil, err
	}
	objScope, err := x.ObjectScope.ToDomain()
	if err != nil {
		return nil, err
	}
	return &domain.DeletePolicyReq{
		Id:           x.Id,
		SubjectScope: *subScope,
		ObjectScope:  *objScope,
		Permission:   *permission,
	}, nil
}

func (x *DeletePolicyReq) FromDomain(req domain.DeletePolicyReq) (*DeletePolicyReq, error) {
	permission := &Permission{}
	permission, err := permission.FromDomain(req.Permission)
	if err != nil {
		return nil, err
	}
	subScope := &Resource{}
	subScope, err = subScope.FromDomain(req.SubjectScope)
	if err != nil {
		return nil, err
	}
	objScope := &Resource{}
	objScope, err = objScope.FromDomain(req.ObjectScope)
	if err != nil {
		return nil, err
	}
	return &DeletePolicyReq{
		Id:           req.Id,
		SubjectScope: subScope,
		ObjectScope:  objScope,
		Permission:   permission,
	}, nil
}

func (x *AdministrationReq) ToDomain() (*domain.AdministrationReq, error) {
	switch x.Kind {
	case AdministrationReq_CreateResource:
		protoReq := &CreateResourceReq{}
		err := proto.Unmarshal(x.ReqMarshalled, protoReq)
		if err != nil {
			return nil, err
		}
		req, err := protoReq.ToDomain()
		if err != nil {
			return nil, err
		}
		return &domain.AdministrationReq{
			ReqKind: domain.CreateResource,
			Request: req,
		}, nil
	case AdministrationReq_DeleteResource:
		protoReq := &DeleteResourceReq{}
		err := proto.Unmarshal(x.ReqMarshalled, protoReq)
		if err != nil {
			return nil, err
		}
		req, err := protoReq.ToDomain()
		if err != nil {
			return nil, err
		}
		return &domain.AdministrationReq{
			ReqKind: domain.DeleteResource,
			Request: req,
		}, nil
	case AdministrationReq_PutAttribute:
		protoReq := &PutAttributeReq{}
		err := proto.Unmarshal(x.ReqMarshalled, protoReq)
		if err != nil {
			return nil, err
		}
		req, err := protoReq.ToDomain()
		if err != nil {
			return nil, err
		}
		return &domain.AdministrationReq{
			ReqKind: domain.PutAttribute,
			Request: req,
		}, nil
	case AdministrationReq_DeleteAttribute:
		protoReq := &DeleteAttributeReq{}
		err := proto.Unmarshal(x.ReqMarshalled, protoReq)
		if err != nil {
			return nil, err
		}
		req, err := protoReq.ToDomain()
		if err != nil {
			return nil, err
		}
		return &domain.AdministrationReq{
			ReqKind: domain.DeleteAttribute,
			Request: req,
		}, nil
	case AdministrationReq_CreateInheritanceRel:
		protoReq := &CreateInheritanceRelReq{}
		err := proto.Unmarshal(x.ReqMarshalled, protoReq)
		if err != nil {
			return nil, err
		}
		req, err := protoReq.ToDomain()
		if err != nil {
			return nil, err
		}
		return &domain.AdministrationReq{
			ReqKind: domain.CreateInheritanceRel,
			Request: req,
		}, nil
	case AdministrationReq_DeleteInheritanceRel:
		protoReq := &DeleteInheritanceRelReq{}
		err := proto.Unmarshal(x.ReqMarshalled, protoReq)
		if err != nil {
			return nil, err
		}
		req, err := protoReq.ToDomain()
		if err != nil {
			return nil, err
		}
		return &domain.AdministrationReq{
			ReqKind: domain.DeleteInheritanceRel,
			Request: req,
		}, nil
	case AdministrationReq_CreatePolicy:
		protoReq := &CreatePolicyReq{}
		err := proto.Unmarshal(x.ReqMarshalled, protoReq)
		if err != nil {
			return nil, err
		}
		req, err := protoReq.ToDomain()
		if err != nil {
			return nil, err
		}
		return &domain.AdministrationReq{
			ReqKind: domain.CreatePolicy,
			Request: req,
		}, nil
	case AdministrationReq_DeletePolicy:
		protoReq := &DeletePolicyReq{}
		err := proto.Unmarshal(x.ReqMarshalled, protoReq)
		if err != nil {
			return nil, err
		}
		req, err := protoReq.ToDomain()
		if err != nil {
			return nil, err
		}
		return &domain.AdministrationReq{
			ReqKind: domain.DeletePolicy,
			Request: req,
		}, nil
	default:
		return nil, errors.New("unknown req kind")
	}
}

func (x *AdministrationReq) FromDomain(req domain.AdministrationReq) (*AdministrationReq, error) {
	switch req.ReqKind {
	case domain.CreateResource:
		protoReq := &CreateResourceReq{}
		protoReq, err := protoReq.FromDomain(*req.Request.(*domain.CreateResourceReq))
		if err != nil {
			return nil, err
		}
		reqMarshalled, err := proto.Marshal(protoReq)
		if err != nil {
			return nil, err
		}
		return &AdministrationReq{
			Kind:          AdministrationReq_CreateResource,
			ReqMarshalled: reqMarshalled,
		}, nil
	case domain.DeleteResource:
		protoReq := &DeleteResourceReq{}
		protoReq, err := protoReq.FromDomain(*req.Request.(*domain.DeleteResourceReq))
		if err != nil {
			return nil, err
		}
		reqMarshalled, err := proto.Marshal(protoReq)
		if err != nil {
			return nil, err
		}
		return &AdministrationReq{
			Kind:          AdministrationReq_DeleteResource,
			ReqMarshalled: reqMarshalled,
		}, nil
	case domain.PutAttribute:
		protoReq := &PutAttributeReq{}
		protoReq, err := protoReq.FromDomain(*req.Request.(*domain.PutAttributeReq))
		if err != nil {
			return nil, err
		}
		reqMarshalled, err := proto.Marshal(protoReq)
		if err != nil {
			return nil, err
		}
		return &AdministrationReq{
			Kind:          AdministrationReq_PutAttribute,
			ReqMarshalled: reqMarshalled,
		}, nil
	case domain.DeleteAttribute:
		protoReq := &DeleteAttributeReq{}
		protoReq, err := protoReq.FromDomain(*req.Request.(*domain.DeleteAttributeReq))
		if err != nil {
			return nil, err
		}
		reqMarshalled, err := proto.Marshal(protoReq)
		if err != nil {
			return nil, err
		}
		return &AdministrationReq{
			Kind:          AdministrationReq_DeleteAttribute,
			ReqMarshalled: reqMarshalled,
		}, nil
	case domain.CreateInheritanceRel:
		protoReq := &CreateInheritanceRelReq{}
		protoReq, err := protoReq.FromDomain(*req.Request.(*domain.CreateInheritanceRelReq))
		if err != nil {
			return nil, err
		}
		reqMarshalled, err := proto.Marshal(protoReq)
		if err != nil {
			return nil, err
		}
		return &AdministrationReq{
			Kind:          AdministrationReq_CreateInheritanceRel,
			ReqMarshalled: reqMarshalled,
		}, nil
	case domain.DeleteInheritanceRel:
		protoReq := &DeleteInheritanceRelReq{}
		protoReq, err := protoReq.FromDomain(*req.Request.(*domain.DeleteInheritanceRelReq))
		if err != nil {
			return nil, err
		}
		reqMarshalled, err := proto.Marshal(protoReq)
		if err != nil {
			return nil, err
		}
		return &AdministrationReq{
			Kind:          AdministrationReq_DeleteInheritanceRel,
			ReqMarshalled: reqMarshalled,
		}, nil
	case domain.CreatePolicy:
		protoReq := &CreatePolicyReq{}
		protoReq, err := protoReq.FromDomain(*req.Request.(*domain.CreatePolicyReq))
		if err != nil {
			return nil, err
		}
		reqMarshalled, err := proto.Marshal(protoReq)
		if err != nil {
			return nil, err
		}
		return &AdministrationReq{
			Kind:          AdministrationReq_CreatePolicy,
			ReqMarshalled: reqMarshalled,
		}, nil
	case domain.DeletePolicy:
		protoReq := &DeletePolicyReq{}
		protoReq, err := protoReq.FromDomain(*req.Request.(*domain.DeletePolicyReq))
		if err != nil {
			return nil, err
		}
		reqMarshalled, err := proto.Marshal(protoReq)
		if err != nil {
			return nil, err
		}
		return &AdministrationReq{
			Kind:          AdministrationReq_DeletePolicy,
			ReqMarshalled: reqMarshalled,
		}, nil
	default:
		return nil, errors.New("unknown req kind")
	}
}

//func NewSyncRespOutboxMessage(reqId string, error string, successful bool) domain.OutboxMessage {
//	resp := domainpb.AsyncSyncResp{
//		Id:         reqId,
//		Error:      error,
//		Successful: successful,
//	}
//	payload, err := proto.Marshal(&resp)
//	if err != nil {
//		log.Println(err)
//		return domain.OutboxMessage{}
//	}
//	return domain.OutboxMessage{
//		Kind:    domain.SyncRespOutboxMessageKind,
//		ReqMarshalled: payload,
//	}
//}
