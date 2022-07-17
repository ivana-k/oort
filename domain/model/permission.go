package model

type PermissionKind int

const (
	Allow PermissionKind = iota
	Deny
)

type Permission struct {
	name      string
	kind      PermissionKind
	condition Condition
}

type Condition struct {
	expression string
}
