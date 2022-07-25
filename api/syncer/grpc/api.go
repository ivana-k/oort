package grpc

import (
	"context"
	"github.com/c12s/oort/domain/handler"
	"github.com/c12s/oort/proto/syncerpb"
)

type SyncerGrpcApi struct {
	syncerpb.UnsafeSyncerServiceServer
	handler handler.SyncerHandler
}

func NewSyncerGrpcApi(handler handler.SyncerHandler) SyncerGrpcApi {
	return SyncerGrpcApi{
		handler: handler,
	}
}

func (s SyncerGrpcApi) ConnectResources(ctx context.Context, req *syncerpb.ConnectResourcesReq) (*syncerpb.SyncResp, error) {
	resp := s.handler.ConnectResources(req.MapToDomain())
	return &syncerpb.SyncResp{}, resp.Resp.Error
}

func (s SyncerGrpcApi) DisconnectResources(ctx context.Context, req *syncerpb.DisconnectResourcesReq) (*syncerpb.SyncResp, error) {
	resp := s.handler.DisconnectResources(req.MapToDomain())
	return &syncerpb.SyncResp{}, resp.Resp.Error
}

func (s SyncerGrpcApi) UpsertAttribute(ctx context.Context, req *syncerpb.UpsertAttributeReq) (*syncerpb.SyncResp, error) {
	domainReq, err := req.MapToDomain()
	if err != nil {
		return &syncerpb.SyncResp{}, err
	}
	resp := s.handler.UpsertAttribute(domainReq)
	return &syncerpb.SyncResp{}, resp.Resp.Error
}

func (s SyncerGrpcApi) RemoveAttribute(ctx context.Context, req *syncerpb.RemoveAttributeReq) (*syncerpb.SyncResp, error) {
	resp := s.handler.RemoveAttribute(req.MapToDomain())
	return &syncerpb.SyncResp{}, resp.Resp.Error
}

func (s SyncerGrpcApi) InsertPermission(ctx context.Context, req *syncerpb.InsertPermissionReq) (*syncerpb.SyncResp, error) {
	resp := s.handler.InsertPermission(req.MapToDomain())
	return &syncerpb.SyncResp{}, resp.Resp.Error
}

func (s SyncerGrpcApi) RemovePermission(ctx context.Context, req *syncerpb.RemovePermissionReq) (*syncerpb.SyncResp, error) {
	resp := s.handler.RemovePermission(req.MapToDomain())
	return &syncerpb.SyncResp{}, resp.Resp.Error
}
