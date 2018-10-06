package lexer

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want *Lexer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newToken(t *testing.T) {
	type args struct {
		tokenName TokenName
		ch        byte
	}
	tests := []struct {
		name string
		args args
		want Token
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newToken(tt.args.tokenName, tt.args.ch); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_readChar(t *testing.T) {
	type fields struct {
		input        string
		position     int
		readPosition int
		ch           byte
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lexer{
				input:        tt.fields.input,
				position:     tt.fields.position,
				readPosition: tt.fields.readPosition,
				ch:           tt.fields.ch,
			}
			l.readChar()
		})
	}
}

func TestLexer_peekChar(t *testing.T) {
	type fields struct {
		input        string
		position     int
		readPosition int
		ch           byte
	}
	tests := []struct {
		name   string
		fields fields
		want   byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lexer{
				input:        tt.fields.input,
				position:     tt.fields.position,
				readPosition: tt.fields.readPosition,
				ch:           tt.fields.ch,
			}
			if got := l.peekChar(); got != tt.want {
				t.Errorf("Lexer.peekChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_delBlank(t *testing.T) {
	type fields struct {
		input        string
		position     int
		readPosition int
		ch           byte
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lexer{
				input:        tt.fields.input,
				position:     tt.fields.position,
				readPosition: tt.fields.readPosition,
				ch:           tt.fields.ch,
			}
			l.delBlank()
		})
	}
}

func TestLexer_readIdentifier(t *testing.T) {
	type fields struct {
		input        string
		position     int
		readPosition int
		ch           byte
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lexer{
				input:        tt.fields.input,
				position:     tt.fields.position,
				readPosition: tt.fields.readPosition,
				ch:           tt.fields.ch,
			}
			if got := l.readIdentifier(); got != tt.want {
				t.Errorf("Lexer.readIdentifier() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_readNumber(t *testing.T) {
	type fields struct {
		input        string
		position     int
		readPosition int
		ch           byte
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lexer{
				input:        tt.fields.input,
				position:     tt.fields.position,
				readPosition: tt.fields.readPosition,
				ch:           tt.fields.ch,
			}
			if got := l.readNumber(); got != tt.want {
				t.Errorf("Lexer.readNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_readText(t *testing.T) {
	type fields struct {
		input        string
		position     int
		readPosition int
		ch           byte
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lexer{
				input:        tt.fields.input,
				position:     tt.fields.position,
				readPosition: tt.fields.readPosition,
				ch:           tt.fields.ch,
			}
			if got := l.readText(); got != tt.want {
				t.Errorf("Lexer.readText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_proveDigit(t *testing.T) {
	type args struct {
		ch byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := proveDigit(tt.args.ch); got != tt.want {
				t.Errorf("proveDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_proveLetter(t *testing.T) {
	type args struct {
		ch byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := proveLetter(tt.args.ch); got != tt.want {
				t.Errorf("proveLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLookupIdent(t *testing.T) {
	type args struct {
		ident string
	}
	tests := []struct {
		name string
		args args
		want TokenName
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LookupIdent(tt.args.ident); got != tt.want {
				t.Errorf("LookupIdent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_NextToken(t *testing.T) {
	type fields struct {
		input        string
		position     int
		readPosition int
		ch           byte
	}
	tests := []struct {
		name   string
		fields fields
		want   Token
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lexer{
				input:        tt.fields.input,
				position:     tt.fields.position,
				readPosition: tt.fields.readPosition,
				ch:           tt.fields.ch,
			}
			if got := l.NextToken(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Lexer.NextToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
