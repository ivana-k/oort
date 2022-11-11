package syncerpb

import (
	"github.com/c12s/oort/domain/async"
	"github.com/c12s/oort/domain/syncer"
	"google.golang.org/protobuf/proto"
)

type protoSyncMessageSerializer struct {
}

func NewProtoSyncMessageSerializer() async.SyncMessageSerializer {
	return protoSyncMessageSerializer{}
}

func (s protoSyncMessageSerializer) Serialize(msg async.SyncMessage) ([]byte, error) {
	return proto.Marshal(msg.(*SyncMessage))
}

func (s protoSyncMessageSerializer) Deserialize(bytes []byte) (async.SyncMessage, error) {
	req := &SyncMessage{}
	err := proto.Unmarshal(bytes, req)
	return req, err
}

func (x *SyncMessage) Request() (syncer.Request, error) {
	var request protoRequest
	var err error
	switch x.Kind {
	case SyncMessage_CONNECT_RESOURCES:
		req := &ConnectResourcesReq{}
		err = proto.Unmarshal(x.Payload, req)
		request = req
	case SyncMessage_DISCONNECT_RESOURCES:
		req := &DisconnectResourcesReq{}
		err = proto.Unmarshal(x.Payload, req)
		request = req
	case SyncMessage_UPSERT_ATTRIBUTE:
		req := &UpsertAttributeReq{}
		err = proto.Unmarshal(x.Payload, req)
		request = req
	case SyncMessage_REMOVE_ATTRIBUTE:
		req := &RemoveAttributeReq{}
		err = proto.Unmarshal(x.Payload, req)
		request = req
	case SyncMessage_INSERT_PERMISSION:
		req := &InsertPermissionReq{}
		err = proto.Unmarshal(x.Payload, req)
		request = req
	case SyncMessage_REMOVE_PERMISSION:
		req := &RemovePermissionReq{}
		err = proto.Unmarshal(x.Payload, req)
		request = req
	default:
		request = nil
	}
	if err != nil {
		return nil, err
	}
	return request.MapToDomain()
}

func (x *SyncMessage) RequestKind() async.SyncMsgKind {
	return async.SyncMsgKind(x.Kind)
}
