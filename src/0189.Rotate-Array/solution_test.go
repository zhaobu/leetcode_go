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

//定义结构
type Input struct {
	nums []int
	k    int
}

type Case struct {
	name   string
	input  *Input
	expect []int
}

func TestRotate1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 3}, expect: []int{5, 6, 7, 1, 2, 3, 4}},
		{name: "test 2", input: &Input{nums: []int{-1, -100, 3, 99}, k: 2}, expect: []int{3, 99, -1, -100}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			Rotate1(c.input.nums, c.input.k)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, c.input.nums, c.input)
		})
	}
}

func TestRotate2(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 3}, expect: []int{5, 6, 7, 1, 2, 3, 4}},
		{name: "test 2", input: &Input{nums: []int{-1, -100, 3, 99}, k: 2}, expect: []int{3, 99, -1, -100}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			Rotate2(c.input.nums, c.input.k)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, c.input.nums, c.input)
		})
	}
}
