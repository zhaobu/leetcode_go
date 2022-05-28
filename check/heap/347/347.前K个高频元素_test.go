package main

import (
	"reflect"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{1, 1, 1, 2, 2, 3},
				k:    2,
			},
			want: []int{1, 2},
		},
		{
			name: "test2",
			args: args{
				nums: []int{1},
				k:    1,
			},
			want: []int{1},
		},
		{
			name: "test3",
			args: args{
				nums: []int{1, 2},
				k:    2,
			},
			want: []int{1, 2},
		},
		{
			name: "test3",
			args: args{
				nums: []int{2, 3, 4, 1, 4, 0, 4, -1, -2, -1},
				k:    2,
			},
			want: []int{4, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
