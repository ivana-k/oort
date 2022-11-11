package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type conditionEvalTestCase struct {
	condition   Condition
	resource    []Attribute
	principal   []Attribute
	env         []Attribute
	result      bool
	description string
}

var conditionEvalTestCases = []conditionEvalTestCase{
	{
		condition:   Condition{expression: ""},
		resource:    []Attribute{},
		principal:   []Attribute{},
		env:         []Attribute{},
		result:      true,
		description: "empty expression",
	},
	//resource
	{
		condition:   Condition{expression: "resource_id == 2"},
		resource:    []Attribute{NewAttribute(NewAttributeId("id", Int64), 2)},
		principal:   []Attribute{},
		env:         []Attribute{},
		result:      true,
		description: "resource - condition that is met",
	},
	{
		condition:   Condition{expression: "resource_id == 2"},
		resource:    []Attribute{NewAttribute(NewAttributeId("id", Int64), 3)},
		principal:   []Attribute{},
		env:         []Attribute{},
		result:      false,
		description: "resource - condition that is not met (incorrect attribute value)",
	},
	{
		condition:   Condition{expression: "resource_id == 2"},
		resource:    []Attribute{NewAttribute(NewAttributeId("id", String), "hello")},
		principal:   []Attribute{},
		env:         []Attribute{},
		result:      false,
		description: "resource - condition that is not met (incorrect attribute type)",
	},
	{
		condition:   Condition{expression: "resource_id == 2"},
		resource:    []Attribute{NewAttribute(NewAttributeId("ID", Int64), 3)},
		principal:   []Attribute{},
		env:         []Attribute{},
		result:      false,
		description: "resource - condition that is not met (incorrect attribute name)",
	},
	//principal
	{
		condition:   Condition{expression: "principal_id == 2"},
		resource:    []Attribute{},
		principal:   []Attribute{NewAttribute(NewAttributeId("id", Int64), 2)},
		env:         []Attribute{},
		result:      true,
		description: "resource - condition that is met",
	},
	{
		condition:   Condition{expression: "principal_id == 2"},
		resource:    []Attribute{},
		principal:   []Attribute{NewAttribute(NewAttributeId("id", Int64), 3)},
		env:         []Attribute{},
		result:      false,
		description: "resource - condition that is not met (incorrect attribute value)",
	},
	{
		condition:   Condition{expression: "principal_id == 2"},
		resource:    []Attribute{},
		principal:   []Attribute{NewAttribute(NewAttributeId("id", String), "hello")},
		env:         []Attribute{},
		result:      false,
		description: "resource - condition that is not met (incorrect attribute type)",
	},
	{
		condition:   Condition{expression: "principal_id == 2"},
		resource:    []Attribute{},
		principal:   []Attribute{NewAttribute(NewAttributeId("ID", Int64), 3)},
		env:         []Attribute{},
		result:      false,
		description: "resource - condition that is not met (incorrect attribute name)",
	},
	//env
	{
		condition:   Condition{expression: "env_id == 2"},
		resource:    []Attribute{},
		principal:   []Attribute{},
		env:         []Attribute{NewAttribute(NewAttributeId("id", Int64), 2)},
		result:      true,
		description: "env - condition that is met",
	},
	{
		condition:   Condition{expression: "env_id == 2"},
		resource:    []Attribute{},
		principal:   []Attribute{},
		env:         []Attribute{NewAttribute(NewAttributeId("id", Int64), 3)},
		result:      false,
		description: "env - condition that is not met (incorrect attribute value)",
	},
	{
		condition:   Condition{expression: "env_id == 2"},
		resource:    []Attribute{},
		principal:   []Attribute{},
		env:         []Attribute{NewAttribute(NewAttributeId("id", String), "hello")},
		result:      false,
		description: "env - condition that is not met (incorrect attribute type)",
	},
	{
		condition:   Condition{expression: "env_id == 2"},
		resource:    []Attribute{},
		principal:   []Attribute{},
		env:         []Attribute{NewAttribute(NewAttributeId("ID", Int64), 3)},
		result:      false,
		description: "env - condition that is not met (incorrect attribute name)",
	},
}

func TestConditionEval(t *testing.T) {
	for _, testCase := range conditionEvalTestCases {
		c := testCase
		t.Run(c.description, func(t *testing.T) {
			t.Parallel()
			result := c.condition.Eval(c.principal, c.resource, c.env)
			assert.Equal(t, result, c.result)
		})
	}
}
