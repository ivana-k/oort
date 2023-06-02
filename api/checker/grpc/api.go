package grpc

import (
	"context"
	"github.com/c12s/oort/domain/checker"
	"github.com/c12s/oort/proto/checkerpb"
	"log"
)

type CheckerGrpcApi struct {
	handler checker.Handler
	checkerpb.UnimplementedCheckerServiceServer
}

func NewCheckerGrpcApi(handler checker.Handler) CheckerGrpcApi {
	return CheckerGrpcApi{
		handler: handler,
	}
}

func (c CheckerGrpcApi) CheckPermission(ctx context.Context, req *checkerpb.CheckPermissionReq) (*checkerpb.CheckResp, error) {
	resp := c.handler.CheckPermission(req.MapToDomain())
	log.Println(resp)
	return &checkerpb.CheckResp{Allowed: resp.Allowed}, resp.Error
}
