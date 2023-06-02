package neo4j

import (
	"github.com/c12s/oort/poller/domain/model"
	"github.com/c12s/oort/poller/domain/store/outbox"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type OutboxStore struct {
	manager *TransactionManager
}

func NewOutboxStore(manager *TransactionManager) outbox.Store {
	return OutboxStore{
		manager: manager,
	}
}

func (store OutboxStore) GetUnprocessed() ([]model.OutboxMessage, error) {
	results, err := store.manager.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(getUnprocessedCypher())
		if err != nil {
			return nil, err
		}
		if result.Err() != nil {
			return nil, result.Err()
		}

		records := make([]interface{}, 0)
		for result.Next() {
			records = append(records, result.Record().Values)
		}
		return records, nil
	})

	if err != nil {
		return []model.OutboxMessage{}, err
	}
	return getOutboxMessages(results), nil
}

func (store OutboxStore) DeleteById(message model.OutboxMessage) error {
	_, err := store.manager.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(deleteByIdCypher(message))
		if err != nil {
			return nil, err
		}
		return nil, result.Err()
	})
	return err
}
