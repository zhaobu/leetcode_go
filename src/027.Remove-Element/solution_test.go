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
	val  int
}

func (i *Input) Copy() *Input {
	res := &Input{
		nums: make([]int, len(i.nums)),
		val:  i.val,
	}
	copy(res.nums, i.nums)
	return res
}

type Case struct {
	name   string
	input  *Input
	expect int
}

func TestRemoveElement1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{nums: []int{3, 2, 2, 3}, val: 3}, expect: 2},
		{name: "test 1", input: &Input{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2}, expect: 5},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := c.input.Copy()
			out := RemoveElement1(c.input.nums, c.input.val)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, input)
		})
	}
}

func TestRemoveElement2(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{nums: []int{3, 2, 2, 3}, val: 3}, expect: 2},
		{name: "test 1", input: &Input{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2}, expect: 5},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := c.input.Copy()
			out := RemoveElement2(c.input.nums, c.input.val)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, input)
		})
	}
}
