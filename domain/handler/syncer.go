package handler

import (
	"github.com/c12s/oort/domain/model/syncer"
	"github.com/c12s/oort/domain/store"
)

type SyncerHandler struct {
	store store.AclStore
}

func NewSyncerHandler(store store.AclStore) SyncerHandler {
	return SyncerHandler{
		store: store,
	}
}

func (h SyncerHandler) ConnectResources(req syncer.ConnectResourcesReq) syncer.ConnectResourcesResp {
	return h.store.ConnectResources(req)
}

func (h SyncerHandler) DisconnectResources(req syncer.DisconnectResourcesReq) syncer.DisconnectResourcesResp {
	return h.store.DisconnectResources(req)
}

func (h SyncerHandler) UpsertAttribute(req syncer.UpsertAttributeReq) syncer.UpsertAttributeResp {
	return h.store.UpsertAttribute(req)
}

func (h SyncerHandler) RemoveAttribute(req syncer.RemoveAttributeReq) syncer.RemoveAttributeResp {
	return h.store.RemoveAttribute(req)
}

func (h SyncerHandler) InsertPermission(req syncer.InsertPermissionReq) syncer.InsertPermissionResp {
	return h.store.InsertPermission(req)
}

func (h SyncerHandler) RemovePermission(req syncer.RemovePermissionReq) syncer.RemovePermissionResp {
	return h.store.RemovePermission(req)
}
