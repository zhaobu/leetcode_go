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

func TestRob1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{nums: []int{1, 2, 3, 1}}, expect: 4},
		{name: "test 2", input: &Input{nums: []int{2, 7, 9, 3, 1}}, expect: 12},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := Rob1(c.input.nums)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}

func TestRob2(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{nums: []int{1, 2, 3, 1}}, expect: 4},
		{name: "test 2", input: &Input{nums: []int{2, 7, 9, 3, 1}}, expect: 12},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := Rob2(c.input.nums)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}

func TestRob3(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{nums: []int{1, 2, 3, 1}}, expect: 4},
		{name: "test 2", input: &Input{nums: []int{2, 7, 9, 3, 1}}, expect: 12},
		{name: "test 3", input: &Input{nums: []int{1, 1}}, expect: 1},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := Rob3(c.input.nums)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}
