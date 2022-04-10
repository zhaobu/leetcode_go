package main

import (
	. "leetcode/check"
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test1",
			args: args{
				lists: []*ListNode{
					// [[-10,-9,-9,-3,-1,-1,0],[-5],[4],[-8],[],[-9,-6,-5,-4,-2,2,3],[-3,-3,-2,-1,0]]
					UnmarshalListBySlice([]int{-10, -9, -9, -3, -1, -1, 0}),
					UnmarshalListBySlice([]int{-5}),
					UnmarshalListBySlice([]int{4}),
					UnmarshalListBySlice([]int{-8}),
					UnmarshalListBySlice([]int{}),
					UnmarshalListBySlice([]int{-9, -6, -5, -4, -2, 2, 3}),
					UnmarshalListBySlice([]int{-3, -3, -2, -1, 0}),
				},
			},
			want: UnmarshalListBySlice([]int{
				-10, -9, -9, -9, -8, -6, -5, -5, -4, -3, -3, -3, -2, -2, -1, -1, -1, 0, 0, 2, 3, 4,
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
