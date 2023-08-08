package domain

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
		resource:    []Attribute{NewAttribute(NewAttributeId("id"), Int64, 2)},
		principal:   []Attribute{},
		env:         []Attribute{},
		result:      true,
		description: "resource - condition that is met",
	},
	{
		condition:   Condition{expression: "resource_id == 2"},
		resource:    []Attribute{NewAttribute(NewAttributeId("id"), Int64, 3)},
		principal:   []Attribute{},
		env:         []Attribute{},
		result:      false,
		description: "resource - condition that is not met (incorrect attribute value)",
	},
	{
		condition:   Condition{expression: "resource_id == 2"},
		resource:    []Attribute{NewAttribute(NewAttributeId("id"), String, "hello")},
		principal:   []Attribute{},
		env:         []Attribute{},
		result:      false,
		description: "resource - condition that is not met (incorrect attribute type)",
	},
	{
		condition:   Condition{expression: "resource_id == 2"},
		resource:    []Attribute{NewAttribute(NewAttributeId("ID"), Int64, 3)},
		principal:   []Attribute{},
		env:         []Attribute{},
		result:      false,
		description: "resource - condition that is not met (incorrect attribute name)",
	},
	//principal
	{
		condition:   Condition{expression: "principal_id == 2"},
		resource:    []Attribute{},
		principal:   []Attribute{NewAttribute(NewAttributeId("id"), Int64, 2)},
		env:         []Attribute{},
		result:      true,
		description: "resource - condition that is met",
	},
	{
		condition:   Condition{expression: "principal_id == 2"},
		resource:    []Attribute{},
		principal:   []Attribute{NewAttribute(NewAttributeId("id"), Int64, 3)},
		env:         []Attribute{},
		result:      false,
		description: "resource - condition that is not met (incorrect attribute value)",
	},
	{
		condition:   Condition{expression: "principal_id == 2"},
		resource:    []Attribute{},
		principal:   []Attribute{NewAttribute(NewAttributeId("id"), String, "hello")},
		env:         []Attribute{},
		result:      false,
		description: "resource - condition that is not met (incorrect attribute type)",
	},
	{
		condition:   Condition{expression: "principal_id == 2"},
		resource:    []Attribute{},
		principal:   []Attribute{NewAttribute(NewAttributeId("ID"), Int64, 3)},
		env:         []Attribute{},
		result:      false,
		description: "resource - condition that is not met (incorrect attribute name)",
	},
	//env
	{
		condition:   Condition{expression: "env_id == 2"},
		resource:    []Attribute{},
		principal:   []Attribute{},
		env:         []Attribute{NewAttribute(NewAttributeId("id"), Int64, 2)},
		result:      true,
		description: "env - condition that is met",
	},
	{
		condition:   Condition{expression: "env_id == 2"},
		resource:    []Attribute{},
		principal:   []Attribute{},
		env:         []Attribute{NewAttribute(NewAttributeId("id"), Int64, 3)},
		result:      false,
		description: "env - condition that is not met (incorrect attribute value)",
	},
	{
		condition:   Condition{expression: "env_id == 2"},
		resource:    []Attribute{},
		principal:   []Attribute{},
		env:         []Attribute{NewAttribute(NewAttributeId("id"), String, "hello")},
		result:      false,
		description: "env - condition that is not met (incorrect attribute type)",
	},
	{
		condition:   Condition{expression: "env_id == 2"},
		resource:    []Attribute{},
		principal:   []Attribute{},
		env:         []Attribute{NewAttribute(NewAttributeId("ID"), Int64, 3)},
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

type expressionValidationTestCase struct {
	expression  string
	err         error
	description string
}

var expressionValidationTestCases = []expressionValidationTestCase{
	{
		expression:  "",
		err:         nil,
		description: "empty expression",
	},
	{
		expression:  "resource_id == 1",
		err:         nil,
		description: "valid operands (resource attribute and int literal) and operation (token.EQL)",
	},
	{
		expression:  "principal_id != \"aaa\"",
		err:         nil,
		description: "valid operands (principal attribute and string literal) and operation (token.NEQ)",
	},
	{
		expression:  "env_latitude > 22.15",
		err:         nil,
		description: "valid operands (env attribute and float literal) and operation (token.GTR)",
	},
	{
		expression:  "2 + 2",
		err:         nil,
		description: "cannot check expr result type, so it is valid",
	},
	{
		expression:  "resource_id == 1 || (env_id == 2 && principal_id == 3)",
		err:         nil,
		description: "parenthesized expression",
	},
	{
		expression:  "fmt.Println(\"hello world\")",
		err:         ErrInvalidNode,
		description: "function call",
	},
	{
		expression:  "var c = 1",
		err:         ErrParsing,
		description: "statement instead of an expression",
	},
	{
		expression:  "id == 5",
		err:         ErrInvalidVariableName,
		description: "incorrectly prefixed variable name",
	},
	{
		expression:  "resource_id >> 2",
		err:         ErrInvalidOperation,
		description: "unsupported operator",
	},
}

func TestConditionExpression(t *testing.T) {
	for _, testCase := range expressionValidationTestCases {
		c := testCase
		t.Run(c.description, func(t *testing.T) {
			t.Parallel()
			err := validate(c.expression)
			assert.ErrorIs(t, err, c.err)
		})
	}
}
