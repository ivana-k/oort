package async

import (
	"github.com/c12s/oort/domain/syncer"
)

type SyncMsgKind int

const (
	ConnectResources SyncMsgKind = iota
	DisconnectResources
	UpsertAttribute
	RemoveAttribute
	InsertPermission
	RemovePermission
)

type SyncMessage interface {
	RequestKind() SyncMsgKind
	Request() (syncer.Request, error)
}

type SyncMessageSerializer interface {
	Serialize(msg SyncMessage) ([]byte, error)
	Deserialize(bytes []byte) (SyncMessage, error)
}
