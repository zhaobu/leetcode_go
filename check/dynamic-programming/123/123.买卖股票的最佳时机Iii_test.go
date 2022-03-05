package main

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "1",
		// 	args: args{
		// 		prices: []int{1, 2, 3, 4, 5},
		// 	},
		// 	want: 4,
		// },
		// {
		// 	name: "1",
		// 	args: args{
		// 		prices: []int{1, 4, 2},
		// 	},
		// 	want: 3,
		// },

		{
			name: "1",
			args: args{
				prices: []int{3, 3, 5, 0, 0, 3, 1, 4},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
