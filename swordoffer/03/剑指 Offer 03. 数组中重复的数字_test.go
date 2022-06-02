package main

import (
	"testing"
)

func Test_findRepeatNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{1, 0, 3, 4, 2, 5},
			},
			want: -1,
		},
		{
			name: "test2",
			args: args{
				nums: []int{3, 4, 2, 0, 0, 1},
			},
			want: 0,
		},
		{
			name: "test3",
			args: args{
				nums: []int{3, 3, 2, 0, 5, 1},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRepeatNumber(tt.args.nums); got != tt.want {
				t.Errorf("findRepeatNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
