package model

type OutboxMessage struct {
	Id      int64
	Kind    string
	Payload []byte
}
