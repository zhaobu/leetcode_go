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
	triangle [][]int
}

type Case struct {
	name   string
	input  *Input
	expect int
}

func TestMinimumTotal1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{triangle: [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}}, expect: 11},
		{name: "test 2", input: &Input{triangle: [][]int{{-1}, {2, 3}, {1, -1, -3}}}, expect: -1},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := MinimumTotal1(c.input.triangle)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}

func TestMinimumTotal2(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{triangle: [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}}, expect: 11},
		{name: "test 2", input: &Input{triangle: [][]int{{-1}, {2, 3}, {1, -1, -3}}}, expect: -1},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := MinimumTotal2(c.input.triangle)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}

func TestMinimumTotal3(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{triangle: [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}}, expect: 11},
		{name: "test 2", input: &Input{triangle: [][]int{{-1}, {2, 3}, {1, -1, -3}}}, expect: -1},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := MinimumTotal3(c.input.triangle)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}

func TestMinimumTotal4(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{triangle: [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}}, expect: 11},
		{name: "test 2", input: &Input{triangle: [][]int{{-1}, {2, 3}, {1, -1, -3}}}, expect: -1},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := MinimumTotal4(c.input.triangle)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}

func TestMinimumTotal5(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{triangle: [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}}, expect: 11},
		{name: "test 2", input: &Input{triangle: [][]int{{-1}, {2, 3}, {1, -1, -3}}}, expect: -1},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := MinimumTotal5(c.input.triangle)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}
