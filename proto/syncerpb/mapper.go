package syncerpb

import (
	"github.com/c12s/oort/domain/model"
	"github.com/c12s/oort/domain/syncer"
	"google.golang.org/protobuf/proto"
)

type protoRequest interface {
	MapToDomain() (syncer.Request, error)
}

func (x *ConnectResourcesReq) MapToDomain() (syncer.Request, error) {
	return syncer.ConnectResourcesReq{
		ReqId:  x.Id,
		Parent: x.Parent.MapToDomain(),
		Child:  x.Child.MapToDomain(),
	}, nil
}

func (x *DisconnectResourcesReq) MapToDomain() (syncer.Request, error) {
	return syncer.DisconnectResourcesReq{
		ReqId:  x.Id,
		Parent: x.Parent.MapToDomain(),
		Child:  x.Child.MapToDomain(),
	}, nil
}

func (x *UpsertAttributeReq) MapToDomain() (syncer.Request, error) {
	attr, err := x.Attribute.MapToDomain()
	if err != nil {
		return syncer.UpsertAttributeReq{}, err
	}
	return syncer.UpsertAttributeReq{
		ReqId:     x.Id,
		Resource:  x.Resource.MapToDomain(),
		Attribute: attr,
	}, nil
}

func (x *RemoveAttributeReq) MapToDomain() (syncer.Request, error) {
	return syncer.RemoveAttributeReq{
		ReqId:       x.Id,
		Resource:    x.Resource.MapToDomain(),
		AttributeId: x.AttributeId.MapToDomain(),
	}, nil
}

func (x *InsertPermissionReq) MapToDomain() (syncer.Request, error) {
	permission, err := x.Permission.MapToDomain()
	if err != nil {
		return nil, err
	}
	return syncer.InsertPermissionReq{
		ReqId:      x.Id,
		Principal:  x.Principal.MapToDomain(),
		Resource:   x.Resource.MapToDomain(),
		Permission: permission,
	}, nil
}

func (x *RemovePermissionReq) MapToDomain() (syncer.Request, error) {
	permission, err := x.Permission.MapToDomain()
	if err != nil {
		return nil, err
	}
	return syncer.RemovePermissionReq{
		ReqId:      x.Id,
		Principal:  x.Principal.MapToDomain(),
		Resource:   x.Resource.MapToDomain(),
		Permission: permission,
	}, nil
}

func NewSyncRespOutboxMessage(reqId string, error string, successful bool) *model.OutboxMessage {
	resp := AsyncSyncResp{
		ReqId:      reqId,
		Error:      error,
		Successful: successful,
	}
	payload, err := proto.Marshal(&resp)
	if err != nil {
		return nil
	}
	return &model.OutboxMessage{
		Kind:    model.SyncRespOutboxMessageKind,
		Payload: payload,
	}
}
