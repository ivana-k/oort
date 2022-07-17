package store

import "github.com/c12s/oort/domain/model"

type AclStore interface {
	ConnectResources(parent, child model.Resource) error
	DisconnectResources(parent, child model.Resource) error
	UpsertAttribute(resource model.Resource, attribute model.Attribute) error
	RemoveAttribute(resource model.Resource, attribute model.Attribute) error
	InsertPermission(principal, resource model.Resource, permission model.Permission) error
	RemovePermission(principal, resource model.Resource, permission model.Permission) error
	CheckPermission(principal, resource model.Resource, permissionName string) error
}
