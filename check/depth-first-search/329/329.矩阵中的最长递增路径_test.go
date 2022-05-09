package main

import "testing"

func Test_longestIncreasingPath(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				matrix: [][]int{
					{9, 9, 4},
					{6, 6, 8},
					{2, 1, 1},
				},
			},
			want: 4,
		},
		// {
		// 	name: "test2",
		// 	args: args{
		// 		matrix: [][]int{
		// 			{13, 6, 16, 6, 16, 4},
		// 			{9, 13, 5, 13, 7, 11},
		// 			{11, 7, 9, 17, 0, 7},
		// 			{7, 8, 5, 14, 11, 8},
		// 			{14, 2, 8, 7, 9, 5},
		// 			{1, 15, 3, 11, 11, 6},
		// 		},
		// 	},
		// 	want: 6,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestIncreasingPath(tt.args.matrix); got != tt.want {
				t.Errorf("longestIncreasingPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
