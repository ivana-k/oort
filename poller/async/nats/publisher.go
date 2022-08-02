package nats

import (
	"errors"
	"github.com/c12s/oort/poller/domain/async"
	"github.com/nats-io/nats.go"
)

type Publisher struct {
	conn *nats.Conn
}

func NewPublisher(conn *nats.Conn) (async.Publisher, error) {
	if conn == nil {
		return nil, errors.New("conn nil")
	}
	return &Publisher{
		conn: conn,
	}, nil
}

func (p Publisher) Publish(subject string, message []byte) error {
	return p.conn.Publish(subject, message)
}
