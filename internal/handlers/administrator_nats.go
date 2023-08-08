package handlers

import (
	"errors"
	"github.com/c12s/oort/internal/domain"
	"github.com/c12s/oort/internal/services"
)

type AdministrationNatsHandler struct {
	marshaller domain.AdministrationReqMarshaller
	service    services.AdministrationService
}

func NewAdministrationNatsHandler(subscriber services.Subscriber, subject, queueGroup string, marshaller domain.AdministrationReqMarshaller, service services.AdministrationService) error {
	s := AdministrationNatsHandler{
		marshaller: marshaller,
		service:    service,
	}
	return subscriber.Subscribe(subject, queueGroup, s.handle)
}

func (s AdministrationNatsHandler) handle(reqMarshalled []byte) error {
	req, err := s.marshaller.Unmarshal(reqMarshalled)
	if err != nil {
		return err
	}
	var respError error
	switch req.ReqKind {
	case domain.CreateResource:
		s.service.CreateResource(req.Request.(domain.CreateResourceReq))
	case domain.DeleteResource:
		s.service.DeleteResource(req.Request.(domain.DeleteResourceReq))
	case domain.PutAttribute:
		s.service.PutAttribute(req.Request.(domain.PutAttributeReq))
	case domain.DeleteAttribute:
		s.service.DeleteAttribute(req.Request.(domain.DeleteAttributeReq))
	case domain.CreateInheritanceRel:
		s.service.CreateInheritanceRel(req.Request.(domain.CreateInheritanceRelReq))
	case domain.DeleteInheritanceRel:
		s.service.DeleteInheritanceRel(req.Request.(domain.DeleteInheritanceRelReq))
	case domain.CreatePolicy:
		s.service.CreatePolicy(req.Request.(domain.CreatePolicyReq))
	case domain.DeletePolicy:
		s.service.DeletePolicy(req.Request.(domain.DeletePolicyReq))
	default:
		respError = errors.New("unknown request kind")
	}
	return respError
}
