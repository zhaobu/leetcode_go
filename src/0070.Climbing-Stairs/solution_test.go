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
	n int
}

type Case struct {
	name   string
	input  *Input
	expect int
}

func TestClimbStairs1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{n: 1}, expect: 1},
		{name: "test 2", input: &Input{n: 2}, expect: 2},
		{name: "test 3", input: &Input{n: 3}, expect: 3},
		{name: "test 4", input: &Input{n: 4}, expect: 5},
		{name: "test 5", input: &Input{n: 5}, expect: 8},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := ClimbStairs1(c.input.n)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}

func TestClimbStairs2(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{n: 1}, expect: 1},
		{name: "test 2", input: &Input{n: 2}, expect: 2},
		{name: "test 3", input: &Input{n: 3}, expect: 3},
		{name: "test 4", input: &Input{n: 4}, expect: 5},
		{name: "test 5", input: &Input{n: 5}, expect: 8},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := ClimbStairs2(c.input.n)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}

func TestClimbStairs3(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{n: 1}, expect: 1},
		{name: "test 2", input: &Input{n: 2}, expect: 2},
		{name: "test 3", input: &Input{n: 3}, expect: 3},
		{name: "test 4", input: &Input{n: 4}, expect: 5},
		{name: "test 5", input: &Input{n: 5}, expect: 8},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := ClimbStairs3(c.input.n)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}

func TestClimbStairs4(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{n: 1}, expect: 1},
		{name: "test 2", input: &Input{n: 2}, expect: 2},
		{name: "test 3", input: &Input{n: 3}, expect: 3},
		{name: "test 4", input: &Input{n: 4}, expect: 5},
		{name: "test 5", input: &Input{n: 5}, expect: 8},
		{name: "test 6", input: &Input{n: 6}, expect: 13},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := ClimbStairs4(c.input.n)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}
