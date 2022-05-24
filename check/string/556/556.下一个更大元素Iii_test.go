package main

import "testing"

func Test_nextGreaterElement(t *testing.T) {
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
			args: args{
				n: 101,
			},
			want: 110,
		},
		{
			name: "test2",
			args: args{
				n: 12,
			},
			want: 21,
		},
		{
			name: "test3",
			args: args{
				n: 230241,
			},
			want: 230412,
		},
		{
			name: "test4",
			args: args{
				n: 2147483486,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElement(tt.args.n); got != tt.want {
				t.Errorf("nextGreaterElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
