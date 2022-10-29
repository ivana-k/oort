package model

type PermissionKind int

const (
	Allow PermissionKind = iota
	Deny
)

type EvalResult int

const (
	Allowed EvalResult = iota
	Denied
	NonEvaluative
)

const DefaultEvalResult = Denied

type PermissionEvalRequest struct {
	Resource  []Attribute
	Principal []Attribute
	Env       []Attribute
}

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

func (p Permission) Condition() Condition {
	return p.condition
}

func (p Permission) eval(req PermissionEvalRequest) EvalResult {
	if !p.condition.Eval(req.Principal, req.Resource, req.Env) {
		return NonEvaluative
	}
	if p.kind == Allow {
		return Allowed
	}
	if p.kind == Deny {
		return Denied
	}
	return NonEvaluative
}

type PermissionLevel []Permission

func (level PermissionLevel) eval(req PermissionEvalRequest) EvalResult {
	res := NonEvaluative
	for _, permission := range level {
		curr := permission.eval(req)
		if curr == Denied {
			return Denied
		}
		if curr != NonEvaluative {
			res = curr
		}
	}
	return res
}

type PermissionHierarchy []PermissionLevel

func (hierarchy PermissionHierarchy) Eval(req PermissionEvalRequest) EvalResult {
	for _, level := range hierarchy {
		if res := level.eval(req); res != NonEvaluative {
			return res
		}
	}
	return DefaultEvalResult
}
