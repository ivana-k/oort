package api

import (
	"google.golang.org/protobuf/proto"
)

type AdministrationReq interface {
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
	Kind() AdministrationAsyncReq_ReqKind
}

func (x *CreateResourceReq) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *CreateResourceReq) Unmarshal(marshalled []byte) error {
	return proto.Unmarshal(marshalled, x)
}

func (x *CreateResourceReq) Kind() AdministrationAsyncReq_ReqKind {
	return AdministrationAsyncReq_CreateResource
}

func (x *DeleteResourceReq) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *DeleteResourceReq) Unmarshal(marshalled []byte) error {
	return proto.Unmarshal(marshalled, x)
}

func (x *DeleteResourceReq) Kind() AdministrationAsyncReq_ReqKind {
	return AdministrationAsyncReq_DeleteResource
}

func (x *PutAttributeReq) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *PutAttributeReq) Unmarshal(marshalled []byte) error {
	return proto.Unmarshal(marshalled, x)
}

func (x *PutAttributeReq) Kind() AdministrationAsyncReq_ReqKind {
	return AdministrationAsyncReq_PutAttribute
}

func (x *DeleteAttributeReq) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *DeleteAttributeReq) Unmarshal(marshalled []byte) error {
	return proto.Unmarshal(marshalled, x)
}

func (x *DeleteAttributeReq) Kind() AdministrationAsyncReq_ReqKind {
	return AdministrationAsyncReq_DeleteAttribute
}

func (x *CreateInheritanceRelReq) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *CreateInheritanceRelReq) Unmarshal(marshalled []byte) error {
	return proto.Unmarshal(marshalled, x)
}

func (x *CreateInheritanceRelReq) Kind() AdministrationAsyncReq_ReqKind {
	return AdministrationAsyncReq_CreateInheritanceRel
}

func (x *DeleteInheritanceRelReq) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *DeleteInheritanceRelReq) Unmarshal(marshalled []byte) error {
	return proto.Unmarshal(marshalled, x)
}

func (x *DeleteInheritanceRelReq) Kind() AdministrationAsyncReq_ReqKind {
	return AdministrationAsyncReq_DeleteInheritanceRel
}

func (x *CreatePolicyReq) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *CreatePolicyReq) Unmarshal(marshalled []byte) error {
	return proto.Unmarshal(marshalled, x)
}

func (x *CreatePolicyReq) Kind() AdministrationAsyncReq_ReqKind {
	return AdministrationAsyncReq_CreatePolicy
}

func (x *DeletePolicyReq) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *DeletePolicyReq) Unmarshal(marshalled []byte) error {
	return proto.Unmarshal(marshalled, x)
}

func (x *DeletePolicyReq) Kind() AdministrationAsyncReq_ReqKind {
	return AdministrationAsyncReq_DeletePolicy
}

func (x *AdministrationAsyncReq) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *AdministrationAsyncReq) Unmarshal(marshalled []byte) error {
	return proto.Unmarshal(marshalled, x)
}

func (x *AdministrationAsyncResp) Marshal() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *AdministrationAsyncResp) Unmarshal(marshalled []byte) error {
	return proto.Unmarshal(marshalled, x)
}
