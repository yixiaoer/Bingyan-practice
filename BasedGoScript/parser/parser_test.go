package parser

import (
	"project/BasedGoScript/lexer"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		l *lexer.Lexer
	}
	tests := []struct {
		name string
		args args
		want *Parser
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
