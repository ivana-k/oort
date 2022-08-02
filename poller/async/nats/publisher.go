package nats

import (
	"github.com/c12s/oort/poller/domain/async"
	"github.com/nats-io/nats.go"
)

type Publisher struct {
	jsContext nats.JetStreamContext
}

func NewPublisher(conn *nats.Conn) (async.Publisher, error) {
	js, err := conn.JetStream()
	if err != nil {
		return nil, err
	}
	_, err = js.AddStream(&nats.StreamConfig{
		Name:     "nzm",
		Subjects: []string{"sync.response"},
	})
	if err != nil {
		return nil, err
	}
	return &Publisher{
		jsContext: js,
	}, nil
}

func (p Publisher) Publish(subject string, message []byte) error {
	_, err := p.jsContext.Publish(subject, message)
	return err
}
