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

func (t *Input) Copy() []int {
	return append(make([]int, 0, len(t.nums)), t.nums...)
}

type Case struct {
	name   string
	input  *Input
	expect []int
}

func TestBubbleSort(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{nums: []int{-2, 2, -6, 2, 9, 3, 42, 5}}, expect: []int{-6, -2, 2, 2, 3, 5, 9, 42}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			copy := c.input.Copy()
			out := BubbleSort(c.input.nums)
			t.Logf("case %s success expect: %v, output: %v, with input: %+v", c.name, c.expect, out, copy)
		})
	}
}

func TestSelectionSort(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{nums: []int{-2, 2, -6, 2, 9, 3, 42, 5}}, expect: []int{-6, -2, 2, 2, 3, 5, 9, 42}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			copy := c.input.Copy()
			out := SelectionSort(c.input.nums)
			t.Logf("case %s success expect: %v, output: %v, with input: %+v", c.name, c.expect, out, copy)
		})
	}
}
