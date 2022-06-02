package main

import (
	"reflect"
	"testing"
)

func Test_findDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{5, 4, 6, 7, 9, 3, 10, 9, 5, 6},
			},
			want: []int{9, 6, 5},
		},
		{
			name: "test2",
			args: args{
				nums: []int{1, 1, 2},
			},
			want: []int{1},
		},
		{
			name: "test3",
			args: args{
				nums: []int{10, 2, 5, 10, 9, 1, 1, 4, 3, 7},
			},
			want: []int{10, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if got := findDuplicates1(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("findDuplicates() = %v, want %v", got, tt.want)
			// }
			if got := findDuplicates(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
