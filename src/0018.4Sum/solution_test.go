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
	nums   []int
	target int
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

func TestFourSum1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{nums: []int{1, 0, -1, 0, -2, 2}, target: 0}, expect: [][]int{{-1, 0, 0, 1}, {-2, -1, 1, 2}, {-2, 0, 0, 2}}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := c.input.Copy()
			out := fourSum1(c.input.nums, c.input.target)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, input)
		})
	}
}
