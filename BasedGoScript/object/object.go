package object //并非面向对象

import (
	"bytes"
	"fmt"
	"project/BasedGoScript/ast"
	"strings"
)

const (
	BOOLEAN_OBJ      = "BOOLEAN"
	ERROR_OBJ        = "ERROR" //输出error
	FUNC_OBJ         = "FUNC"
	NUMBER_OBJ       = "NUMBER"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	TEXT_OBJ         = "TEXT"

	BUILTIN_OBJ = "BUILTIN"
)

type (
	ObjectType string

	Object interface {
		Type() ObjectType
		Inspect() string
	}

	Boolean struct {
		Value bool
	}

	Error struct {
		Message string
	}

	Function struct {
		Parameters []*ast.Identifier
		Body       *ast.BlockStatement
		Env        *Environment
	}

	Null struct{}

	Number struct {
		Value int64
	}

	ReturnValue struct {
		Value Object
	}

	Text struct {
		Value string
	}

	Builtin struct {
		Fn BuiltinFunction
	}
	BuiltinFunction func(args ...Object) Object
)

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }

func (f *Function) Type() ObjectType { return FUNC_OBJ }
func (f *Function) Inspect() string {

	var out bytes.Buffer
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

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string {
	return "null"
}

func (n *Number) Type() ObjectType { return NUMBER_OBJ }
func (n *Number) Inspect() string { //在源码中有数字(于此而言即是整数),先转换ast.NumberLiteral
	return fmt.Sprintf("%d", n.Value) //当evaluate那个AST节点时转换成object.Integer
}

func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }

func (t *Text) Type() ObjectType { return TEXT_OBJ }
func (t *Text) Inspect() string  { return t.Value }

func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string  { return "builtin function" }
