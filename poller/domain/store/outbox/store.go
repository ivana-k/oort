package outbox

import "github.com/c12s/oort/poller/domain/model"

type Store interface {
	GetUnprocessed() ([]model.OutboxMessage, error)
	DeleteById(message model.OutboxMessage) error
}
