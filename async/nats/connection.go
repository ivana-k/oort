package nats

import (
	"github.com/nats-io/nats.go"
)

func NewConnection(uri string) (*nats.Conn, error) {
	connection, err := nats.Connect(uri)
	if err != nil {
		return nil, err
	}
	return connection, nil
}
