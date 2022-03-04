package main

import (
	"testing"
)

func Test_waysToChange(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{n: 10},
			want: 4,
		},
		{
			name: "2",
			args: args{n: 5},
			want: 2,
		},
		{
			name: "3",
			args: args{n: 900000},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := waysToChange(tt.args.n); got != tt.want {
				t.Errorf("waysToChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
