package main

import (
	"fmt"
	"testing"
)

func Test_numEnclaves(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "test1",
		// 	args: args{
		// 		grid: [][]int{
		// 			{0, 1, 1, 0},
		// 			{0, 0, 1, 0},
		// 			{0, 0, 1, 0},
		// 			{0, 0, 0, 0},
		// 		},
		// 	},
		// 	want: 0,
		// },

		// {
		// 	name: "test2",
		// 	args: args{
		// 		grid: [][]int{
		// 			{0, 0, 0, 0},
		// 			{1, 0, 1, 0},
		// 			{0, 1, 1, 0},
		// 			{0, 0, 0, 0},
		// 		},
		// 	},
		// 	want: 3,
		// },
		{
			name: "test3",
			args: args{
				grid: [][]int{
					{0, 0, 0, 1, 1, 1, 0, 1, 0, 0},
					{1, 1, 0, 0, 0, 1, 0, 1, 1, 1},
					{0, 0, 0, 1, 1, 1, 0, 1, 0, 0},
					{0, 1, 1, 0, 0, 0, 1, 0, 1, 0},
					{0, 1, 1, 1, 1, 1, 0, 0, 1, 0},
					{0, 0, 1, 0, 1, 1, 1, 1, 0, 1},
					{0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
					{0, 0, 1, 0, 0, 1, 0, 1, 0, 1},
					{1, 0, 1, 0, 1, 1, 0, 0, 0, 0},
					{0, 0, 0, 0, 1, 1, 0, 0, 0, 1},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numEnclaves(tt.args.grid); got != tt.want {
				t.Errorf("numEnclaves() = %v, want %v", got, tt.want)
			}
		})
	}
}

func printGrid(grid [][]int) {
	fmt.Printf("\n")
	for i := 0; i < len(grid); i++ {
		fmt.Printf("grid[%d]=%+v\n", i, grid[i])
	}
}
