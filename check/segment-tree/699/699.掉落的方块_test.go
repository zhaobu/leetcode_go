package main

import (
	"reflect"
	"testing"
)

func Test_fallingSquares(t *testing.T) {
	type args struct {
		positions [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				positions: [][]int{
					{6, 1},
					{9, 2},
					{2, 4}},
				/*
					            a1(6,6,1)
											a2(9,10,2)
							a3(2,5,4)

				*/
			},
			want: []int{1, 2, 4},
		},
		{
			name: "test2",
			args: args{
				positions: [][]int{
					{9, 6},
					{2, 2},
					{2, 6}},
				/*
					            a0(9,14,6)
					a1(2,3,2)
					a2(2,7,6)

				*/
			},
			want: []int{6, 6, 8},
		},
		{
			name: "test3",
			args: args{
				positions: [][]int{
					{2, 1},
					{2, 9},
					{1, 8}},
				/*
				 */
			},
			want: []int{1, 10, 18},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fallingSquares(tt.args.positions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fallingSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
