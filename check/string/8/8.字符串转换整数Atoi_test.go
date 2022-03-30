package main

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				s: "-91283472332",
			},
			want: -2147483648,
		},
		{
			name: "test2",
			args: args{
				s: "+1",
			},
			want: 1,
		},
		{
			name: "test3",
			args: args{
				s: "+-12",
			},
			want: 0,
		},
		{
			name: "test4",
			args: args{
				s: "2147483648",
			},
			want: 2147483647,
		},
		{
			name: "test5",
			args: args{
				s: "42",
			},
			want: 42,
		},
		{
			name: "test6",
			args: args{
				s: "2147483646",
			},
			want: 2147483646,
		},
		{
			name: "test7",
			args: args{
				s: "00000-42a1234",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
