package Solution

import (
	"testing"
)

// func init() {
// 	zaplog.InitLog(&zaplog.Config{
// 		Logdir:   "./",
// 		LogName:  "main.log",
// 		Loglevel: "debug",
// 		Debug:    true},
// 	)
// }

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//定义结构
type Input struct {
	root *TreeNode
}

type Case struct {
	name   string
	input  *Input
	expect int
}

func TestRob1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{root: []int{2, 3, 2}}, expect: 3},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := rob1(c.input.nums)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}
