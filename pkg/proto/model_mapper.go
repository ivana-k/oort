package proto

import (
	"errors"
	"github.com/c12s/oort/internal/domain"
	"github.com/golang/protobuf/proto"
)

func (x *AttributeId) ToDomain() (*domain.AttributeId, error) {
	return domain.NewAttributeId(x.Name)
}

func (x *AttributeId) FromDomain(id domain.AttributeId) (*AttributeId, error) {
	return &AttributeId{
		Name: id.Name(),
	}, nil
}

func (x *Attribute) ToDomain() (*domain.Attribute, error) {
	value, err := x.valueToDomain()
	if err != nil {
		return nil, err
	}
	id, err := x.Id.ToDomain()
	if err != nil {
		return nil, err
	}
	return domain.NewAttribute(*id, domain.AttributeKind(x.Kind), value)
}

func (x *Attribute) FromDomain(attr domain.Attribute) (*Attribute, error) {
	value, kind, err := x.valueKindFromDomain(attr)
	if err != nil {
		return nil, err
	}
	return &Attribute{
		Id: &AttributeId{
			Name: attr.Name(),
		},
		Kind:  kind,
		Value: value,
	}, nil
}

func (x *AttributeList) ToDomain() ([]domain.Attribute, error) {
	attributes := make([]domain.Attribute, 0)
	for _, attribute := range x.Attributes {
		domainAttribute, err := attribute.ToDomain()
		if err != nil {
			return nil, err
		}
		attributes = append(attributes, *domainAttribute)
	}
	return attributes, nil
}

func (x *Attribute) valueToDomain() (interface{}, error) {
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

func (x *Attribute) valueKindFromDomain(attr domain.Attribute) ([]byte, Attribute_AttributeKind, error) {
	switch attr.Kind() {
	case domain.Int64:
		value := &Int64Attribute{Value: attr.Value().(int64)}
		marshalled, err := proto.Marshal(value)
		return marshalled, Attribute_INT64, err
	case domain.Float64:
		value := &Float64Attribute{Value: attr.Value().(float64)}
		marshalled, err := proto.Marshal(value)
		return marshalled, Attribute_FLOAT64, err
	case domain.String:
		value := &StringAttribute{Value: attr.Value().(string)}
		marshalled, err := proto.Marshal(value)
		return marshalled, Attribute_STRING, err
	case domain.Bool:
		value := &BoolAttribute{Value: attr.Value().(bool)}
		marshalled, err := proto.Marshal(value)
		return marshalled, Attribute_BOOL, err
	default:
		return nil, Attribute_INT64, errors.New("unknown kind")
	}
}

func (x *Resource) ToDomain() (*domain.Resource, error) {
	return domain.NewResource(x.Id, x.Kind)
}

func (x *Resource) FromDomain(res domain.Resource) (*Resource, error) {
	return &Resource{
		Id:   res.Id(),
		Kind: res.Kind(),
	}, nil
}

func (x *Permission) ToDomain() (*domain.Permission, error) {
	condition, err := domain.NewCondition(x.Condition.Expression)
	if err != nil {
		return nil, err
	}
	return domain.NewPermission(x.Name,
		domain.PermissionKind(x.Kind),
		*condition)
}

func (x *Permission) FromDomain(perm domain.Permission) (*Permission, error) {
	return &Permission{
		Name: perm.Name(),
		Kind: Permission_PermissionKind(perm.Kind()),
		Condition: &Condition{
			Expression: perm.Condition().Expression(),
		},
	}, nil
}
