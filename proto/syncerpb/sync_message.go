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
	case SyncMessage_PutAttribute:
		req := &PutAttributeReq{}
		err = proto.Unmarshal(x.Payload, req)
		request = req
	case SyncMessage_DeleteAttribute:
		req := &DeleteAttributeReq{}
		err = proto.Unmarshal(x.Payload, req)
		request = req
	case SyncMessage_CreateInheritanceRel:
		req := &CreateInheritanceRelReq{}
		err = proto.Unmarshal(x.Payload, req)
		request = req
	case SyncMessage_DeleteInheritanceRel:
		req := &DeleteInheritanceRelReq{}
		err = proto.Unmarshal(x.Payload, req)
		request = req
	case SyncMessage_CreatePolicy:
		req := &CreatePolicyReq{}
		err = proto.Unmarshal(x.Payload, req)
		request = req
	case SyncMessage_DeletePolicy:
		req := &DeletePolicyReq{}
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
