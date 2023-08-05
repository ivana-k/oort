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

func (h Handler) PutAttribute(req PutAttributeReq) SyncResp {
	resp := h.store.PutAttribute(acl.PutAttributeReq{
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

func (h Handler) CreateInheritanceRel(req CreateInheritanceRelReq) SyncResp {
	resp := h.store.CreateInheritanceRel(acl.CreateInheritanceRelReq{
		From:     req.From,
		To:       req.To,
		Callback: h.outboxMessageCallback(req),
	})
	return SyncResp{
		Error: resp.Error,
	}
}

func (h Handler) DeleteInheritanceRel(req DeleteInheritanceRelReq) SyncResp {
	resp := h.store.DeleteInheritanceRel(acl.DeleteInheritanceRelReq{
		From:     req.From,
		To:       req.To,
		Callback: h.outboxMessageCallback(req),
	})
	return SyncResp{
		Error: resp.Error,
	}
}

func (h Handler) CreatePolicy(req CreatePolicyReq) SyncResp {
	if req.SubjectScope == nil {
		req.SubjectScope = &model.RootResource
	}
	if req.ObjectScope == nil {
		req.ObjectScope = &model.RootResource
	}
	resp := h.store.CreatePolicy(acl.CreatePolicyReq{
		SubjectScope: *req.SubjectScope,
		ObjectScope:  *req.ObjectScope,
		Permission:   req.Permission,
		Callback:     h.outboxMessageCallback(req),
	})
	return SyncResp{
		Error: resp.Error,
	}
}

func (h Handler) DeletePolicy(req DeletePolicyReq) SyncResp {
	if req.SubjectScope == nil {
		req.SubjectScope = &model.RootResource
	}
	if req.ObjectScope == nil {
		req.ObjectScope = &model.RootResource
	}
	resp := h.store.DeletePolicy(acl.DeletePolicyReq{
		SubjectScope: *req.SubjectScope,
		ObjectScope:  *req.ObjectScope,
		Permission:   req.Permission,
		Callback:     h.outboxMessageCallback(req),
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
