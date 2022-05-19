package main

import (
	. "leetcode/check/tree"
	"testing"
)

func Test_isSubStructure(t *testing.T) {
	type args struct {
		A *TreeNode
		B *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubStructure(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("isSubStructure() = %v, want %v", got, tt.want)
			}
		})
	}
}
