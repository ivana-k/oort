package async

import (
	"errors"
	"github.com/c12s/oort/domain/async"
	"github.com/c12s/oort/domain/handler"
	"github.com/c12s/oort/domain/model/syncer"
)

type SyncerAsyncApi struct {
	serializer async.SyncMessageSerializer
	handler    handler.SyncerHandler
}

func NewSyncerAsyncApi(subscriber async.Subscriber, subject, queueGroup string, serializer async.SyncMessageSerializer, handler handler.SyncerHandler) error {
	s := SyncerAsyncApi{serializer: serializer, handler: handler}
	//req := syncerpb.ConnectResourcesReq{
	//	Parent: &common.Resource{
	//		Id:   "id",
	//		Kind: "kind",
	//	},
	//	Child: &common.Resource{
	//		Id:   "cid",
	//		Kind: "ckind",
	//	},
	//}
	//b, _ := proto.Marshal(&req)
	//msg := syncerpb.SyncReq{
	//	Id:      "id",
	//	Kind:    syncerpb.SyncReq_CONNECT_RESOURCES,
	//	Payload: b,
	//}
	//msgProto, _ := proto.Marshal(&msg)
	//err := s.handle(msgProto)
	//if err != nil {
	//	panic(err)
	//}
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
	switch msg.MsgKind() {
	case async.ConnectResources:
		respError = s.handler.ConnectResources(request.(syncer.ConnectResourcesReq)).GetError()
	case async.DisconnectResources:
		respError = s.handler.DisconnectResources(request.(syncer.DisconnectResourcesReq)).GetError()
	case async.UpsertAttribute:
		respError = s.handler.UpsertAttribute(request.(syncer.UpsertAttributeReq)).GetError()
	case async.RemoveAttribute:
		respError = s.handler.RemoveAttribute(request.(syncer.RemoveAttributeReq)).GetError()
	case async.InsertPermission:
		respError = s.handler.InsertPermission(request.(syncer.InsertPermissionReq)).GetError()
	case async.RemovePermission:
		respError = s.handler.RemovePermission(request.(syncer.RemovePermissionReq)).GetError()
	default:
		respError = errors.New("unknown message kind")
	}
	return respError
}
