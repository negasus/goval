package main

import "testing"

func Test_isRuleVar(t *testing.T) {
	type args struct {
		rule string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{
			name:  "empty",
			args:  args{rule: ""},
			want:  "",
			want1: false,
		},
		{
			name:  "small size",
			args:  args{rule: "{}"},
			want:  "",
			want1: false,
		},
		{
			name:  "no brace 1",
			args:  args{rule: "a{a}"},
			want:  "",
			want1: false,
		},
		{
			name:  "no brace 2",
			args:  args{rule: "{a}a"},
			want:  "",
			want1: false,
		},
		{
			name:  "ok",
			args:  args{rule: "{foo}"},
			want:  "foo",
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := isRuleVar(tt.args.rule)
			if got != tt.want {
				t.Errorf("isRuleVar() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("isRuleVar() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
