package lexer

const (
	ILLEGAL = "ILLEGAL" //未知类型
	EOF     = "EOF"     //文本结束，以此告知这一句的语法分析结束

	IDENT  = "IDENT"  //
	NUMBER = "NUMBER" //数字类型，在这个脚本语言中仅表示整型
	TEXT   = "TEXT"   //文本类型，在这个脚本语言中仅支持英文

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT     = "<"
	GT     = ">"
	EQ     = "=="
	NOT_EQ = "!="

	//分隔符
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	//关键字
	FUNC   = "FUNC"
	DEF    = "DEF"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
	IF     = "IF"
	ELSE   = "ELSE"
	RETURN = "RETURN"
)

type TokenName string //更清晰地表示各种token的类别名称

type Token struct { //一个token的基本信息
	Name    TokenName
	Literal string
}

type Lexer struct {
	input        string //当前输入
	position     int    //当前输入的指向当前字符的位置
	readPosition int    //再下一个字符(peek)
	ch           byte   //当前字符
}

var keywords = map[string]TokenName{
	"Func":   FUNC,
	"Def":    DEF,
	"True":   TRUE,
	"False":  FALSE,
	"If":     IF,
	"Else":   ELSE,
	"Return": RETURN,
}

func (l *Lexer) readChar() { //逐一读取字符
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position++
	l.readPosition++
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func newToken(tokenName TokenName, ch byte) Token {
	return Token{Name: tokenName, Literal: string(ch)}
}

func proveLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'

}

func (l *Lexer) reeadIdentifier() string {
	position := l.position
	for proveLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func proveDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for proveDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) delBlank() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func LookupIdent(ident string) TokenName { //判断得到的indentifier是否是关键字，如果是可以直接返回tokenname确定类型
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

func (l *Lexer) peekChar() byte { //只预读一个字符看是否是空即可
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readText() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' {
			break
		}
	}
	return l.input[position:l.position]
}

func (l *Lexer) NextToken() Token { //将读取的字符进行基本词法识别
	var tok Token
	l.delBlank()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Name: EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(PLUS, l.ch)
	case '-':
		tok = newToken(MINUS, l.ch)
	case '/':
		tok = newToken(SLASH, l.ch)
	case '*':
		tok = newToken(ASTERISK, l.ch)
	case '<':
		tok = newToken(LT, l.ch)
	case '>':
		tok = newToken(GT, l.ch)
	case ',':
		tok = newToken(COMMA, l.ch)
	case ';':
		tok = newToken(SEMICOLON, l.ch)
	case '(':
		tok = newToken(LPAREN, l.ch)
	case ')':
		tok = newToken(RPAREN, l.ch)
	case '{':
		tok = newToken(LBRACE, l.ch)
	case '}':
		tok = newToken(RBRACE, l.ch)
	case '"':
		tok.Name = TEXT
		tok.Literal = l.readText()
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Name: NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(BANG, l.ch)
		}
	case 0:
		tok.Literal = ""
		tok.Name = EOF
	default:
		if proveLetter(l.ch) {
			tok.Literal = l.reeadIdentifier()
			tok.Name = LookupIdent(tok.Literal)
			return tok
		} else if proveDigit(l.ch) {
			tok.Name = NUMBER
			tok.Literal = l.readNumber()
		} else {
			tok = newToken(ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}
