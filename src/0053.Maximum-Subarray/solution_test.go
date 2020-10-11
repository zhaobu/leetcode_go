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
}

type Case struct {
	name   string
	input  *Input
	expect int
}

func TestMaxSubArray1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}}, expect: 6},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := MaxSubArray1(c.input.nums)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}
func TestMaxSubArray2(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}}, expect: 6},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := MaxSubArray2(c.input.nums)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}

func TestMaxSubArray3(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}}, expect: 6},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := MaxSubArray3(c.input.nums)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}
