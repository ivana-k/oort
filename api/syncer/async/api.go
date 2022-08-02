package async

import (
	"errors"
	"github.com/c12s/oort/domain/async"
	"github.com/c12s/oort/domain/syncer"
	"github.com/c12s/oort/proto/common"
	"github.com/c12s/oort/proto/syncerpb"
	"google.golang.org/protobuf/proto"
)

type SyncerAsyncApi struct {
	serializer async.SyncMessageSerializer
	handler    syncer.Handler
}

func NewSyncerAsyncApi(subscriber async.Subscriber, subject, queueGroup string, serializer async.SyncMessageSerializer, handler syncer.Handler) error {
	s := SyncerAsyncApi{serializer: serializer, handler: handler}
	req := syncerpb.ConnectResourcesReq{
		Id: "reqid",
		Parent: &common.Resource{
			Id:   "id",
			Kind: "kind",
		},
		Child: &common.Resource{
			Id:   "cid",
			Kind: "ckind",
		},
	}
	b, _ := proto.Marshal(&req)
	msg := syncerpb.SyncReq{
		Kind:    syncerpb.SyncReq_CONNECT_RESOURCES,
		Payload: b,
	}
	msgProto, _ := proto.Marshal(&msg)
	err := s.handle(msgProto)
	if err != nil {
		panic(err)
	}
	return subscriber.Subscribe(subject, queueGroup, s.handle)
}

func (s SyncerAsyncApi) handle(message []byte) error {
	msg, err := s.serializer.Deserialize(message)
	if err != nil {
		return err
	}
	request, err := msg.Request()
	if err != nil {
		return err
	}
	var respError error
	switch msg.MessageKind() {
	case async.ConnectResources:
		s.handler.ConnectResources(request.(syncer.ConnectResourcesReq))
	case async.DisconnectResources:
		s.handler.DisconnectResources(request.(syncer.DisconnectResourcesReq))
	case async.UpsertAttribute:
		s.handler.UpsertAttribute(request.(syncer.UpsertAttributeReq))
	case async.RemoveAttribute:
		s.handler.RemoveAttribute(request.(syncer.RemoveAttributeReq))
	case async.InsertPermission:
		s.handler.InsertPermission(request.(syncer.InsertPermissionReq))
	case async.RemovePermission:
		s.handler.RemovePermission(request.(syncer.RemovePermissionReq))
	default:
		respError = errors.New("unknown message kind")
	}
	return respError
}
