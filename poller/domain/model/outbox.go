package model

type OutboxMessage struct {
	Id      string
	Kind    string
	Payload []byte
}
