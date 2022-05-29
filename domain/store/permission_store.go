package store

import (
	"github.com/c12s/oort/domain/model"
)

type PermissionStore interface {
	AddResource(resource model.Resource) error
	AddIdentity(resource model.Resource) error
	AddResourceToPath(resource model.Resource, path []model.Resource) error
	AddIdentityToPath(resource model.Resource, path []model.Resource) error
}
