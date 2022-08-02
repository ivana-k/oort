package outbox

import "github.com/c12s/oort/poller/domain/model"

type Store interface {
	ReserveAndGetUnprocessed() ([]model.OutboxMessage, error)
	SetUnprocessed(message model.OutboxMessage) error
	DeleteById(message model.OutboxMessage) error
}
