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

func (i *Input) Copy() *Input {
	res := &Input{
		nums: make([]int, len(i.nums)),
	}
	copy(res.nums, i.nums)
	return res
}

type Case struct {
	name   string
	input  *Input
	expect [][]int
}

func TestThreeSum1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{nums: []int{-1, 0, 1, 2, -1, -4}}, expect: [][]int{{-1, 0, 1}, {-1, -1, 2}}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := c.input.Copy()
			out := ThreeSum1(c.input.nums)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, input)
		})
	}
}

func TestThreeSum2(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{nums: []int{-1, 0, 1, 2, -1, -4}}, expect: [][]int{{-1, 0, 1}, {-1, -1, 2}}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := c.input.Copy()
			out := ThreeSum2(c.input.nums)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, input)
		})
	}
}

func TestThreeSum3(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{nums: []int{-1, 0, 1, 2, -1, -4}}, expect: [][]int{{-1, 0, 1}, {-1, -1, 2}}},
		{name: "test 2", input: &Input{nums: []int{0, 0}}, expect: [][]int{}},
		{name: "test 3", input: &Input{nums: []int{1, 2, -2, -1}}, expect: [][]int{}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := c.input.Copy()
			out := ThreeSum3(c.input.nums)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, input)
		})
	}
}
