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

func NewSyncerGrpcApi(handler syncer.Handler) SyncerGrpcApi {
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

func (s SyncerGrpcApi) CreateAggregationRel(ctx context.Context, req *syncerpb.CreateAggregationRelReq) (*syncerpb.SyncResp, error) {
	request, err := req.MapToDomain()
	if err != nil {
		return nil, err
	}
	resp := s.handler.CreateAggregationRelReq(request.(syncer.CreateAggregationRelReq))
	return &syncerpb.SyncResp{}, resp.Error
}

func (s SyncerGrpcApi) DeleteAggregationRel(ctx context.Context, req *syncerpb.DeleteAggregationRelReq) (*syncerpb.SyncResp, error) {
	request, err := req.MapToDomain()
	if err != nil {
		return nil, err
	}
	resp := s.handler.DeleteAggregationRelReq(request.(syncer.DeleteAggregationRelReq))
	return &syncerpb.SyncResp{}, resp.Error
}

func (s SyncerGrpcApi) CreateCompositionRel(ctx context.Context, req *syncerpb.CreateCompositionRelReq) (*syncerpb.SyncResp, error) {
	request, err := req.MapToDomain()
	if err != nil {
		return nil, err
	}
	resp := s.handler.CreateCompositionRelReq(request.(syncer.CreateCompositionRelReq))
	return &syncerpb.SyncResp{}, resp.Error
}

func (s SyncerGrpcApi) DeleteCompositionRel(ctx context.Context, req *syncerpb.DeleteCompositionRelReq) (*syncerpb.SyncResp, error) {
	request, err := req.MapToDomain()
	if err != nil {
		return nil, err
	}
	resp := s.handler.DeleteCompositionRelReq(request.(syncer.DeleteCompositionRelReq))
	return &syncerpb.SyncResp{}, resp.Error
}

func (s SyncerGrpcApi) CreateAttribute(ctx context.Context, req *syncerpb.CreateAttributeReq) (*syncerpb.SyncResp, error) {
	request, err := req.MapToDomain()
	if err != nil {
		return nil, err
	}
	resp := s.handler.CreateAttribute(request.(syncer.CreateAttributeReq))
	return &syncerpb.SyncResp{}, resp.Error
}

func (s SyncerGrpcApi) UpdateAttribute(ctx context.Context, req *syncerpb.UpdateAttributeReq) (*syncerpb.SyncResp, error) {
	request, err := req.MapToDomain()
	if err != nil {
		return nil, err
	}
	resp := s.handler.UpdateAttribute(request.(syncer.UpdateAttributeReq))
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

func (s SyncerGrpcApi) CreatePermission(ctx context.Context, req *syncerpb.CreatePermissionReq) (*syncerpb.SyncResp, error) {
	request, err := req.MapToDomain()
	if err != nil {
		return nil, err
	}
	resp := s.handler.CreatePermission(request.(syncer.CreatePermissionReq))
	return &syncerpb.SyncResp{}, resp.Error
}

func (s SyncerGrpcApi) DeletePermission(ctx context.Context, req *syncerpb.DeletePermissionReq) (*syncerpb.SyncResp, error) {
	request, err := req.MapToDomain()
	if err != nil {
		return nil, err
	}
	resp := s.handler.DeletePermission(request.(syncer.DeletePermissionReq))
	return &syncerpb.SyncResp{}, resp.Error
}
