package ast

import (
	"bytes"
	"project/BasedGoScript/lexer"
	"strings"
)

type (
	Node interface { //每个node接口都要提供TokenLiteral方法
		TokenLiteral() string
		String() string
	}

	Expression interface { //表达式
		Node
		expressionNode()
	}

	Identifier struct {
		Token lexer.Token // the lexer.Token=IDENT
		Value string
	}

	Program struct { //AST的根节点
		Statements []Statement
	}

	Statement interface { //语句
		Node
		statementNode()
	}

	Boolean struct {
		Token lexer.Token
		Value bool
	}

	FuncLiteral struct {
		Token      lexer.Token // lexer.Token=FUNC
		Parameters []*Identifier
		Body       *BlockStatement
	}

	NumberLiteral struct {
		Token lexer.Token //the lexer.Token=NUMBER
		Value int64
	}

	TextLiteral struct {
		Token lexer.Token
		Value string
	}

	CallExpression struct { //调用函数用到"()"
		Token     lexer.Token // lexer.Token="("
		Function  Expression  // Identifier或者FuncLiteral
		Arguments []Expression
	}

	IfExpression struct {
		Token       lexer.Token // lexer.Token=IF
		Condition   Expression
		Consequence *BlockStatement
		Alternative *BlockStatement
	}

	InfixExpression struct {
		Token    lexer.Token //置于中间的token(infix token)(eg.*)
		Left     Expression  //左边的表达式
		Operator string
		Right    Expression //右边的表达式
	}

	PrefixExpression struct {
		Token    lexer.Token //前置的token(prefix token)(eg.!)
		Operator string
		Right    Expression //包含了这个操作符右边的表达式
	}

	BlockStatement struct {
		Token      lexer.Token // the { token
		Statements []Statement
	}

	DefStatement struct {
		Token lexer.Token //lexer.Token=DEF
		Name  *Identifier //变量名或函数名
		Value Expression  //值
	}

	ExpressionStatement struct { //表达式可以作为一个单独的完整语句
		Token      lexer.Token // 表达式的第一个token
		Expression Expression
	}

	ReturnStatement struct {
		Token       lexer.Token // the lexer.token=RETURN
		ReturnValue Expression
	}
)

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String()) //WriteString方法将一个字符串放到缓冲器的尾部
	}

	return out.String()
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }

func (fl *FuncLiteral) expressionNode()      {}
func (fl *FuncLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FuncLiteral) String() string {

	var out bytes.Buffer
	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}

func (nl *NumberLiteral) expressionNode()      {}
func (nl *NumberLiteral) TokenLiteral() string { return nl.Token.Literal }
func (nl *NumberLiteral) String() string       { return nl.Token.Literal }

func (tl *TextLiteral) expressionNode()      {}
func (tl *TextLiteral) TokenLiteral() string { return tl.Token.Literal }
func (tl *TextLiteral) String() string       { return tl.Token.Literal }

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {

	var out bytes.Buffer
	out.WriteString("If")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("Else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {

	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {

	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

func (ds *DefStatement) statementNode()       {}
func (ds *DefStatement) TokenLiteral() string { return ds.Token.Literal }
func (ds *DefStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ds.TokenLiteral() + " ")
	out.WriteString(ds.Name.String())
	out.WriteString(" = ")

	if ds.Value != nil {
		out.WriteString(ds.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {

	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}
