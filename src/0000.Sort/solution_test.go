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
		{name: "test 2", input: &Input{nums: []int{-9, 100, 0, 2, 1, 11, 7, 4}}, expect: []int{-9, 0, 1, 2, 4, 7, 11, 100}},
		{name: "test 3", input: &Input{nums: []int{299, 10, -5, -3, -3, 2, -3, 5}}, expect: []int{-5, -3, -3, -3, 2, 5, 10, 299}},
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
		{name: "test 2", input: &Input{nums: []int{-9, 100, 0, 2, 1, 11, 7, 4}}, expect: []int{-9, 0, 1, 2, 4, 7, 11, 100}},
		{name: "test 3", input: &Input{nums: []int{299, 10, -5, -3, -3, 2, -3, 5}}, expect: []int{-5, -3, -3, -3, 2, 5, 10, 299}},
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

func TestInsertionSort(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{nums: []int{-2, 2, -6, 2, 9, 3, 42, 5}}, expect: []int{-6, -2, 2, 2, 3, 5, 9, 42}},
		{name: "test 2", input: &Input{nums: []int{-9, 100, 0, 2, 1, 11, 7, 4}}, expect: []int{-9, 0, 1, 2, 4, 7, 11, 100}},
		{name: "test 3", input: &Input{nums: []int{299, 10, -5, -3, -3, 2, -3, 5}}, expect: []int{-5, -3, -3, -3, 2, 5, 10, 299}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			copy := c.input.Copy()
			out := InsertionSort(c.input.nums)
			t.Logf("case %s success expect: %v, output: %v, with input: %+v", c.name, c.expect, out, copy)
		})
	}
}

func TestShellSort(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{nums: []int{5, 7, 8, 3, 1, 2, 4, 6}}, expect: []int{1, 2, 3, 4, 5, 6, 7, 8}},
		{name: "test 2", input: &Input{nums: []int{40, 100, 9, 5, -2, 3, 2, -7, 2}}, expect: []int{-7, -2, 2, 2, 3, 5, 9, 40, 100}},
		{name: "test 3", input: &Input{nums: []int{42, 9, 5, -2, 3, 2, -6, 2}}, expect: []int{-6, -2, 2, 2, 3, 5, 9, 42}},
		{name: "test 4", input: &Input{nums: []int{-9, 100, 0, 2, 1, 11, 7, 4}}, expect: []int{-9, 0, 1, 2, 4, 7, 11, 100}},
		{name: "test 5", input: &Input{nums: []int{299, 10, -5, -3, -3, 2, -3, 5}}, expect: []int{-5, -3, -3, -3, 2, 5, 10, 299}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			copy := c.input.Copy()
			out := ShellSort(c.input.nums)
			t.Logf("case %s success expect: %v, output: %v, with input: %+v", c.name, c.expect, out, copy)
		})
	}
}

func TestMergeSort(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{nums: []int{5, 7, 8, 3, 1, 2, 4, 6}}, expect: []int{1, 2, 3, 4, 5, 6, 7, 8}},
		{name: "test 2", input: &Input{nums: []int{40, 100, 9, 5, -2, 3, 2, -7, 2}}, expect: []int{-7, -2, 2, 2, 3, 5, 9, 40, 100}},
		{name: "test 3", input: &Input{nums: []int{42, 9, 5, -2, 3, 2, -6, 2}}, expect: []int{-6, -2, 2, 2, 3, 5, 9, 42}},
		{name: "test 4", input: &Input{nums: []int{-9, 100, 0, 2, 1, 11, 7, 4}}, expect: []int{-9, 0, 1, 2, 4, 7, 11, 100}},
		{name: "test 5", input: &Input{nums: []int{299, 10, -5, -3, -3, 2, -3, 5}}, expect: []int{-5, -3, -3, -3, 2, 5, 10, 299}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			copy := c.input.Copy()
			out := MergeSort(c.input.nums)
			t.Logf("case %s success expect: %v, output: %v, with input: %+v", c.name, c.expect, out, copy)
		})
	}
}
