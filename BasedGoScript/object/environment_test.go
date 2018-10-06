package object

import (
	"reflect"
	"testing"
)

func TestEnvironment_Get(t *testing.T) {
	type fields struct {
		store map[string]Object
		outer *Environment
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Object
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Environment{
				store: tt.fields.store,
				outer: tt.fields.outer,
			}
			got, got1 := e.Get(tt.args.name)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Environment.Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Environment.Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestEnvironment_Set(t *testing.T) {
	type fields struct {
		store map[string]Object
		outer *Environment
	}
	type args struct {
		name string
		val  Object
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Object
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Environment{
				store: tt.fields.store,
				outer: tt.fields.outer,
			}
			if got := e.Set(tt.args.name, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Environment.Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewEnvironment(t *testing.T) {
	tests := []struct {
		name string
		want *Environment
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEnvironment(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEnvironment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewEnclosedEnvironment(t *testing.T) {
	type args struct {
		outer *Environment
	}
	tests := []struct {
		name string
		args args
		want *Environment
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEnclosedEnvironment(tt.args.outer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEnclosedEnvironment() = %v, want %v", got, tt.want)
			}
		})
	}
}
