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
	case async.PutAttribute:
		s.handler.PutAttribute(request.(syncer.PutAttributeReq))
	case async.DeleteAttribute:
		s.handler.DeleteAttribute(request.(syncer.DeleteAttributeReq))
	case async.CreateInheritanceRel:
		s.handler.CreateInheritanceRel(request.(syncer.CreateInheritanceRelReq))
	case async.DeleteInheritanceRel:
		s.handler.DeleteInheritanceRel(request.(syncer.DeleteInheritanceRelReq))
	case async.CreatePolicy:
		s.handler.CreatePolicy(request.(syncer.CreatePolicyReq))
	case async.DeletePolicy:
		s.handler.DeletePolicy(request.(syncer.DeletePolicyReq))
	default:
		respError = errors.New("unknown message kind")
	}
	return respError
}
