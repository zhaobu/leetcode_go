package main

import "testing"

func Test_findKthNumber(t *testing.T) {
	type args struct {
		m int
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				m: 45,
				n: 12,
				k: 471,
			},
			want: 312,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthNumber(tt.args.m, tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("findKthNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
