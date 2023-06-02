package async

import (
	"errors"
	"github.com/c12s/oort/domain/async"
	"github.com/c12s/oort/domain/syncer"
)

type SyncerAsyncApi struct {
	serializer async.SyncMessageSerializer
	handler    syncer.Handler
}

func NewSyncerAsyncApi(subscriber async.Subscriber, subject, queueGroup string, serializer async.SyncMessageSerializer, handler syncer.Handler) error {
	s := SyncerAsyncApi{serializer: serializer, handler: handler}
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
	switch msg.RequestKind() {
	case async.CreateResource:
		s.handler.CreateResource(request.(syncer.CreateResourceReq))
	case async.DeleteResource:
		s.handler.DeleteResource(request.(syncer.DeleteResourceReq))
	case async.CreateAttribute:
		s.handler.CreateAttribute(request.(syncer.CreateAttributeReq))
	case async.UpdateAttribute:
		s.handler.UpdateAttribute(request.(syncer.UpdateAttributeReq))
	case async.DeleteAttribute:
		s.handler.DeleteAttribute(request.(syncer.DeleteAttributeReq))
	case async.CreateAggregationRel:
		s.handler.CreateAggregationRelReq(request.(syncer.CreateAggregationRelReq))
	case async.DeleteAggregationRel:
		s.handler.DeleteAggregationRelReq(request.(syncer.DeleteAggregationRelReq))
	case async.CreateCompositionRel:
		s.handler.CreateCompositionRelReq(request.(syncer.CreateCompositionRelReq))
	case async.DeleteCompositionRel:
		s.handler.DeleteCompositionRelReq(request.(syncer.DeleteCompositionRelReq))
	case async.CreatePermission:
		s.handler.CreatePermission(request.(syncer.CreatePermissionReq))
	case async.DeletePermission:
		s.handler.DeletePermission(request.(syncer.DeletePermissionReq))
	default:
		respError = errors.New("unknown message kind")
	}
	return respError
}
