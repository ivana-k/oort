package async

import (
	"github.com/c12s/oort/domain/syncer"
)

type SyncMsgKind int

const (
	CreateResource SyncMsgKind = iota
	DeleteResource
	CreateAttribute
	UpdateAttribute
	DeleteAttribute
	CreateAggregationRel
	DeleteAggregationRel
	CreateCompositionRel
	DeleteCompositionRel
	CreatePermission
	DeletePermission
)

type SyncMessage interface {
	RequestKind() SyncMsgKind
	Request() (syncer.Request, error)
}

type SyncMessageSerializer interface {
	Serialize(msg SyncMessage) ([]byte, error)
	Deserialize(bytes []byte) (SyncMessage, error)
}
