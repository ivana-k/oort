package async

import "github.com/c12s/oort/domain/model/syncer"

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
	MsgId() string
	MsgKind() SyncMsgKind
	Request() (syncer.Request, error)
}

type SyncMessageSerializer interface {
	Serialize(msg SyncMessage) ([]byte, error)
	Deserialize(bytes []byte) (SyncMessage, error)
}
