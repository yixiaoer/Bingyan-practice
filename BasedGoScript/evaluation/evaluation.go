package evaluation

import (
	"fmt"
	"project/BasedGoScript/ast"
	//"project/BasedGoScript/lexer"
	"project/BasedGoScript/object"
	//"project/BasedGoScript/parser"
)

var (
	NULL = &object.Null{}

	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

func proveError(obj object.Object) bool { //然error出现在合适的位置
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}
	return false
}

func proveTruthy(obj object.Object) bool {
	switch obj {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}

func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	}

	if builtin, ok := builtins[node.Value]; ok {
		return builtin
	}

	return newError("identifier not found: " + node.Value)
}

func evalProgram(program *ast.Program, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range program.Statements {
		result = Eval(statement, env)

		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}
	return result
}

func nativeBoolToBooleanObject(input bool) *object.Boolean { //true与false都是bool object
	if input {
		return TRUE
	} else {
		return FALSE
	}
}

func evalExpressions(exps []ast.Expression, env *object.Environment) []object.Object {
	var result []object.Object

	for _, e := range exps {
		evaluated := Eval(e, env)
		if proveError(evaluated) {
			return []object.Object{evaluated}
		}
		result = append(result, evaluated)
	}

	return result
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

func evalIfExpression(ie *ast.IfExpression, env *object.Environment) object.Object {
	condition := Eval(ie.Condition, env)
	if proveError(condition) {
		return condition
	}
	if proveTruthy(condition) {
		return Eval(ie.Consequence, env)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative, env)
	} else {
		return NULL
	}
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		return newError("unknown operator: %s%s", operator, right.Type())
	}
}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.NUMBER_OBJ {
		return newError("unknown operator: -%s", right.Type())
	}

	value := right.(*object.Number).Value
	return &object.Number{Value: -value}
}

func evalInfixExpression(operator string, left, right object.Object) object.Object {
	switch {
	case left.Type() == object.NUMBER_OBJ && right.Type() == object.NUMBER_OBJ:
		return evalNumberInfixExpression(operator, left, right)
	case operator == "==":
		return nativeBoolToBooleanObject(left == right)
	case operator == "!=":
		return nativeBoolToBooleanObject(left != right)
	case left.Type() == object.TEXT_OBJ && right.Type() == object.TEXT_OBJ:
		return evalTextInfixExpression(operator, left, right)
	case left.Type() != right.Type():
		return newError("type mismatch: %s %s %s", left.Type(), operator, right.Type())
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalNumberInfixExpression(operator string, left, right object.Object) object.Object { //实在找不出来为什么QAQ OTZ
	leftVal := left.(*object.Number).Value //为什么1+1不能显示2，1 +1可以，左边与运算符之间要有间隔才可以，文本类型的相加也不需要
	rightVal := right.(*object.Number).Value

	switch operator {
	case "+":
		return &object.Number{Value: leftVal + rightVal}
	case "-":
		return &object.Number{Value: leftVal - rightVal}
	case "*":
		return &object.Number{Value: leftVal * rightVal}
	case "/":
		return &object.Number{Value: leftVal / rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalTextInfixExpression(operator string, left, right object.Object) object.Object {
	if operator != "+" {
		return newError("unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}

	leftVal := left.(*object.Text).Value
	rightVal := right.(*object.Text).Value
	return &object.Text{Value: leftVal + rightVal}
}

func evalBlockStatement(block *ast.BlockStatement, env *object.Environment) object.Object { //在嵌套时能够实现
	var result object.Object

	for _, statement := range block.Statements {
		result = Eval(statement, env)

		if result != nil {
			rt := result.Type()
			if rt == object.RETURN_VALUE_OBJ || rt == object.ERROR_OBJ {
				return result
			}
		}
	}
	return result
}

func applyFunction(fn object.Object, args []object.Object) object.Object {
	switch fn := fn.(type) {

	case *object.Function:
		extendedEnv := extendFunctionEnv(fn, args)
		evaluated := Eval(fn.Body, extendedEnv)
		return unwrapReturnValue(evaluated)

	case *object.Builtin:
		return fn.Fn(args...)

	default:
		return newError("not a function: %s", fn.Type())
	}
}

func extendFunctionEnv(fn *object.Function, args []object.Object) *object.Environment {
	env := object.NewEnclosedEnvironment(fn.Env)

	for paramIdx, param := range fn.Parameters {
		env.Set(param.Value, args[paramIdx])
	}

	return env
}

func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}
	return obj
}

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.BlockStatement:
		return evalBlockStatement(node, env)
	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	case *ast.CallExpression:
		function := Eval(node.Function, env)
		if proveError(function) { //无处不在的返回errorQAQ
			return function
		}
		args := evalExpressions(node.Arguments, env)
		if len(args) == 1 && proveError(args[0]) {
			return args[0]
		}
		return applyFunction(function, args)
	case *ast.DefStatement:
		val := Eval(node.Value, env)
		if proveError(val) {
			return val
		}
		env.Set(node.Name.Value, val)
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case *ast.FuncLiteral:
		params := node.Parameters
		body := node.Body
		return &object.Function{Parameters: params, Env: env, Body: body}
	case *ast.Identifier: //得到这些values
		return evalIdentifier(node, env)
	case *ast.IfExpression:
		return evalIfExpression(node, env)
	case *ast.InfixExpression:
		left := Eval(node.Left, env)
		if proveError(left) {
			return left
		}
		right := Eval(node.Right, env)
		if proveError(right) {
			return right
		}
		return evalInfixExpression(node.Operator, left, right)
	case *ast.NumberLiteral:
		return &object.Number{Value: node.Value}
	case *ast.PrefixExpression:
		right := Eval(node.Right, env)
		if proveError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)
	case *ast.Program:
		return evalProgram(node, env)
	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue, env)
		if proveError(val) {
			return val
		}
		return &object.ReturnValue{Value: val}
	case *ast.TextLiteral:
		return &object.Text{Value: node.Value}
	}
	return nil
}
