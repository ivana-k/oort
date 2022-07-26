package async

import "github.com/nats-io/nats.go"

type Subscriber interface {
	Subscribe(handler func(*nats.Msg)) error
	Unsubscribe() error
}
