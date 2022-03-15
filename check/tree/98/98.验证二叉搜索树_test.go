package main

import (
	. "leetcode/check/tree"
	"testing"
)

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// {
		// 	/*
		// 		  	  0
		// 		  1      2
		// 		3   4  5    6
		// 	*/
		// 	name: "test1",
		// 	args: args{
		// 		root: BuildTree("546##37"),
		// 	},
		// 	want: false,
		// },
		{
			/*
				  	  0
				  1      2
				3   4  5    6
			*/
			name: "test1",
			args: args{
				root: BuildTree("514##36"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
