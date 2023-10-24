package proto

import (
	"errors"
	"github.com/c12s/oort/internal/domain"
	"github.com/c12s/oort/pkg/api"
	"github.com/golang/protobuf/proto"
)

func AttributeIdToDomain(id *api.AttributeId) (*domain.AttributeId, error) {
	return domain.NewAttributeId(id.Name)
}

func AttributeToDomain(attr *api.Attribute) (*domain.Attribute, error) {
	value, err := AttributeValueToDomain(attr)
	if err != nil {
		return nil, err
	}
	id, err := AttributeIdToDomain(attr.Id)
	if err != nil {
		return nil, err
	}
	return domain.NewAttribute(*id, domain.AttributeKind(attr.Kind), value)
}

func AttributeValueToDomain(attr *api.Attribute) (interface{}, error) {
	switch attr.Kind {
	case api.Attribute_INT64:
		var value api.Int64Attribute
		err := proto.Unmarshal(attr.Value, &value)
		if err != nil {
			return nil, err
		}
		return value.Value, nil
	case api.Attribute_FLOAT64:
		var value api.Float64Attribute
		err := proto.Unmarshal(attr.Value, &value)
		if err != nil {
			return nil, err
		}
		return value.Value, nil
	case api.Attribute_STRING:
		var value api.StringAttribute
		err := proto.Unmarshal(attr.Value, &value)
		if err != nil {
			return nil, err
		}
		return value.Value, nil
	case api.Attribute_BOOL:
		var value api.BoolAttribute
		err := proto.Unmarshal(attr.Value, &value)
		if err != nil {
			return nil, err
		}
		return value.Value, nil
	default:
		return nil, errors.New("unknown kind")
	}
}

func ResourceToDomain(res *api.Resource) (*domain.Resource, error) {
	return domain.NewResource(res.Id, res.Kind)
}

func PermissionToDomain(perm *api.Permission) (*domain.Permission, error) {
	condition, err := domain.NewCondition(perm.Condition.Expression)
	if err != nil {
		return nil, err
	}
	return domain.NewPermission(perm.Name,
		domain.PermissionKind(perm.Kind),
		*condition)
}
