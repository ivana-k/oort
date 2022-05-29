package model

type Permission interface {
	GetName() string
}

type permission struct {
	name string
}

func NewPermission(name string) Permission {
	return &permission{
		name: name,
	}
}

func (p *permission) GetName() string {
	return p.name
}
