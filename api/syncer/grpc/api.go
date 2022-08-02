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

func (s SyncerGrpcApi) ConnectResources(ctx context.Context, req *syncerpb.ConnectResourcesReq) (*syncerpb.SyncResp, error) {
	request, err := req.MapToDomain()
	if err != nil {
		return nil, err
	}
	resp := s.handler.ConnectResources(request.(syncer.ConnectResourcesReq))
	return &syncerpb.SyncResp{}, resp.Error
}

func (s SyncerGrpcApi) DisconnectResources(ctx context.Context, req *syncerpb.DisconnectResourcesReq) (*syncerpb.SyncResp, error) {
	request, err := req.MapToDomain()
	if err != nil {
		return nil, err
	}
	resp := s.handler.DisconnectResources(request.(syncer.DisconnectResourcesReq))
	return &syncerpb.SyncResp{}, resp.Error
}

func (s SyncerGrpcApi) UpsertAttribute(ctx context.Context, req *syncerpb.UpsertAttributeReq) (*syncerpb.SyncResp, error) {
	domainReq, err := req.MapToDomain()
	if err != nil {
		return &syncerpb.SyncResp{}, err
	}
	resp := s.handler.UpsertAttribute(domainReq.(syncer.UpsertAttributeReq))
	return &syncerpb.SyncResp{}, resp.Error
}

func (s SyncerGrpcApi) RemoveAttribute(ctx context.Context, req *syncerpb.RemoveAttributeReq) (*syncerpb.SyncResp, error) {
	request, err := req.MapToDomain()
	if err != nil {
		return nil, err
	}
	resp := s.handler.RemoveAttribute(request.(syncer.RemoveAttributeReq))
	return &syncerpb.SyncResp{}, resp.Error
}

func (s SyncerGrpcApi) InsertPermission(ctx context.Context, req *syncerpb.InsertPermissionReq) (*syncerpb.SyncResp, error) {
	request, err := req.MapToDomain()
	if err != nil {
		return nil, err
	}
	resp := s.handler.InsertPermission(request.(syncer.InsertPermissionReq))
	return &syncerpb.SyncResp{}, resp.Error
}

func (s SyncerGrpcApi) RemovePermission(ctx context.Context, req *syncerpb.RemovePermissionReq) (*syncerpb.SyncResp, error) {
	request, err := req.MapToDomain()
	if err != nil {
		return nil, err
	}
	resp := s.handler.RemovePermission(request.(syncer.RemovePermissionReq))
	return &syncerpb.SyncResp{}, resp.Error
}
