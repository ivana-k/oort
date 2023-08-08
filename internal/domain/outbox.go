package domain

const (
	SyncRespOutboxMessageKind = "sync.response"
)

type OutboxMessage struct {
	Kind    string
	Payload []byte
}
