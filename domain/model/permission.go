package model

type Permission struct {
	name      string
	kind      PermissionKind
	condition Condition
}

func NewPermission(name string, kind PermissionKind, condition Condition) Permission {
	return Permission{
		name:      name,
		kind:      kind,
		condition: condition,
	}
}

func (p Permission) Name() string {
	return p.name
}

func (p Permission) Kind() PermissionKind {
	return p.kind
}

type PermissionKind int

const (
	Allow PermissionKind = iota
	Deny
)

type EvalResult = int

const (
	Allowed EvalResult = iota
	Denied
	Unknown
)

const DefaultEvalResult EvalResult = Denied

func (p Permission) Condition() Condition {
	return p.condition
}

func (p Permission) Eval(principal, resource []Attribute, env map[string]interface{}) EvalResult {
	if !p.condition.Eval(principal, resource, env) {
		return Unknown
	}
	if p.kind == Allow {
		return Allowed
	}
	return Denied
}

type PermissionList []Permission

func (level PermissionList) Eval(principal, resource []Attribute, env map[string]interface{}) EvalResult {
	res := Unknown
	for _, permission := range level {
		curr := permission.Eval(principal, resource, env)
		if curr != Unknown {
			res = curr
		}
		if curr == Denied {
			break
		}
	}
	return res
}

type PermissionHierarchy []PermissionList

func (hierarchy PermissionHierarchy) Eval(principal, resource []Attribute, env map[string]interface{}) EvalResult {
	res := Unknown
	for _, level := range hierarchy {
		res = level.Eval(principal, resource, env)
		if res != Unknown {
			break
		}
	}
	if res == Unknown {
		return DefaultEvalResult
	}
	return res
}
