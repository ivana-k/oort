package model

import (
	"errors"
	"fmt"
	"github.com/Knetic/govaluate"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"
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

func validate(expression string) error {
	expr, err := parser.ParseExpr(expression)
	if err != nil {
		return err
	}
	ast.Inspect(expr, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.Ident:
			if !strings.HasPrefix(x.Name, "resource_") && !strings.HasPrefix(x.Name, "principal_") && !strings.HasPrefix(x.Name, "env_") {
				err = errors.New(fmt.Sprintf("invalid variable name %s", x.Name))
			}
		case *ast.BasicLit:
		case *ast.BinaryExpr:
			operation := x.Op
			supported := false
			for _, sop := range getSupportedOperations() {
				if sop == operation {
					supported = true
					break
				}
			}
			if !supported {
				err = errors.New(fmt.Sprintf("operation %s not supported", operation))
			}
		default:
			err = errors.New("expression nodes must be literals, variable names or supported operations")
		}
		return err == nil
	})
	return err
}

func getSupportedOperations() []token.Token {
	return []token.Token{token.ADD, token.SUB, token.MUL, token.QUO, token.REM, token.LAND,
		token.LOR, token.EQL, token.LSS, token.GTR, token.NEQ, token.LEQ, token.GEQ}
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
		parameters["env_"+key] = value
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
