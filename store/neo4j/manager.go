package neo4j

import (
	"github.com/c12s/oort/config"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Manager struct {
	config config.Neo4j
	driver neo4j.Driver
}

func NewManager(config config.Neo4j) *Manager {
	return &Manager{
		config: config,
	}
}

func (n *Manager) GetDriver() (neo4j.Driver, error) {
	if n.driver != nil {
		return n.driver, nil
	}
	err := n.setDriver()
	if err != nil {
		return nil, err
	}
	return n.driver, nil
}

func (n *Manager) setDriver() error {
	driver, err := neo4j.NewDriver(n.config.Uri(), neo4j.BasicAuth(n.config.Username(), n.config.Password(), ""))
	n.driver = driver
	return err
}

func (n *Manager) GetSession(mode neo4j.AccessMode) (neo4j.Session, error) {
	driver, err := n.GetDriver()
	if err != nil {
		return nil, err
	}
	return driver.NewSession(neo4j.SessionConfig{AccessMode: mode, DatabaseName: n.config.DbName()}), nil
}
