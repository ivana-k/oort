package syncerpb

import (
	"github.com/c12s/oort/domain/async"
	"github.com/c12s/oort/domain/model/syncer"
	"google.golang.org/protobuf/proto"
)

func (x *ConnectResourcesReq) MapToDomain() (syncer.Request, error) {
	return syncer.ConnectResourcesReq{
		Parent: x.Parent.MapToDomain(),
		Child:  x.Child.MapToDomain(),
	}, nil
}

func (x *DisconnectResourcesReq) MapToDomain() (syncer.Request, error) {
	return syncer.DisconnectResourcesReq{
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
		Resource:  x.Resource.MapToDomain(),
		Attribute: attr,
	}, nil
}

func (x *RemoveAttributeReq) MapToDomain() (syncer.Request, error) {
	return syncer.RemoveAttributeReq{
		Resource:    x.Resource.MapToDomain(),
		AttributeId: x.AttributeId.MapToDomain(),
	}, nil
}

func (x *InsertPermissionReq) MapToDomain() (syncer.Request, error) {
	return syncer.InsertPermissionReq{
		Principal:  x.Principal.MapToDomain(),
		Resource:   x.Resource.MapToDomain(),
		Permission: x.Permission.MapToDomain(),
	}, nil
}

func (x *RemovePermissionReq) MapToDomain() (syncer.Request, error) {
	return syncer.RemovePermissionReq{
		Principal:  x.Principal.MapToDomain(),
		Resource:   x.Resource.MapToDomain(),
		Permission: x.Permission.MapToDomain(),
	}, nil
}

func (x *SyncReq) Request() (syncer.Request, error) {
	var request ProtoRequest
	var err error
	switch x.Kind {
	case SyncReq_CONNECT_RESOURCES:
		req := &ConnectResourcesReq{}
		err = proto.Unmarshal(x.Payload, req)
		request = req
	case SyncReq_DISCONNECT_RESOURCES:
		req := &DisconnectResourcesReq{}
		err = proto.Unmarshal(x.Payload, req)
		request = req
	case SyncReq_UPSERT_ATTRIBUTE:
		req := &UpsertAttributeReq{}
		err = proto.Unmarshal(x.Payload, req)
		request = req
	case SyncReq_REMOVE_ATTRIBUTE:
		req := &RemoveAttributeReq{}
		err = proto.Unmarshal(x.Payload, req)
		request = req
	case SyncReq_INSERT_PERMISSION:
		req := &InsertPermissionReq{}
		err = proto.Unmarshal(x.Payload, req)
		request = req
	case SyncReq_REMOVE_PERMISSION:
		req := &RemovePermissionReq{}
		err = proto.Unmarshal(x.Payload, req)
		request = req
	default:
		request = nil
		request = nil
	}
	if err != nil {
		return nil, err
	}
	return request.MapToDomain()
}

func (x *SyncReq) MsgId() string {
	return x.Id
}

func (x *SyncReq) MsgKind() async.SyncMsgKind {
	return async.SyncMsgKind(x.Kind)
}
