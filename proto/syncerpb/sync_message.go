package syncerpb

import (
	"errors"
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
	case SyncMessage_CreateResource:
		req := &CreateResourceReq{}
		err = proto.Unmarshal(x.Payload, req)
		request = req
	case SyncMessage_DeleteResource:
		req := &DeleteResourceReq{}
		err = proto.Unmarshal(x.Payload, req)
		request = req
	case SyncMessage_CreateAttribute:
		req := &CreateAttributeReq{}
		err = proto.Unmarshal(x.Payload, req)
		request = req
	case SyncMessage_UpdateAttribute:
		req := &UpdateAttributeReq{}
		err = proto.Unmarshal(x.Payload, req)
		request = req
	case SyncMessage_DeleteAttribute:
		req := &DeleteAttributeReq{}
		err = proto.Unmarshal(x.Payload, req)
		request = req
	case SyncMessage_CreateAggregationRel:
		req := &CreateAggregationRelReq{}
		err = proto.Unmarshal(x.Payload, req)
		request = req
	case SyncMessage_DeleteAggregationRel:
		req := &DeleteAggregationRelReq{}
		err = proto.Unmarshal(x.Payload, req)
		request = req
	case SyncMessage_CreateCompositionRel:
		req := &CreateCompositionRelReq{}
		err = proto.Unmarshal(x.Payload, req)
		request = req
	case SyncMessage_DeleteCompositionRel:
		req := &DeleteCompositionRelReq{}
		err = proto.Unmarshal(x.Payload, req)
		request = req
	case SyncMessage_CreatePermission:
		req := &CreatePermissionReq{}
		err = proto.Unmarshal(x.Payload, req)
		request = req
	case SyncMessage_DeletePermission:
		req := &DeletePermissionReq{}
		err = proto.Unmarshal(x.Payload, req)
		request = req
	default:
		request = nil
	}
	if err != nil {
		return nil, err
	}
	if request == nil {
		return nil, errors.New("unknown msg kind")
	}
	return request.MapToDomain()
}

func (x *SyncMessage) RequestKind() async.SyncMsgKind {
	return async.SyncMsgKind(x.Kind)
}
