package handlers

import (
	"context"
	"github.com/c12s/oort/internal/services"
	"github.com/c12s/oort/pkg/proto"
)

type oortAdministratorGrpcServer struct {
	proto.UnimplementedOortAdministratorServer
	service services.AdministrationService
}

func NewOortAdministratorGrpcServer(service services.AdministrationService) (proto.OortAdministratorServer, error) {
	return &oortAdministratorGrpcServer{
		service: service,
	}, nil
}

func (o *oortAdministratorGrpcServer) CreateResource(ctx context.Context, req *proto.CreateResourceReq) (*proto.AdministrationResp, error) {
	request, err := req.ToDomain()
	if err != nil {
		return nil, err
	}
	resp := o.service.CreateResource(*request)
	return &proto.AdministrationResp{}, resp.Error
}

func (o *oortAdministratorGrpcServer) DeleteResource(ctx context.Context, req *proto.DeleteResourceReq) (*proto.AdministrationResp, error) {
	request, err := req.ToDomain()
	if err != nil {
		return nil, err
	}
	resp := o.service.DeleteResource(*request)
	return &proto.AdministrationResp{}, resp.Error
}

func (o *oortAdministratorGrpcServer) CreateInheritanceRel(ctx context.Context, req *proto.CreateInheritanceRelReq) (*proto.AdministrationResp, error) {
	request, err := req.ToDomain()
	if err != nil {
		return nil, err
	}
	resp := o.service.CreateInheritanceRel(*request)
	return &proto.AdministrationResp{}, resp.Error
}

func (o *oortAdministratorGrpcServer) DeleteInheritanceRel(ctx context.Context, req *proto.DeleteInheritanceRelReq) (*proto.AdministrationResp, error) {
	request, err := req.ToDomain()
	if err != nil {
		return nil, err
	}
	resp := o.service.DeleteInheritanceRel(*request)
	return &proto.AdministrationResp{}, resp.Error
}

func (o *oortAdministratorGrpcServer) PutAttribute(ctx context.Context, req *proto.PutAttributeReq) (*proto.AdministrationResp, error) {
	request, err := req.ToDomain()
	if err != nil {
		return nil, err
	}
	resp := o.service.PutAttribute(*request)
	return &proto.AdministrationResp{}, resp.Error
}

func (o *oortAdministratorGrpcServer) DeleteAttribute(ctx context.Context, req *proto.DeleteAttributeReq) (*proto.AdministrationResp, error) {
	request, err := req.ToDomain()
	if err != nil {
		return nil, err
	}
	resp := o.service.DeleteAttribute(*request)
	return &proto.AdministrationResp{}, resp.Error
}

func (o *oortAdministratorGrpcServer) CreatePolicy(ctx context.Context, req *proto.CreatePolicyReq) (*proto.AdministrationResp, error) {
	request, err := req.ToDomain()
	if err != nil {
		return nil, err
	}
	resp := o.service.CreatePolicy(*request)
	return &proto.AdministrationResp{}, resp.Error
}

func (o *oortAdministratorGrpcServer) DeletePolicy(ctx context.Context, req *proto.DeletePolicyReq) (*proto.AdministrationResp, error) {
	request, err := req.ToDomain()
	if err != nil {
		return nil, err
	}
	resp := o.service.DeletePolicy(*request)
	return &proto.AdministrationResp{}, resp.Error
}
