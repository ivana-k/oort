package model

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
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
)

func validate(expression string) error {
	expr, err := parser.ParseExpr(expression)
	if err != nil {
		return err
	}
	ast.Inspect(expr, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.BasicLit:
		case *ast.Ident:
			if !validVariableNamePrefix(x.Name) {
				err = ErrInvalidVariableName
			}
		case *ast.BinaryExpr:
			if !validOperation(x.Op) {
				err = ErrInvalidOperation
			}
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
