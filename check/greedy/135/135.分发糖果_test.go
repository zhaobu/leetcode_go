package main

import "testing"

func Test_candy(t *testing.T) {
	type args struct {
		ratings []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				ratings: []int{1, 0, 2},
				/*
				  1   2
				    0
				  2, 1, 2
				*/
			},
			want: 5,
		},
		{
			name: "test2",
			args: args{
				ratings: []int{1, 2, 2},
			},
			want: 4,
		},
		{
			name: "test3",
			args: args{
				ratings: []int{1, 2, 3, 4, 5, 6},
			},
			want: 21,
		},
		{
			name: "test4",
			args: args{
				ratings: []int{1, 2, 2, 2, 3, 4, 3, 5},
				/*

							  4   5
					        3   3
					  2 2 2
					1

					1+2+1+1+2+3+1+2
				*/
			},
			want: 13,
		},
		{
			name: "test5",
			args: args{
				ratings: []int{1, 3, 2, 2, 1},
				/*
						 3
						1  2 2
						       1
					   1+2+1+2+1

				*/
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := candy(tt.args.ratings); got != tt.want {
				t.Errorf("candy() = %v, want %v", got, tt.want)
			}
		})
	}
}
