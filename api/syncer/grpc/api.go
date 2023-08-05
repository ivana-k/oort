package grpc

import (
	"context"
	"github.com/c12s/oort/domain/syncer"
	"github.com/c12s/oort/proto/syncerpb"
)

type SyncerGrpcApi struct {
	syncerpb.UnsafeSyncerServiceServer
	handler syncer.Handler
}

func NewSyncerGrpcApi(handler syncer.Handler) syncerpb.SyncerServiceServer {
	return SyncerGrpcApi{
		handler: handler,
	}
}

func (s SyncerGrpcApi) CreateResource(ctx context.Context, req *syncerpb.CreateResourceReq) (*syncerpb.SyncResp, error) {
	request, err := req.MapToDomain()
	if err != nil {
		return nil, err
	}
	resp := s.handler.CreateResource(request.(syncer.CreateResourceReq))
	return &syncerpb.SyncResp{}, resp.Error
}

func (s SyncerGrpcApi) DeleteResource(ctx context.Context, req *syncerpb.DeleteResourceReq) (*syncerpb.SyncResp, error) {
	request, err := req.MapToDomain()
	if err != nil {
		return nil, err
	}
	resp := s.handler.DeleteResource(request.(syncer.DeleteResourceReq))
	return &syncerpb.SyncResp{}, resp.Error
}

func (s SyncerGrpcApi) CreateInheritanceRel(ctx context.Context, req *syncerpb.CreateInheritanceRelReq) (*syncerpb.SyncResp, error) {
	request, err := req.MapToDomain()
	if err != nil {
		return nil, err
	}
	resp := s.handler.CreateInheritanceRel(request.(syncer.CreateInheritanceRelReq))
	return &syncerpb.SyncResp{}, resp.Error
}

func (s SyncerGrpcApi) DeleteInheritanceRel(ctx context.Context, req *syncerpb.DeleteInheritanceRelReq) (*syncerpb.SyncResp, error) {
	request, err := req.MapToDomain()
	if err != nil {
		return nil, err
	}
	resp := s.handler.DeleteInheritanceRel(request.(syncer.DeleteInheritanceRelReq))
	return &syncerpb.SyncResp{}, resp.Error
}

func (s SyncerGrpcApi) PutAttribute(ctx context.Context, req *syncerpb.PutAttributeReq) (*syncerpb.SyncResp, error) {
	request, err := req.MapToDomain()
	if err != nil {
		return nil, err
	}
	resp := s.handler.PutAttribute(request.(syncer.PutAttributeReq))
	return &syncerpb.SyncResp{}, resp.Error
}

func (s SyncerGrpcApi) DeleteAttribute(ctx context.Context, req *syncerpb.DeleteAttributeReq) (*syncerpb.SyncResp, error) {
	request, err := req.MapToDomain()
	if err != nil {
		return nil, err
	}
	resp := s.handler.DeleteAttribute(request.(syncer.DeleteAttributeReq))
	return &syncerpb.SyncResp{}, resp.Error
}

func (s SyncerGrpcApi) CreatePolicy(ctx context.Context, req *syncerpb.CreatePolicyReq) (*syncerpb.SyncResp, error) {
	request, err := req.MapToDomain()
	if err != nil {
		return nil, err
	}
	resp := s.handler.CreatePolicy(request.(syncer.CreatePolicyReq))
	return &syncerpb.SyncResp{}, resp.Error
}

func (s SyncerGrpcApi) DeletePolicy(ctx context.Context, req *syncerpb.DeletePolicyReq) (*syncerpb.SyncResp, error) {
	request, err := req.MapToDomain()
	if err != nil {
		return nil, err
	}
	resp := s.handler.DeletePolicy(request.(syncer.DeletePolicyReq))
	return &syncerpb.SyncResp{}, resp.Error
}
