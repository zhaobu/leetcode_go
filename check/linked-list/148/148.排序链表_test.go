package main

import (
	. "leetcode/check"
	"reflect"
	"testing"
)

func Test_sortList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// {
		// 	name: "test1",
		// 	args: args{
		// 		head: UnmarshalListBySlice([]int{4, 2, 1, 3}),
		// 	},
		// 	want: UnmarshalListBySlice([]int{1, 2, 3, 4}),
		// },
		// {
		// 	name: "test2",
		// 	args: args{
		// 		head: UnmarshalListBySlice([]int{4, 2}),
		// 	},
		// 	want: UnmarshalListBySlice([]int{2, 4}),
		// },
		// {
		// 	name: "test3",
		// 	args: args{
		// 		head: UnmarshalListBySlice([]int{-1, 5, 3, 4, 0}),
		// 	},
		// 	want: UnmarshalListBySlice([]int{-1, 0, 3, 4, 5}),
		// },
		{
			name: "test4",
			args: args{
				head: UnmarshalListBySlice([]int{4, 19, 14, 5, -3, 1, 8, 5, 11, 15}),
			},
			want: UnmarshalListBySlice([]int{-3, 1, 4, 5, 5, 8, 11, 14, 15, 19}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
