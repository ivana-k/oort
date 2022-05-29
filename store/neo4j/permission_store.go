package neo4j

import (
	"github.com/c12s/oort/domain/store"
)

const (
	resourceVar        = "r"
	resourceLabel      = "Resource"
	identityLabel      = "Identity"
	parentRelationship = "PARENT"
)

//TODO: error handling

type permissionStore struct {
	manager *Manager
}

func NewNeo4jPermissionStore(manager *Manager) store.PermissionStore {
	return &permissionStore{
		manager: manager,
	}
}
