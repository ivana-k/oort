package domain

import (
	"log"
	"sort"
)

type PermissionKind int

const (
	PermissionKindAllow PermissionKind = iota
	PermissionKindDeny
)

type EvalResult int

const (
	EvalResultAllowed EvalResult = iota
	EvalResultDenied
	EvalResultNonEvaluative
)

const DefaultEvalResult = EvalResultDenied

type PermissionEvalRequest struct {
	Subject []Attribute
	Object  []Attribute
	Env     []Attribute
}

type Permission struct {
	name      string
	kind      PermissionKind
	condition Condition
}

func NewPermission(name string, kind PermissionKind, condition Condition) (*Permission, error) {
	return &Permission{
		name:      name,
		kind:      kind,
		condition: condition,
	}, nil
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
	log.Println("perm eval")
	if !p.condition.Eval(req.Subject, req.Object, req.Env) {
		return EvalResultNonEvaluative
	}
	if p.kind == PermissionKindAllow {
		return EvalResultAllowed
	}
	if p.kind == PermissionKindDeny {
		return EvalResultDenied
	}
	return EvalResultNonEvaluative
}

type PermissionLevel []Permission

func (level PermissionLevel) eval(req PermissionEvalRequest) EvalResult {
	res := EvalResultNonEvaluative
	for _, permission := range level {
		curr := permission.eval(req)
		if curr == EvalResultDenied {
			return EvalResultDenied
		}
		if curr != EvalResultNonEvaluative {
			res = curr
		}
	}
	return res
}

type PermissionPriority int
type PermissionObjHierarchy map[PermissionPriority]PermissionLevel

func (hierarchy PermissionObjHierarchy) eval(req PermissionEvalRequest) EvalResult {
	for _, level := range hierarchy.sortByPriorityDesc() {
		if res := level.eval(req); res != EvalResultNonEvaluative {
			return res
		}
	}
	return DefaultEvalResult
}

func (hierarchy PermissionObjHierarchy) sortByPriorityDesc() []PermissionLevel {
	keys := make([]PermissionPriority, 0, len(hierarchy))
	for k := range hierarchy {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})
	levels := make([]PermissionLevel, 0, len(hierarchy))
	for _, key := range keys {
		levels = append(levels, hierarchy[key])
	}
	return levels
}

type PermissionHierarchy map[PermissionPriority]PermissionObjHierarchy

func (hierarchy PermissionHierarchy) Eval(req PermissionEvalRequest) EvalResult {
	for _, objHierarchy := range hierarchy.sortByPriorityDesc() {
		if res := objHierarchy.eval(req); res != EvalResultNonEvaluative {
			return res
		}
	}
	return DefaultEvalResult
}

func (hierarchy PermissionHierarchy) sortByPriorityDesc() []PermissionObjHierarchy {
	keys := make([]PermissionPriority, 0, len(hierarchy))
	for k := range hierarchy {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})
	levels := make([]PermissionObjHierarchy, 0, len(hierarchy))
	for _, key := range keys {
		levels = append(levels, hierarchy[key])
	}
	return levels
}
