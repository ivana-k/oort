package store

import (
	checker2 "github.com/c12s/oort/domain/model/checker"
	syncer2 "github.com/c12s/oort/domain/model/syncer"
)

type AclStore interface {
	ConnectResources(req syncer2.ConnectResourcesReq) syncer2.ConnectResourcesResp
	DisconnectResources(req syncer2.DisconnectResourcesReq) syncer2.DisconnectResourcesResp
	UpsertAttribute(req syncer2.UpsertAttributeReq) syncer2.UpsertAttributeResp
	RemoveAttribute(req syncer2.RemoveAttributeReq) syncer2.RemoveAttributeResp
	GetAttributes(req checker2.GetAttributeReq) checker2.GetAttributeResp
	InsertPermission(req syncer2.InsertPermissionReq) syncer2.InsertPermissionResp
	RemovePermission(req syncer2.RemovePermissionReq) syncer2.RemovePermissionResp
	GetPermissionByPrecedence(req checker2.GetPermissionReq) checker2.GetPermissionByPrecedenceResp
}
