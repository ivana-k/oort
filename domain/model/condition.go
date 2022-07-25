package model

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"log"
)

type Condition struct {
	expression string
}

func NewCondition(expression string) Condition {
	return Condition{
		expression: expression,
	}
}

func (c Condition) Expression() string {
	return c.expression
}

func (c Condition) Eval(principal, resource []Attribute, env map[string]interface{}) bool {
	if c.expression == "" {
		return true
	}

	goeExpr, err := govaluate.NewEvaluableExpression(c.expression)
	if err != nil {
		log.Print(err)
		return false
	}

	parameters := make(map[string]interface{}, 8)
	for _, attr := range principal {
		parameters["principal_"+attr.Name()] = attr.Value()
	}
	for _, attr := range resource {
		parameters["resource_"+attr.Name()] = attr.Value()
	}
	for key, value := range env {
		parameters[key] = value
	}
	fmt.Println("PARAMS ", parameters)

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
