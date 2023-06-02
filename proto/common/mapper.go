package common

import (
	"errors"
	"github.com/c12s/oort/domain/model"
	"github.com/golang/protobuf/proto"
)

func (x *AttributeId) MapToDomain() model.AttributeId {
	return model.NewAttributeId(x.Name)
}

func (x *Attribute) MapToDomain() (model.Attribute, error) {
	value, err := x.originalValue()
	if err != nil {
		return model.Attribute{}, err
	}
	return model.NewAttribute(x.Id.MapToDomain(), model.AttributeKind(x.Kind), value), nil
}

func (x *AttributeList) MapToDomain() ([]model.Attribute, error) {
	attributes := make([]model.Attribute, 0)
	for _, attribute := range x.Attributes {
		domainAttribute, err := attribute.MapToDomain()
		if err != nil {
			return nil, err
		}
		attributes = append(attributes, domainAttribute)
	}
	return attributes, nil
}

func (x *Attribute) originalValue() (interface{}, error) {
	switch x.Kind {
	case Attribute_INT64:
		var value Int64Attribute
		err := proto.Unmarshal(x.Value, &value)
		if err != nil {
			return nil, err
		}
		return value.Value, nil
	case Attribute_FLOAT64:
		var value Float64Attribute
		err := proto.Unmarshal(x.Value, &value)
		if err != nil {
			return nil, err
		}
		return value.Value, nil
	case Attribute_STRING:
		var value StringAttribute
		err := proto.Unmarshal(x.Value, &value)
		if err != nil {
			return nil, err
		}
		return value.Value, nil
	case Attribute_BOOL:
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

func (x *Permission) MapToDomain() (model.Permission, error) {
	condition, err := model.NewCondition(x.Condition.Expression)
	if err != nil {
		return model.Permission{}, err
	}
	return model.NewPermission(x.Name,
		model.PermissionKind(x.Kind),
		*condition), nil
}
