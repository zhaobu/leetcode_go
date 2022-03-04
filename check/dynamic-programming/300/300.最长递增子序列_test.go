package main

import "testing"

func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		// {
		// 	name: "1",
		// 	args: args{
		// 		nums: []int{1, 3, 6, 7, 9, 4, 10, 5, 6},
		// 	},
		// 	want: 6,
		// },

		{
			name: "2",
			args: args{
				nums: []int{10, 2, 2, 5, 1, 7, 101, 18},
			},
			want: 4,
		},
		{
			name: "3",
			args: args{
				nums: []int{10, 9, 2, 5, 3, 4},
			},
			want: 3,
		},
		{
			name: "4",
			args: args{
				nums: []int{3, 5, 6, 2, 5, 4, 19, 5, 6, 7, 12},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
