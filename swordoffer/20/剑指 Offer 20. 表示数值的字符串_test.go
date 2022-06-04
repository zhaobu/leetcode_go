package main

import (
	"testing"
)

func Test_isNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test3",
			args: args{
				s: "0e",
			},
			want: false,
		},
		{
			name: "test4",
			args: args{
				s: "e9",
			},
			want: false,
		},
		{
			name: "test5",
			args: args{
				s: "3.",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNumber(tt.args.s); got != tt.want {
				t.Errorf("isNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
