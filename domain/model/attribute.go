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
}

func NewAttributeId(name string) AttributeId {
	return AttributeId{
		name: name,
	}
}

func (attr AttributeId) Name() string {
	return attr.name
}

type Attribute struct {
	id    AttributeId
	kind  AttributeKind
	value interface{}
}

func NewAttribute(id AttributeId, kind AttributeKind, value interface{}) Attribute {
	return Attribute{
		id:    id,
		kind:  kind,
		value: value,
	}
}

func (attr Attribute) Name() string {
	return attr.id.name
}

func (attr Attribute) Kind() AttributeKind {
	return attr.kind
}

func (attr Attribute) Value() interface{} {
	return attr.value
}
