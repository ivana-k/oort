package model

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

const (
	ResourceVarNamePrefix  = "resource_"
	PrincipalVarNamePrefix = "principal_"
	EnvVarNamePrefix       = "env_"
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
	if strings.HasPrefix(varName, ResourceVarNamePrefix) {
		return true
	}
	if strings.HasPrefix(varName, PrincipalVarNamePrefix) {
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
