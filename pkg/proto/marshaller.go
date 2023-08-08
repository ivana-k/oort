package proto

import (
	"github.com/c12s/oort/internal/domain"
	"google.golang.org/protobuf/proto"
)

type administrationReqMarshaller struct {
}

func NewAdministrationReqMarshaller() domain.AdministrationReqMarshaller {
	return administrationReqMarshaller{}
}

func (s administrationReqMarshaller) Marshal(req domain.AdministrationReq) ([]byte, error) {
	protoReq := &AdministrationReq{}
	protoReq, err := protoReq.FromDomain(req)
	if err != nil {
		return nil, err
	}
	return proto.Marshal(protoReq)
}

func (s administrationReqMarshaller) Unmarshal(bytes []byte) (*domain.AdministrationReq, error) {
	req := &AdministrationReq{}
	err := proto.Unmarshal(bytes, req)
	if err != nil {
		return nil, err
	}
	return req.ToDomain()
}
