package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type permissionEvalTestCase struct {
	permission  Permission
	evalRequest PermissionEvalRequest
	result      EvalResult
	description string
}

var permissionEvalTestCases = []permissionEvalTestCase{
	{
		permission: NewPermission("cluster.list", PermissionKindAllow, Condition{""}),
		evalRequest: PermissionEvalRequest{
			Resource:  []Attribute{},
			Principal: []Attribute{},
			Env:       []Attribute{},
		},
		result:      EvalResultAllowed,
		description: "allow permission without a condition",
	},
	{
		permission: NewPermission("cluster.list", PermissionKindDeny, Condition{""}),
		evalRequest: PermissionEvalRequest{
			Resource:  []Attribute{},
			Principal: []Attribute{},
			Env:       []Attribute{},
		},
		result:      EvalResultDenied,
		description: "allow permission without a condition",
	},
	{
		permission: NewPermission("cluster.list", PermissionKindAllow, Condition{"resource_id == 1"}),
		evalRequest: PermissionEvalRequest{
			Resource:  []Attribute{NewAttribute(NewAttributeId("id", Int64), 1)},
			Principal: []Attribute{},
			Env:       []Attribute{},
		},
		result:      EvalResultAllowed,
		description: "allow permission with a condition that is met",
	},
	{
		permission: NewPermission("cluster.list", PermissionKindAllow, Condition{"resource_id == 1"}),
		evalRequest: PermissionEvalRequest{
			Resource:  []Attribute{NewAttribute(NewAttributeId("id", Int64), 2)},
			Principal: []Attribute{},
			Env:       []Attribute{},
		},
		result:      EvalResultNonEvaluative,
		description: "allow permission with a condition that is not met",
	},
}

func TestPermissionEval(t *testing.T) {
	for _, testCase := range permissionEvalTestCases {
		c := testCase
		t.Run(c.description, func(t *testing.T) {
			t.Parallel()
			result := c.permission.eval(c.evalRequest)
			assert.Equal(t, result, c.result)
		})
	}
}

type permissionLevelEvalTestCase struct {
	level       PermissionLevel
	evalRequest PermissionEvalRequest
	result      EvalResult
	description string
}

var permissionLevelEvalTestCases = []permissionLevelEvalTestCase{
	{
		level: []Permission{},
		evalRequest: PermissionEvalRequest{
			Resource:  []Attribute{},
			Principal: []Attribute{},
			Env:       []Attribute{},
		},
		result:      EvalResultNonEvaluative,
		description: "empty permission list cannot cannot produce a definite eval result",
	},
	{
		level: []Permission{
			NewPermission("cluster.list", PermissionKindAllow, Condition{""}),
			NewPermission("cluster.list", PermissionKindAllow, Condition{""}),
		},
		evalRequest: PermissionEvalRequest{
			Resource:  []Attribute{},
			Principal: []Attribute{},
			Env:       []Attribute{},
		},
		result:      EvalResultAllowed,
		description: "only allow permissions in a level",
	},
	{
		level: []Permission{
			NewPermission("cluster.list", PermissionKindDeny, Condition{""}),
			NewPermission("cluster.list", PermissionKindDeny, Condition{""}),
		},
		evalRequest: PermissionEvalRequest{
			Resource:  []Attribute{},
			Principal: []Attribute{},
			Env:       []Attribute{},
		},
		result:      EvalResultDenied,
		description: "only deny permissions in a level",
	},
	{
		level: []Permission{
			NewPermission("cluster.list", PermissionKindDeny, Condition{"resource_id == 1"}),
			NewPermission("cluster.list", PermissionKindDeny, Condition{"resource_id == 2"}),
		},
		evalRequest: PermissionEvalRequest{
			Resource:  []Attribute{NewAttribute(NewAttributeId("id", Int64), 3)},
			Principal: []Attribute{},
			Env:       []Attribute{},
		},
		result:      EvalResultNonEvaluative,
		description: "only non-evaluative permissions in a level",
	},
	{
		level: []Permission{
			NewPermission("cluster.list", PermissionKindDeny, Condition{"resource_id == 1"}),
			NewPermission("cluster.list", PermissionKindAllow, Condition{"principal_id == 2"}),
		},
		evalRequest: PermissionEvalRequest{
			Resource:  []Attribute{NewAttribute(NewAttributeId("id", Int64), 1)},
			Principal: []Attribute{NewAttribute(NewAttributeId("id", Int64), 2)},
			Env:       []Attribute{},
		},
		result:      EvalResultDenied,
		description: "deny permission kind wins",
	},
}

func TestPermissionLevelEval(t *testing.T) {
	for _, testCase := range permissionLevelEvalTestCases {
		c := testCase
		t.Run(c.description, func(t *testing.T) {
			t.Parallel()
			result := c.level.eval(c.evalRequest)
			assert.Equal(t, result, c.result)
		})
	}
}

type permissionHierarchyEvalTestCase struct {
	hierarchy   PermissionHierarchy
	evalRequest PermissionEvalRequest
	result      EvalResult
	description string
}

var permissionHierarchyEvalTestCases = []permissionHierarchyEvalTestCase{
	{
		hierarchy: []PermissionLevel{},
		evalRequest: PermissionEvalRequest{
			Resource:  []Attribute{},
			Principal: []Attribute{},
			Env:       []Attribute{},
		},
		result:      DefaultEvalResult,
		description: "empty permission hierarchy should fall back to the default eval result (deny)",
	},
	{
		hierarchy: []PermissionLevel{
			//level 0 - non-evaluative
			{
				NewPermission("cluster.list", PermissionKindDeny, Condition{"resource_id == 1"}),
			},
			//level 1 - non-evaluative
			{
				NewPermission("cluster.list", PermissionKindDeny, Condition{"resource_id == 1"}),
			},
			//level 2 - non-evaluative
			{
				NewPermission("cluster.list", PermissionKindAllow, Condition{"resource_id == 1"}),
			},
		},
		evalRequest: PermissionEvalRequest{
			Resource:  []Attribute{},
			Principal: []Attribute{},
			Env:       []Attribute{},
		},
		result:      DefaultEvalResult,
		description: "hierarchy with no evaluative permission levels should fall back to the default eval result",
	},
	{
		hierarchy: []PermissionLevel{
			//level 0 - non-evaluative
			{
				NewPermission("cluster.list", PermissionKindDeny, Condition{"resource_id == 1"}),
			},
			//level 1 - deny
			{
				NewPermission("cluster.list", PermissionKindDeny, Condition{""}),
			},
			//level 2 - allow
			{
				NewPermission("cluster.list", PermissionKindAllow, Condition{""}),
			},
		},
		evalRequest: PermissionEvalRequest{
			Resource:  []Attribute{},
			Principal: []Attribute{},
			Env:       []Attribute{},
		},
		result:      EvalResultDenied,
		description: "the first evaluative level should give a final eval result - deny",
	},
	{
		hierarchy: []PermissionLevel{
			//level 0 - non-evaluative
			{
				NewPermission("cluster.list", PermissionKindDeny, Condition{"resource_id == 1"}),
			},
			//level 1 - allow
			{
				NewPermission("cluster.list", PermissionKindAllow, Condition{""}),
			},
			//level 2 - deny
			{
				NewPermission("cluster.list", PermissionKindDeny, Condition{""}),
			},
		},
		evalRequest: PermissionEvalRequest{
			Resource:  []Attribute{},
			Principal: []Attribute{},
			Env:       []Attribute{},
		},
		result:      EvalResultAllowed,
		description: "the first evaluative level should give a final eval result - allow",
	},
}

func TestPermissionHierarchyEval(t *testing.T) {
	for _, testCase := range permissionHierarchyEvalTestCases {
		c := testCase
		t.Run(c.description, func(t *testing.T) {
			t.Parallel()
			result := c.hierarchy.Eval(c.evalRequest)
			assert.Equal(t, result, c.result)
		})
	}
}
