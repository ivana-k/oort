package domain

import (
	"errors"
	"github.com/Knetic/govaluate"
	"go/ast"
	"go/parser"
	"go/token"
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
func (c Condition) IsEmpty() bool {
	return c.expression == ""
}
func (c Condition) Eval(sub, obj, env []Attribute) bool {
	if c.IsEmpty() {
		return true
	}

	goeExpr, err := govaluate.NewEvaluableExpression(c.expression)
	if err != nil {
		return false
	}

	parameters := make(map[string]interface{}, 8)
	for _, attr := range sub {
		parameters[SubVarNamePrefix+attr.Name()] = attr.Value()
	}
	for _, attr := range obj {
		parameters[ObjVarNamePrefix+attr.Name()] = attr.Value()
	}
	for _, attr := range env {
		parameters[EnvVarNamePrefix+attr.Name()] = attr.Value()
	}

	result, err := goeExpr.Evaluate(parameters)
	if err != nil {
		return false
	}
	boolResult, ok := result.(bool)
	if !ok {
		return false
	}
	return boolResult
}

const (
	SubVarNamePrefix = "sub_"
	ObjVarNamePrefix = "obj_"
	EnvVarNamePrefix = "env_"
)

var validOperations = []token.Token{
	token.ADD,
	token.SUB,
	token.MUL,
	token.QUO,
	token.REM,
	token.LAND,
	token.LOR,
	token.EQL,
	token.LSS,
	token.GTR,
	token.NEQ,
	token.LEQ,
	token.GEQ,
}

var (
	ErrInvalidOperation    = errors.New("expression operation invalid")
	ErrInvalidVariableName = errors.New("expression variable name invalid")
	ErrInvalidNode         = errors.New("expression nodes must be literals, variable names or supported operations")
	ErrParsing             = errors.New("not an expression")
)

func validate(expression string) error {
	if len(expression) == 0 {
		return nil
	}
	expr, err := parser.ParseExpr(expression)
	if err != nil {
		return ErrParsing
	}
	ast.Inspect(expr, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.BasicLit:
		case *ast.ParenExpr:
		case *ast.Ident:
			if !validVariableNamePrefix(x.Name) {
				err = ErrInvalidVariableName
			}
		case *ast.BinaryExpr:
			if !validOperation(x.Op) {
				err = ErrInvalidOperation
			}
		case nil:
		default:
			err = ErrInvalidNode
		}
		return err == nil
	})
	return err
}

func validVariableNamePrefix(varName string) bool {
	if strings.HasPrefix(varName, SubVarNamePrefix) {
		return true
	}
	if strings.HasPrefix(varName, ObjVarNamePrefix) {
		return true
	}
	if strings.HasPrefix(varName, EnvVarNamePrefix) {
		return true
	}
	return false
}

func validOperation(operation token.Token) bool {
	for _, supportedOperation := range validOperations {
		if operation == supportedOperation {
			return true
		}
	}
	return false
}
