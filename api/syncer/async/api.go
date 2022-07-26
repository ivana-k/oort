package async

import (
	"github.com/c12s/oort/domain/async"
	"github.com/nats-io/nats.go"
)

type SyncerAsyncApi struct {
}

func NewSyncerAsyncApi(subscriber async.Subscriber, subject, queueGroup string) error {
	s := SyncerAsyncApi{}
	return subscriber.Subscribe(subject, queueGroup, s.Handle)
}

func (s SyncerAsyncApi) Handle(message *nats.Msg) {

}
