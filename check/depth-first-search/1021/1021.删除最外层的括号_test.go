package main

import "testing"

func Test_removeOuterParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{s: "(()())(())"},
			want: "()()()",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeOuterParentheses(tt.args.s); got != tt.want {
				t.Errorf("removeOuterParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
