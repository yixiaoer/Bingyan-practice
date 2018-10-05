package object //并非面向对象

import (
	"bytes"
	"fmt"
	"project/BasedGoScript/ast"
	"strings"
)

const (
	NUMBER_OBJ       = "NUMBER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR" //输出error
	FUNC_OBJ         = "FUNC"
	TEXT_OBJ         = "TEXT"
)

type (
	ObjectType string

	Object interface {
		Type() ObjectType
		Inspect() string
	}

	Number struct {
		Value int64
	}

	Boolean struct {
		Value bool
	}

	Null struct{}

	ReturnValue struct {
		Value Object
	}

	Error struct {
		Message string
	}

	Function struct {
		Parameters []*ast.Identifier
		Body       *ast.BlockStatement
		Env        *Environment
	}

	Text struct {
		Value string
	}
)

var out bytes.Buffer

func (n *Number) Type() ObjectType { return NUMBER_OBJ }
func (n *Number) Inspect() string { //在源码中有数字(于此而言即是整数),先转换ast.NumberLiteral
	return fmt.Sprintf("%d", n.Value) //当evaluate那个AST节点时转换成object.Integer
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string {
	return "null"
}

func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }

func (f *Function) Type() ObjectType { return FUNC_OBJ }
func (f *Function) Inspect() string {

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out.WriteString("Func")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(f.Body.String())

	return out.String()
}

func (t *Text) Type() ObjectType { return TEXT_OBJ }
func (t *Text) Inspect() string  { return t.Value }
