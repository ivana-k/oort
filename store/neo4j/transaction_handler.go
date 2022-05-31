package neo4j

import (
	"fmt"
	"github.com/c12s/oort/config"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type TransactionHandler struct {
	config config.Neo4j
	driver neo4j.Driver
}

func NewTransactionHandler(config config.Neo4j) *TransactionHandler {
	return &TransactionHandler{
		config: config,
	}
}

func (h *TransactionHandler) Write(cypher string, params map[string]interface{}) (neo4j.Result, error) {
	return h.transaction(neo4j.AccessModeWrite, cypher, params)
}

func (h *TransactionHandler) Read(cypher string, params map[string]interface{}) (neo4j.Result, error) {
	return h.transaction(neo4j.AccessModeRead, cypher, params)
}

func (h *TransactionHandler) transaction(mode neo4j.AccessMode, cypher string, params map[string]interface{}) (neo4j.Result, error) {
	session, err := h.getSession(mode)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	result, err := session.WriteTransaction(h.handleTransaction(cypher, params))
	if err != nil {
		return nil, err
	}
	return result.(neo4j.Result), err
}

func (h *TransactionHandler) handleTransaction(cypher string, params map[string]interface{}) func(transaction neo4j.Transaction) (interface{}, error) {
	return func(transaction neo4j.Transaction) (interface{}, error) {
		fmt.Println(cypher)
		result, err := transaction.Run(cypher, params)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
}

func (h *TransactionHandler) getSession(mode neo4j.AccessMode) (neo4j.Session, error) {
	driver, err := h.getDriver()
	if err != nil {
		return nil, err
	}
	return driver.NewSession(neo4j.SessionConfig{AccessMode: mode, DatabaseName: h.config.DbName()}), nil
}

func (h *TransactionHandler) getDriver() (neo4j.Driver, error) {
	if h.driver != nil {
		return h.driver, nil
	}
	err := h.setDriver()
	if err != nil {
		return nil, err
	}
	return h.driver, nil
}

func (h *TransactionHandler) setDriver() error {
	driver, err := neo4j.NewDriver(h.config.Uri(), neo4j.BasicAuth(h.config.Username(), h.config.Password(), ""))
	h.driver = driver
	return err
}
