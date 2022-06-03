package main

import "testing"

func Test_consecutiveNumbersSum(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{n: 79932},
			want: 4,
		},
		{
			name: "test2",
			args: args{n: 1000000000},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := consecutiveNumbersSum(tt.args.n); got != tt.want {
				t.Errorf("consecutiveNumbersSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
