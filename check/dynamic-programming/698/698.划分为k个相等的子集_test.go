package main

import "testing"

func Test_canPartitionKSubsets(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				nums: []int{1, 2, 2, 2, 2},
				k:    3,
			},
			want: false,
		},
		{

			name: "test2",
			args: args{
				nums: []int{1, 1, 1, 1, 9},
				k:    2,
			},
			want: false,
		},
		{

			name: "test3",
			args: args{
				nums: []int{3522, 181, 521, 515, 304, 123, 2512, 312, 922, 407, 146, 1932, 4037, 2646, 3871, 269},
				k:    5,
			},
			want: true,
		},
		{

			name: "test4",
			args: args{
				nums: []int{3, 3, 10, 2, 6, 5, 10, 6, 8, 3, 2, 1, 6, 10, 7, 2},
				k:    6,
			},
			want: false,
		},
		{

			name: "test4",
			args: args{
				nums: []int{10, 1, 10, 9, 6, 1, 9, 5, 9, 10, 7, 8, 5, 2, 10, 8},
				k:    11,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPartitionKSubsets(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("canPartitionKSubsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
