package syncerpb

import (
	"github.com/c12s/oort/domain/model"
	"github.com/c12s/oort/domain/syncer"
	"google.golang.org/protobuf/proto"
	"log"
)

type protoRequest interface {
	MapToDomain() (syncer.Request, error)
}

func (x *CreateResourceReq) MapToDomain() (syncer.Request, error) {
	return syncer.CreateResourceReq{
		ReqId:    x.Id,
		Resource: x.Resource.MapToDomain(),
	}, nil
}

func (x *DeleteResourceReq) MapToDomain() (syncer.Request, error) {
	return syncer.DeleteResourceReq{
		ReqId:    x.Id,
		Resource: x.Resource.MapToDomain(),
	}, nil
}

func (x *PutAttributeReq) MapToDomain() (syncer.Request, error) {
	attr, err := x.Attribute.MapToDomain()
	if err != nil {
		return syncer.PutAttributeReq{}, err
	}
	return syncer.PutAttributeReq{
		ReqId:     x.Id,
		Resource:  x.Resource.MapToDomain(),
		Attribute: attr,
	}, nil
}

func (x *DeleteAttributeReq) MapToDomain() (syncer.Request, error) {
	return syncer.DeleteAttributeReq{
		ReqId:       x.Id,
		Resource:    x.Resource.MapToDomain(),
		AttributeId: x.AttributeId.MapToDomain(),
	}, nil
}

func (x *CreateInheritanceRelReq) MapToDomain() (syncer.Request, error) {
	return syncer.CreateInheritanceRelReq{
		ReqId: x.Id,
		From:  x.From.MapToDomain(),
		To:    x.To.MapToDomain(),
	}, nil
}

func (x *DeleteInheritanceRelReq) MapToDomain() (syncer.Request, error) {
	return syncer.DeleteInheritanceRelReq{
		ReqId: x.Id,
		From:  x.From.MapToDomain(),
		To:    x.To.MapToDomain(),
	}, nil
}

func (x *CreatePolicyReq) MapToDomain() (syncer.Request, error) {
	permission, err := x.Permission.MapToDomain()
	if err != nil {
		return nil, err
	}
	subjectScope := x.SubjectScope.MapToDomain()
	objectScope := x.ObjectScope.MapToDomain()
	return syncer.CreatePolicyReq{
		ReqId:        x.Id,
		SubjectScope: &subjectScope,
		ObjectScope:  &objectScope,
		Permission:   permission,
	}, nil
}

func (x *DeletePolicyReq) MapToDomain() (syncer.Request, error) {
	permission, err := x.Permission.MapToDomain()
	if err != nil {
		return nil, err
	}
	subjectScope := x.SubjectScope.MapToDomain()
	objectScope := x.ObjectScope.MapToDomain()
	return syncer.DeletePolicyReq{
		ReqId:        x.Id,
		SubjectScope: &subjectScope,
		ObjectScope:  &objectScope,
		Permission:   permission,
	}, nil
}

func NewSyncRespOutboxMessage(reqId string, error string, successful bool) model.OutboxMessage {
	resp := AsyncSyncResp{
		ReqId:      reqId,
		Error:      error,
		Successful: successful,
	}
	payload, err := proto.Marshal(&resp)
	if err != nil {
		log.Println(err)
		return model.OutboxMessage{}
	}
	return model.OutboxMessage{
		Kind:    model.SyncRespOutboxMessageKind,
		Payload: payload,
	}
}
