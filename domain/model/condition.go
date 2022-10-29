package model

import (
	"github.com/Knetic/govaluate"
	"log"
)

const (
	ResourceVarNamePrefix  = "resource_"
	PrincipalVarNamePrefix = "principal_"
	EnvVarNamePrefix       = "env_"
)

type Condition struct {
	expression string
}

func NewCondition(expression string) (*Condition, error) {
	if err := validate(expression); err != nil {
		return nil, err
	}
	return &Condition{
		expression: expression,
	}, nil
}

func (c Condition) Expression() string {
	return c.expression
}
func (c Condition) IsEmpty() bool {
	return c.expression == ""
}
func (c Condition) Eval(principal, resource, env []Attribute) bool {
	if c.IsEmpty() {
		return true
	}

	goeExpr, err := govaluate.NewEvaluableExpression(c.expression)
	if err != nil {
		log.Print(err)
		return false
	}

	parameters := make(map[string]interface{}, 8)
	for _, attr := range principal {
		parameters[PrincipalVarNamePrefix+attr.Name()] = attr.Value()
	}
	for _, attr := range resource {
		parameters[ResourceVarNamePrefix+attr.Name()] = attr.Value()
	}
	for _, attr := range env {
		parameters[EnvVarNamePrefix+attr.Name()] = attr.Value()
	}

	result, err := goeExpr.Evaluate(parameters)
	if err != nil {
		log.Print(err)
		return false
	}
	boolResult, ok := result.(bool)
	if !ok {
		return false
	}
	return boolResult
}
