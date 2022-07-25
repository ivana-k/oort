package model

type AttributeKind int64

const (
	Int64 AttributeKind = iota
	Float64
	String
	Bool
)

type AttributeId struct {
	name string
	kind AttributeKind
}

func NewAttributeId(name string, kind AttributeKind) AttributeId {
	return AttributeId{
		name: name,
		kind: kind,
	}
}

func (attr AttributeId) Name() string {
	return attr.name
}

func (attr AttributeId) Kind() AttributeKind {
	return attr.kind
}

type Attribute struct {
	id    AttributeId
	value interface{}
}

func NewAttribute(id AttributeId, value interface{}) Attribute {
	return Attribute{
		id:    id,
		value: value,
	}
}

func (attr Attribute) Name() string {
	return attr.id.name
}

func (attr Attribute) Kind() AttributeKind {
	return attr.id.kind
}

func (attr Attribute) Value() interface{} {
	return attr.value
}
