package main

import "testing"

func Test_findNumberOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				nums: []int{1, 3, 5, 4, 7},
			},
			want: 2,
		},
		{
			name: "test2",
			args: args{
				nums: []int{2, 2, 2, 2, 2},
			},
			want: 5,
		},
		{
			name: "test3",
			args: args{
				nums: []int{1, 2, 4, 3, 5, 4, 7, 2},
			},
			want: 3,
		},
		{
			name: "test4",
			args: args{
				nums: []int{3, 1, 2},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumberOfLIS(tt.args.nums); got != tt.want {
				t.Errorf("findNumberOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
