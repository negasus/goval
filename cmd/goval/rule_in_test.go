package main

import (
	"reflect"
	"testing"
)

func Test_parseInput(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "empty",
			args: args{input: ""},
			want: []string{""},
		},
		{
			name: "simple1",
			args: args{input: "foo,bar"},
			want: []string{"foo", "bar"},
		},
		{
			name: "spaces",
			args: args{input: " foo , bar , ,"},
			want: []string{" foo ", " bar ", " ", ""},
		},
		{
			name: "quotes",
			args: args{input: "foo,'bar','baz\\''"},
			want: []string{"foo", "bar", "baz'"},
		},
		{
			name: "comma in quotes",
			args: args{input: "foo,'ba,r','baz\\''"},
			want: []string{"foo", "ba,r", "baz'"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseStringInput(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeString(t *testing.T) {
	type args struct {
		slice []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty",
			args: args{slice: []string{}},
			want: "{}",
		},
		{
			name: "one",
			args: args{slice: []string{"foo"}},
			want: `{"foo"}`,
		},
		{
			name: "many",
			args: args{slice: []string{"foo", "bar", " baz "}},
			want: `{"foo", "bar", " baz "}`,
		},
		{
			name: "escape",
			args: args{slice: []string{`fo " bar`}},
			want: `{"fo \" bar"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeStringFromStrings(tt.args.slice); got != tt.want {
				t.Errorf("makeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
