package syncer

import (
	"github.com/c12s/oort/domain/model"
	"github.com/c12s/oort/domain/store/acl"
)

type Handler struct {
	store           acl.Store
	syncRespFactory func(string, string, bool) *model.OutboxMessage
}

func NewHandler(store acl.Store, syncRespFactory func(string, string, bool) *model.OutboxMessage) Handler {
	return Handler{
		store:           store,
		syncRespFactory: syncRespFactory,
	}
}

func (h Handler) ConnectResources(req ConnectResourcesReq) SyncResp {
	resp := h.store.ConnectResources(acl.ConnectResourcesReq{
		Parent:   req.Parent,
		Child:    req.Child,
		Callback: h.outboxMessageCallback(req),
	})
	return SyncResp{
		Error: resp.Error,
	}
}

func (h Handler) DisconnectResources(req DisconnectResourcesReq) SyncResp {
	resp := h.store.DisconnectResources(acl.DisconnectResourcesReq{
		Parent:   req.Parent,
		Child:    req.Child,
		Callback: h.outboxMessageCallback(req),
	})
	return SyncResp{
		Error: resp.Error,
	}
}

func (h Handler) UpsertAttribute(req UpsertAttributeReq) SyncResp {
	resp := h.store.UpsertAttribute(acl.UpsertAttributeReq{
		Resource:  req.Resource,
		Attribute: req.Attribute,
		Callback:  h.outboxMessageCallback(req),
	})
	return SyncResp{
		Error: resp.Error,
	}
}

func (h Handler) RemoveAttribute(req RemoveAttributeReq) SyncResp {
	resp := h.store.RemoveAttribute(acl.RemoveAttributeReq{
		Resource:    req.Resource,
		AttributeId: req.AttributeId,
		Callback:    h.outboxMessageCallback(req),
	})
	return SyncResp{
		Error: resp.Error,
	}
}

func (h Handler) InsertPermission(req InsertPermissionReq) SyncResp {
	resp := h.store.InsertPermission(acl.InsertPermissionReq{
		Principal:  req.Principal,
		Resource:   req.Resource,
		Permission: req.Permission,
		Callback:   h.outboxMessageCallback(req),
	})
	return SyncResp{
		Error: resp.Error,
	}
}

func (h Handler) RemovePermission(req RemovePermissionReq) SyncResp {
	resp := h.store.RemovePermission(acl.RemovePermissionReq{
		Principal:  req.Principal,
		Resource:   req.Resource,
		Permission: req.Permission,
		Callback:   h.outboxMessageCallback(req),
	})
	return SyncResp{
		Error: resp.Error,
	}
}

func (h Handler) outboxMessageCallback(req Request) func(error) *model.OutboxMessage {
	return func(err error) *model.OutboxMessage {
		if err != nil {
			return h.syncRespFactory(req.Id(), err.Error(), false)
		}
		return h.syncRespFactory(req.Id(), "", true)
	}
}
