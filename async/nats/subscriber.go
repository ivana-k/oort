package nats

import (
	"errors"
	"github.com/c12s/oort/domain/async"
	"github.com/nats-io/nats.go"
)

const ackMaxRetry int8 = 5

type Subscriber struct {
	jsContext    nats.JetStreamContext
	subscription *nats.Subscription
}

func NewSubscriber(conn *nats.Conn) (async.Subscriber, error) {
	js, err := conn.JetStream()
	if err != nil {
		return nil, err
	}
	_, err = js.AddStream(&nats.StreamConfig{
		Name:     "SYNC",
		Subjects: []string{"sync"},
	})
	if err != nil {
		return nil, err
	}
	return &Subscriber{
		jsContext: js,
	}, nil
}

func (s *Subscriber) Subscribe(subject, queueGroup string, handler func(msg []byte) error) error {
	if s.subscription != nil {
		return errors.New("already subscribed")
	}
	subscription, err := s.jsContext.QueueSubscribe(subject, queueGroup, func(msg *nats.Msg) {
		err := handler(msg.Data)
		if err != nil {
			s.ackRetry(msg.Nak, ackMaxRetry)
		}
		s.ackRetry(msg.Ack, ackMaxRetry)
	})
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

func (s *Subscriber) ackRetry(ackFunc func(opts ...nats.AckOpt) error, maxRetry int8) {
	for i := maxRetry; i > 0; i-- {
		err := ackFunc(nil)
		if err == nil {
			break
		}
	}
}
