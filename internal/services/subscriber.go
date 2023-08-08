package services

type Subscriber interface {
	Subscribe(subject, queueGroup string, handler func([]byte) error) error
	Unsubscribe() error
}
