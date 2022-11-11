package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
