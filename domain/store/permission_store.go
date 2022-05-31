package store

import (
	"github.com/c12s/oort/domain/model"
	storemodel "github.com/c12s/oort/store/neo4j/model"
)

type PermissionStore interface {
	AddResource(resource model.Resource) error
	AddIdentity(resource model.Resource) error
	AddResourceToPath(resource model.Resource, path storemodel.Path) error
	AddIdentityToPath(resource model.Resource, path storemodel.Path) error
	Connect(parentPath storemodel.Path, childPath storemodel.Path) error
	AddPermission(identityPath storemodel.Path, resourcePath storemodel.Path, permission model.Permission) error
	CheckPermission(identityPath storemodel.Path, resourcePath storemodel.Path, permission model.Permission) (bool, error)
}
