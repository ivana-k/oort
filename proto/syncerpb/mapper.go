package syncerpb

import (
	"github.com/c12s/oort/domain/model"
	"github.com/c12s/oort/domain/syncer"
	"google.golang.org/protobuf/proto"
	"log"
)

type protoRequest interface {
	MapToDomain() (syncer.Request, error)
}

func (x *CreateResourceReq) MapToDomain() (syncer.Request, error) {
	return syncer.CreateResourceReq{
		ReqId:    x.Id,
		Resource: x.Resource.MapToDomain(),
	}, nil
}

func (x *DeleteResourceReq) MapToDomain() (syncer.Request, error) {
	return syncer.DeleteResourceReq{
		ReqId:    x.Id,
		Resource: x.Resource.MapToDomain(),
	}, nil
}

func (x *CreateAttributeReq) MapToDomain() (syncer.Request, error) {
	attr, err := x.Attribute.MapToDomain()
	if err != nil {
		return syncer.CreateAttributeReq{}, err
	}
	return syncer.CreateAttributeReq{
		ReqId:     x.Id,
		Resource:  x.Resource.MapToDomain(),
		Attribute: attr,
	}, nil
}

func (x *UpdateAttributeReq) MapToDomain() (syncer.Request, error) {
	attr, err := x.Attribute.MapToDomain()
	if err != nil {
		return syncer.UpdateAttributeReq{}, err
	}
	return syncer.UpdateAttributeReq{
		ReqId:     x.Id,
		Resource:  x.Resource.MapToDomain(),
		Attribute: attr,
	}, nil
}

func (x *DeleteAttributeReq) MapToDomain() (syncer.Request, error) {
	return syncer.DeleteAttributeReq{
		ReqId:       x.Id,
		Resource:    x.Resource.MapToDomain(),
		AttributeId: x.AttributeId.MapToDomain(),
	}, nil
}

func (x *CreateAggregationRelReq) MapToDomain() (syncer.Request, error) {
	return syncer.CreateAggregationRelReq{
		ReqId:  x.Id,
		Parent: x.Parent.MapToDomain(),
		Child:  x.Child.MapToDomain(),
	}, nil
}

func (x *DeleteAggregationRelReq) MapToDomain() (syncer.Request, error) {
	return syncer.DeleteAggregationRelReq{
		ReqId:  x.Id,
		Parent: x.Parent.MapToDomain(),
		Child:  x.Child.MapToDomain(),
	}, nil
}

func (x *CreateCompositionRelReq) MapToDomain() (syncer.Request, error) {
	return syncer.CreateCompositionRelReq{
		ReqId:  x.Id,
		Parent: x.Parent.MapToDomain(),
		Child:  x.Child.MapToDomain(),
	}, nil
}

func (x *DeleteCompositionRelReq) MapToDomain() (syncer.Request, error) {
	return syncer.DeleteCompositionRelReq{
		ReqId:  x.Id,
		Parent: x.Parent.MapToDomain(),
		Child:  x.Child.MapToDomain(),
	}, nil
}

func (x *CreatePermissionReq) MapToDomain() (syncer.Request, error) {
	permission, err := x.Permission.MapToDomain()
	if err != nil {
		return nil, err
	}
	subject := x.Subject.MapToDomain()
	object := x.Object.MapToDomain()
	return syncer.CreatePermissionReq{
		ReqId:      x.Id,
		Subject:    &subject,
		Object:     &object,
		Permission: permission,
	}, nil
}

func (x *DeletePermissionReq) MapToDomain() (syncer.Request, error) {
	permission, err := x.Permission.MapToDomain()
	if err != nil {
		return nil, err
	}
	subject := x.Subject.MapToDomain()
	object := x.Object.MapToDomain()
	return syncer.DeletePermissionReq{
		ReqId:      x.Id,
		Subject:    &subject,
		Object:     &object,
		Permission: permission,
	}, nil
}

func NewSyncRespOutboxMessage(reqId string, error string, successful bool) model.OutboxMessage {
	resp := AsyncSyncResp{
		ReqId:      reqId,
		Error:      error,
		Successful: successful,
	}
	payload, err := proto.Marshal(&resp)
	if err != nil {
		log.Println(err)
		return model.OutboxMessage{}
	}
	return model.OutboxMessage{
		Kind:    model.SyncRespOutboxMessageKind,
		Payload: payload,
	}
}
