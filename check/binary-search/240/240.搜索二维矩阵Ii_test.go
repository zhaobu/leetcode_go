package main

import "testing"

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// {
		// 	name: "test1",
		// 	args: args{
		// 		matrix: [][]int{{-5}},
		// 		target: -2,
		// 	},
		// 	want: false,
		// },
		{
			name: "test2",
			args: args{
				matrix: [][]int{{-1, 3}},
				target: 3,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
