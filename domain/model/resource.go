package model

type AttributeKind int

const (
	Int64 AttributeKind = iota
	Float64
	String
	Bool
)

type ConnectionKind int

func (connKind ConnectionKind) String() string {
	switch connKind {
	case Inherits:
		return "Inherits"
	default:
		return "Includes"
	}
}

const (
	Inherits ConnectionKind = iota
	Includes
)

type AttributeId struct {
	name string
	kind AttributeKind
}

type Attribute struct {
	id    AttributeId
	value []byte
}

func NewAttribute(name string, kind AttributeKind, value []byte) Attribute {
	return Attribute{
		id: AttributeId{
			name: name,
			kind: kind,
		},
		value: value,
	}
}

func (attr Attribute) Name() string {
	return attr.id.name
}

func (attr Attribute) Kind() AttributeKind {
	return attr.id.kind
}

func (attr Attribute) Value() []byte {
	return attr.value
}

type ResourceId struct {
	id   string
	kind string
}

type Resource struct {
	id ResourceId
}

func NewResource(id, kind string) Resource {
	return Resource{
		id: ResourceId{
			id:   id,
			kind: kind,
		},
	}
}

func (r Resource) Id() string {
	return r.id.id
}

func (r Resource) Kind() string {
	return r.id.kind
}
