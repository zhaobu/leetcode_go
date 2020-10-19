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

type Case struct {
	name   string
	input  *Input
	expect int
}

func TestSearch1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{nums: []int{-1, 0, 3, 5, 9, 12}, target: -2}, expect: -1},
		{name: "test 2", input: &Input{nums: []int{-1, 0, 3, 5, 9, 12}, target: -1}, expect: 0},
		{name: "test 3", input: &Input{nums: []int{-1, 0, 3, 5, 9, 12}, target: 9}, expect: 4},
		{name: "test 4", input: &Input{nums: []int{-1, 0, 3, 5, 9, 12}, target: 4}, expect: -1},
		{name: "test 5", input: &Input{nums: []int{-1, 0, 3, 5, 9, 12}, target: 12}, expect: 5},
		{name: "test 6", input: &Input{nums: []int{-1, 0, 3, 5, 9, 12}, target: 15}, expect: -1},
		{name: "test 7", input: &Input{nums: []int{3}, target: 3}, expect: 0},
		{name: "test 8", input: &Input{nums: []int{}, target: 3}, expect: -1},
	}
	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := Search1(c.input.nums, c.input.target)
			t.Logf("case %s success expect: %v, output: %v, with input: %+v", c.name, c.expect, out, c.input)
		})
	}
}

func TestSearch2(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{nums: []int{-1, 0, 3, 5, 9, 12}, target: -2}, expect: -1},
		{name: "test 2", input: &Input{nums: []int{-1, 0, 3, 5, 9, 12}, target: -1}, expect: 0},
		{name: "test 3", input: &Input{nums: []int{-1, 0, 3, 5, 9, 12}, target: 9}, expect: 4},
		{name: "test 4", input: &Input{nums: []int{-1, 0, 3, 5, 9, 12}, target: 4}, expect: -1},
		{name: "test 5", input: &Input{nums: []int{-1, 0, 3, 5, 9, 12}, target: 12}, expect: 5},
		{name: "test 6", input: &Input{nums: []int{-1, 0, 3, 5, 9, 12}, target: 15}, expect: -1},
		{name: "test 7", input: &Input{nums: []int{3}, target: 3}, expect: 0},
		{name: "test 8", input: &Input{nums: []int{}, target: 3}, expect: -1},
	}
	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := Search2(c.input.nums, c.input.target)
			t.Logf("case %s success expect: %v, output: %v, with input: %+v", c.name, c.expect, out, c.input)
		})
	}
}

func TestSearch3(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{nums: []int{-1, 0, 3, 5, 9, 12}, target: -2}, expect: -1},
		{name: "test 2", input: &Input{nums: []int{-1, 0, 3, 5, 9, 12}, target: -1}, expect: 0},
		{name: "test 3", input: &Input{nums: []int{-1, 0, 3, 5, 9, 12}, target: 9}, expect: 4},
		{name: "test 4", input: &Input{nums: []int{-1, 0, 3, 5, 9, 12}, target: 4}, expect: -1},
		{name: "test 5", input: &Input{nums: []int{-1, 0, 3, 5, 9, 12}, target: 12}, expect: 5},
		{name: "test 6", input: &Input{nums: []int{-1, 0, 3, 5, 9, 12}, target: 15}, expect: -1},
		{name: "test 7", input: &Input{nums: []int{3}, target: 3}, expect: 0},
		{name: "test 8", input: &Input{nums: []int{}, target: 3}, expect: -1},
	}
	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := Search3(c.input.nums, c.input.target)
			t.Logf("case %s success expect: %v, output: %v, with input: %+v", c.name, c.expect, out, c.input)
		})
	}
}

func TestLeftBound1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 0}, expect: -1},
		{name: "test 2", input: &Input{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 1}, expect: 0},
		{name: "test 3", input: &Input{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 4}, expect: 4},
		{name: "test 4", input: &Input{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 3}, expect: -1},
		{name: "test 5", input: &Input{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 5}, expect: 6},
		{name: "test 6", input: &Input{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 6}, expect: -1},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := LeftBound1(c.input.nums, c.input.target)
			t.Logf("case %s success expect: %v, output: %v, with input: %+v", c.name, c.expect, out, c.input)
		})
	}
}

func TestLeftBound2(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 0}, expect: -1},
		{name: "test 2", input: &Input{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 1}, expect: 0},
		{name: "test 3", input: &Input{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 4}, expect: 4},
		{name: "test 4", input: &Input{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 3}, expect: -1},
		{name: "test 5", input: &Input{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 5}, expect: 6},
		{name: "test 6", input: &Input{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 6}, expect: -1},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := LeftBound2(c.input.nums, c.input.target)
			t.Logf("case %s success expect: %v, output: %v, with input: %+v", c.name, c.expect, out, c.input)
		})
	}
}

func TestRightBound1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 0}, expect: -1},
		{name: "test 2", input: &Input{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 1}, expect: 1},
		{name: "test 3", input: &Input{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 4}, expect: 5},
		{name: "test 4", input: &Input{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 3}, expect: -1},
		{name: "test 5", input: &Input{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 5}, expect: 7},
		{name: "test 6", input: &Input{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 6}, expect: -1},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := RightBound1(c.input.nums, c.input.target)
			t.Logf("case %s success expect: %v, output: %v, with input: %+v", c.name, c.expect, out, c.input)
		})
	}
}

func TestRightBound2(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 0}, expect: -1},
		{name: "test 2", input: &Input{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 1}, expect: 1},
		{name: "test 3", input: &Input{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 4}, expect: 5},
		{name: "test 4", input: &Input{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 3}, expect: -1},
		{name: "test 5", input: &Input{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 5}, expect: 7},
		{name: "test 6", input: &Input{nums: []int{1, 1, 2, 2, 4, 4, 5, 5}, target: 6}, expect: -1},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := RightBound2(c.input.nums, c.input.target)
			t.Logf("case %s success expect: %v, output: %v, with input: %+v", c.name, c.expect, out, c.input)
		})
	}
}
