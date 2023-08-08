package startup

import "github.com/nats-io/nats.go"

func newNatsConn(uri string) (*nats.Conn, error) {
	connection, err := nats.Connect(uri)
	if err != nil {
		return nil, err
	}
	return connection, nil
}
