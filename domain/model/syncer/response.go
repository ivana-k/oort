package syncer

type SyncResp struct {
	Error error
}

type ConnectResourcesResp struct {
	Resp SyncResp
}

type DisconnectResourcesResp struct {
	Resp SyncResp
}

type UpsertAttributeResp struct {
	Resp SyncResp
}

type RemoveAttributeResp struct {
	Resp SyncResp
}

type InsertPermissionResp struct {
	Resp SyncResp
}

type RemovePermissionResp struct {
	Resp SyncResp
}
