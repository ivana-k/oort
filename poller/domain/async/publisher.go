package async

type Publisher interface {
	Publish(subject string, message []byte) error
}
