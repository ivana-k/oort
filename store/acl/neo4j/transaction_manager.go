package neo4j

import (
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

func (manager *TransactionManager) WriteTransaction(txFunc TransactionFunction) (interface{}, error) {
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

func (manager *TransactionManager) ReadTransaction(txFunc TransactionFunction) (interface{}, error) {
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
