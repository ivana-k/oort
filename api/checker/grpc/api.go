package grpc

import (
	"context"
	"github.com/c12s/oort/domain/handler"
	"github.com/c12s/oort/proto/checkerpb"
)

type CheckerGrpcApi struct {
	handler handler.CheckerHandler
	checkerpb.UnimplementedCheckerServiceServer
}

func NewCheckerGrpcApi(handler handler.CheckerHandler) CheckerGrpcApi {
	return CheckerGrpcApi{
		handler: handler,
	}
}

func (c CheckerGrpcApi) CheckPermission(ctx context.Context, req *checkerpb.CheckPermissionReq) (*checkerpb.CheckResp, error) {
	resp := c.handler.CheckPermission(req.MapToDomain())
	return &checkerpb.CheckResp{Allowed: resp.Allowed}, resp.Error
}
