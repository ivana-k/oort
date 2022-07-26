package jetstream

import (
	"errors"
	"github.com/c12s/oort/domain/async"
	"github.com/nats-io/nats.go"
)

type Subscriber struct {
	jsContext    nats.JetStreamContext
	subject      string
	subscription *nats.Subscription
}

func NewSubscriber(conn *nats.Conn, subject string) (async.Subscriber, error) {
	js, err := conn.JetStream()
	if err != nil {
		return nil, err
	}
	return &Subscriber{
		jsContext: js,
		subject:   subject,
	}, nil
}

func (s *Subscriber) Subscribe(handler func(msg *nats.Msg)) error {
	if s.subscription != nil {
		return errors.New("already subscribed")
	}
	subscription, err := s.jsContext.Subscribe(s.subject, handler)
	if err != nil {
		return err
	}
	s.subscription = subscription
	return nil
}

func (s *Subscriber) Unsubscribe() error {
	if s.subscription == nil {
		return nil
	}
	err := s.subscription.Unsubscribe()
	return err
}
