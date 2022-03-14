package main

import (
	. "leetcode/check/tree"
	"reflect"
	"testing"
)

func Test_postorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			/*
				  	  0
				  1      2
				3   4  5    6
			*/
			name: "test1",
			args: args{
				root: BuildTree("0123456"),
			},
			want: []int{3, 4, 1, 5, 6, 2, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
