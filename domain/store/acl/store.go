package acl

type Store interface {
	ConnectResources(req ConnectResourcesReq) SyncResp
	DisconnectResources(req DisconnectResourcesReq) SyncResp
	UpsertAttribute(req UpsertAttributeReq) SyncResp
	RemoveAttribute(req RemoveAttributeReq) SyncResp
	GetAttributes(req GetAttributeReq) GetAttributeResp
	InsertPermission(req InsertPermissionReq) SyncResp
	RemovePermission(req RemovePermissionReq) SyncResp
	GetPermissionByPrecedence(req GetPermissionReq) GetPermissionByPrecedenceResp
}
