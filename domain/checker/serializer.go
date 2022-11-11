package checker

import (
	"github.com/c12s/oort/domain/model"
)

type AttributeSerializer interface {
	Serialize(attributes []model.Attribute) ([]byte, error)
	Deserialize(bytes []byte) ([]model.Attribute, error)
}

type CheckPermissionResponseSerializer interface {
	Serialize(resp CheckPermissionResp) ([]byte, error)
	Deserialize(bytes []byte) (CheckPermissionResp, error)
}
