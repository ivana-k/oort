package syncer

type Response interface {
	GetError() error
}

type SyncResp struct {
	Error error
}

type ConnectResourcesResp struct {
	Resp SyncResp
}

func (r ConnectResourcesResp) GetError() error {
	return r.Resp.Error
}

type DisconnectResourcesResp struct {
	Resp SyncResp
}

func (r DisconnectResourcesResp) GetError() error {
	return r.Resp.Error
}

type UpsertAttributeResp struct {
	Resp SyncResp
}

func (r UpsertAttributeResp) GetError() error {
	return r.Resp.Error
}

type RemoveAttributeResp struct {
	Resp SyncResp
}

func (r RemoveAttributeResp) GetError() error {
	return r.Resp.Error
}

type InsertPermissionResp struct {
	Resp SyncResp
}

func (r InsertPermissionResp) GetError() error {
	return r.Resp.Error
}

type RemovePermissionResp struct {
	Resp SyncResp
}

func (r RemovePermissionResp) GetError() error {
	return r.Resp.Error
}
