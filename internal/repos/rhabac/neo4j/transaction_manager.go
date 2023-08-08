package neo4j

import (
	"github.com/c12s/oort/internal/domain"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"log"
)

type TransactionManager struct {
	driver neo4j.Driver
	dbName string
}

func NewTransactionManager(uri, dbName string) (*TransactionManager, error) {
	driver, err := neo4j.NewDriver(uri, neo4j.NoAuth())
	if err != nil {
		return nil, err
	}
	return &TransactionManager{
		driver: driver,
		dbName: dbName,
	}, nil
}

type TransactionFunction func(transaction neo4j.Transaction) (interface{}, error)

func (manager *TransactionManager) WriteTransaction(cypher string, params map[string]interface{}, generator domain.OutboxMsgGenerator) error {
	_, err := manager.writeTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(cypher, params)
		if err != nil {
			_ = transaction.Rollback()
		}
		// todo: otkomentarisi ovo kasnije
		//if callback == nil {
		//	return nil, nil
		//}
		//outboxMessage := callback(err)
		//_, err = transaction.Run(manager.getOutboxMessageCypher(outboxMessage))
		//if err != nil {
		//	_ = transaction.Rollback()
		//	return nil, errors.New("outbox message could not be stored - " + err.Error())
		//}
		if result == nil {
			return nil, nil
		}
		return nil, result.Err()
	})
	return err
}

func (manager *TransactionManager) WriteTransactions(cyphers []string, params []map[string]interface{}, generator domain.OutboxMsgGenerator) error {
	_, err := manager.writeTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		var txErr error = nil
		for i := range cyphers {
			cypher := cyphers[i]
			param := params[i]
			result, err := transaction.Run(cypher, param)
			if err != nil || result.Err() != nil {
				_ = transaction.Rollback()
				if err != nil {
					txErr = err
				} else {
					txErr = result.Err()
				}
				break
			}
		}
		//if callback == nil {
		//	return nil, nil
		//}
		//outboxMessage := callback(txErr)
		//result, err := transaction.Run(manager.getOutboxMessageCypher(outboxMessage))
		//if err != nil || result.Err() != nil {
		//	_ = transaction.Rollback()
		//	return nil, errors.New("outbox message could not be stored - " + err.Error())
		//}
		return txErr, nil
	})
	return err
}

func (manager *TransactionManager) ReadTransaction(cypher string, params map[string]interface{}) (interface{}, error) {
	return manager.readTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(cypher, params)
		if err != nil {
			return nil, err
		}
		if result.Err() != nil {
			return nil, result.Err()
		}
		return result.Collect()
	})
}

func (manager *TransactionManager) writeTransaction(txFunc TransactionFunction) (interface{}, error) {
	session := manager.driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: manager.dbName})
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {
			log.Println(err)
		}
	}(session)

	result, err := session.WriteTransaction(neo4j.TransactionWork(txFunc))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (manager *TransactionManager) readTransaction(txFunc TransactionFunction) (interface{}, error) {
	session := manager.driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
		DatabaseName: manager.dbName})
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {
			log.Println(err)
		}
	}(session)

	result, err := session.ReadTransaction(neo4j.TransactionWork(txFunc))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (manager *TransactionManager) Stop() {
	err := manager.driver.Close()
	if err != nil {
		log.Println("error while closing neo4j conn: ", err)
	}
}

//func (manager *TransactionManager) getOutboxMessageCypher(message model.OutboxMessage) (string, map[string]interface{}) {
//	return "CREATE (:OutboxMessage{kind: $kind, payload: $payload, processing: $processing})",
//		map[string]interface{}{"kind": message.Kind,
//			"payload":    message.Payload,
//			"processing": false}
//}
