package syncer

import (
	"github.com/c12s/oort/domain/model"
	"github.com/c12s/oort/domain/store/acl"
)

type Handler struct {
	store           acl.Store
	syncRespFactory func(string, string, bool) model.OutboxMessage
}

func NewHandler(store acl.Store, syncRespFactory func(string, string, bool) model.OutboxMessage) Handler {
	return Handler{
		store:           store,
		syncRespFactory: syncRespFactory,
	}
}

func (h Handler) CreateResource(req CreateResourceReq) SyncResp {
	resp := h.store.CreateResource(acl.CreateResourceReq{
		Resource: req.Resource,
		Callback: h.outboxMessageCallback(req),
	})
	return SyncResp{
		Error: resp.Error,
	}
}

func (h Handler) DeleteResource(req DeleteResourceReq) SyncResp {
	resp := h.store.DeleteResource(acl.DeleteResourceReq{
		Resource: req.Resource,
		Callback: h.outboxMessageCallback(req),
	})
	return SyncResp{
		Error: resp.Error,
	}
}

func (h Handler) CreateAttribute(req CreateAttributeReq) SyncResp {
	resp := h.store.CreateAttribute(acl.CreateAttributeReq{
		Resource:  req.Resource,
		Attribute: req.Attribute,
		Callback:  h.outboxMessageCallback(req),
	})
	return SyncResp{
		Error: resp.Error,
	}
}

func (h Handler) UpdateAttribute(req UpdateAttributeReq) SyncResp {
	resp := h.store.UpdateAttribute(acl.UpdateAttributeReq{
		Resource:  req.Resource,
		Attribute: req.Attribute,
		Callback:  h.outboxMessageCallback(req),
	})
	return SyncResp{
		Error: resp.Error,
	}
}

func (h Handler) DeleteAttribute(req DeleteAttributeReq) SyncResp {
	resp := h.store.DeleteAttribute(acl.DeleteAttributeReq{
		Resource:    req.Resource,
		AttributeId: req.AttributeId,
		Callback:    h.outboxMessageCallback(req),
	})
	return SyncResp{
		Error: resp.Error,
	}
}

func (h Handler) CreateAggregationRelReq(req CreateAggregationRelReq) SyncResp {
	resp := h.store.CreateAggregationRel(acl.CreateAggregationRelReq{
		Parent:   req.Parent,
		Child:    req.Child,
		Callback: h.outboxMessageCallback(req),
	})
	return SyncResp{
		Error: resp.Error,
	}
}

func (h Handler) DeleteAggregationRelReq(req DeleteAggregationRelReq) SyncResp {
	resp := h.store.DeleteAggregationRel(acl.DeleteAggregationRelReq{
		Parent:   req.Parent,
		Child:    req.Child,
		Callback: h.outboxMessageCallback(req),
	})
	return SyncResp{
		Error: resp.Error,
	}
}

func (h Handler) CreateCompositionRelReq(req CreateCompositionRelReq) SyncResp {
	resp := h.store.CreateCompositionRel(acl.CreateCompositionRelReq{
		Parent:   req.Parent,
		Child:    req.Child,
		Callback: h.outboxMessageCallback(req),
	})
	return SyncResp{
		Error: resp.Error,
	}
}

func (h Handler) DeleteCompositionRelReq(req DeleteCompositionRelReq) SyncResp {
	resp := h.store.DeleteCompositionRel(acl.DeleteCompositionRelReq{
		Parent:   req.Parent,
		Child:    req.Child,
		Callback: h.outboxMessageCallback(req),
	})
	return SyncResp{
		Error: resp.Error,
	}
}

func (h Handler) CreatePermission(req CreatePermissionReq) SyncResp {
	if req.Subject == nil {
		req.Subject = &model.RootResource
	}
	if req.Object == nil {
		req.Object = &model.RootResource
	}
	resp := h.store.CreatePermission(acl.CreatePermissionReq{
		Subject:    *req.Subject,
		Object:     *req.Object,
		Permission: req.Permission,
		Callback:   h.outboxMessageCallback(req),
	})
	return SyncResp{
		Error: resp.Error,
	}
}

func (h Handler) DeletePermission(req DeletePermissionReq) SyncResp {
	if req.Subject == nil {
		req.Subject = &model.RootResource
	}
	if req.Object == nil {
		req.Object = &model.RootResource
	}
	resp := h.store.DeletePermission(acl.DeletePermissionReq{
		Subject:    *req.Subject,
		Object:     *req.Object,
		Permission: req.Permission,
		Callback:   h.outboxMessageCallback(req),
	})
	return SyncResp{
		Error: resp.Error,
	}
}

func (h Handler) outboxMessageCallback(req Request) func(error) model.OutboxMessage {
	return func(err error) model.OutboxMessage {
		if err != nil {
			return h.syncRespFactory(req.Id(), err.Error(), false)
		}
		return h.syncRespFactory(req.Id(), "", true)
	}
}
