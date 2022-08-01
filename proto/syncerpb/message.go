package syncerpb

import (
	"github.com/c12s/oort/domain/async"
	"github.com/c12s/oort/domain/model/syncer"
	"google.golang.org/protobuf/proto"
)

type ProtoRequest interface {
	MapToDomain() (syncer.Request, error)
}

type syncMessageProtoSerializer struct {
}

func NewSyncMessageProtoSerializer() async.SyncMessageSerializer {
	return syncMessageProtoSerializer{}
}

func (s syncMessageProtoSerializer) Serialize(msg async.SyncMessage) ([]byte, error) {
	return proto.Marshal(msg.(*SyncReq))
}

func (s syncMessageProtoSerializer) Deserialize(bytes []byte) (async.SyncMessage, error) {
	req := &SyncReq{}
	err := proto.Unmarshal(bytes, req)
	return req, err
}
