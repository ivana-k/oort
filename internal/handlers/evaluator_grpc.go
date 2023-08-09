package handlers

import (
	"context"
	"github.com/c12s/oort/internal/mappers/proto"
	"github.com/c12s/oort/internal/services"
	"github.com/c12s/oort/pkg/api"
	"log"
)

type oortEvaluatorGrpcServer struct {
	service services.EvaluationService
	api.UnimplementedOortEvaluatorServer
}

func NewOortEvaluatorGrpcServer(service services.EvaluationService) (api.OortEvaluatorServer, error) {
	return &oortEvaluatorGrpcServer{
		service: service,
	}, nil
}

func (o *oortEvaluatorGrpcServer) Authorize(ctx context.Context, req *api.AuthorizationReq) (*api.AuthorizationResp, error) {
	reqDomain, err := proto.AuthorizationReqToDomain(req)
	if err != nil {
		return nil, err
	}
	resp := o.service.Authorize(*reqDomain)
	log.Println(resp)
	return &api.AuthorizationResp{Authorized: resp.Authorized}, resp.Error
}
