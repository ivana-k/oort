package handlers

import (
	"context"
	"github.com/c12s/oort/internal/services"
	"github.com/c12s/oort/pkg/proto"
	"log"
)

type oortEvaluatorGrpcServer struct {
	service services.EvaluationService
	proto.UnimplementedOortEvaluatorServer
}

func NewOortEvaluatorGrpcServer(service services.EvaluationService) (proto.OortEvaluatorServer, error) {
	return &oortEvaluatorGrpcServer{
		service: service,
	}, nil
}

func (o *oortEvaluatorGrpcServer) Authorize(ctx context.Context, req *proto.AuthorizationReq) (*proto.AuthorizationResp, error) {
	reqDomain, err := req.ToDomain()
	if err != nil {
		return nil, err
	}
	log.Println(reqDomain.Subject.Name())
	log.Println(reqDomain.Object.Name())
	resp := o.service.Authorize(*reqDomain)
	log.Println(resp)
	return &proto.AuthorizationResp{Allowed: resp.Allowed}, resp.Error
}
