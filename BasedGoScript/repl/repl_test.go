package repl

import (
	"bytes"
	"io"
	"testing"
)

func TestStart(t *testing.T) {
	type args struct {
		in io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			Start(tt.args.in, out)
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("Start() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
