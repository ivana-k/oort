package common

import (
	"errors"
	"github.com/c12s/oort/domain/model"
	"github.com/golang/protobuf/proto"
)

func (x *AttributeId) MapToDomain() model.AttributeId {
	return model.NewAttributeId(x.Name, model.AttributeKind(x.Kind))
}

func (x *Attribute) MapToDomain() (model.Attribute, error) {
	value, err := x.originalValue()
	if err != nil {
		return model.Attribute{}, err
	}
	return model.NewAttribute(x.Id.MapToDomain(), value), nil
}

func (x *Attribute) originalValue() (interface{}, error) {
	switch x.Id.Kind {
	case AttributeId_INT64:
		var value Int64Attribute
		err := proto.Unmarshal(x.Value, &value)
		if err != nil {
			return nil, err
		}
		return value.Value, nil
	case AttributeId_FLOAT64:
		var value Float64Attribute
		err := proto.Unmarshal(x.Value, &value)
		if err != nil {
			return nil, err
		}
		return value.Value, nil
	case AttributeId_STRING:
		var value StringAttribute
		err := proto.Unmarshal(x.Value, &value)
		if err != nil {
			return nil, err
		}
		return value.Value, nil
	case AttributeId_BOOL:
		var value BoolAttribute
		err := proto.Unmarshal(x.Value, &value)
		if err != nil {
			return nil, err
		}
		return value.Value, nil
	default:
		return nil, errors.New("unknown kind")
	}
}

func (x *Resource) MapToDomain() model.Resource {
	return model.NewResource(x.Id, x.Kind)
}

func (x *Permission) MapToDomain() model.Permission {
	return model.NewPermission(x.Name,
		model.PermissionKind(x.Kind),
		model.NewCondition(x.Condition.Expression))
}
