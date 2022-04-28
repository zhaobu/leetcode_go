package main

import "testing"

func Test_findKthNumber(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "test1",
		// 	args: args{
		// 		n: 13,
		// 		k: 2,
		// 	},
		// 	want: 10,
		// },
		{
			name: "test2",
			args: args{
				n: 20,
				k: 15,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthNumber(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("findKthNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
