package model

type resourceId struct {
	id   string
	kind string
}

type Resource struct {
	id resourceId
}

func NewResource(id, kind string) Resource {
	return Resource{
		id: resourceId{
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
