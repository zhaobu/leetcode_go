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
	N int
}

type Case struct {
	name   string
	input  *Input
	expect int
}

func TestFib1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{N: 1}, expect: 1},
		{name: "test 2", input: &Input{N: 3}, expect: 2},
		{name: "test 3", input: &Input{N: 0}, expect: 0},
		{name: "test 4", input: &Input{N: 2}, expect: 1},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := fib1(c.input.N)
			t.Logf("case %s success expect: %v, output: %v, with input: %+v", c.name, c.expect, out, c.input)
		})
	}
}

func TestFib2(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{N: 1}, expect: 1},
		{name: "test 2", input: &Input{N: 3}, expect: 2},
		{name: "test 3", input: &Input{N: 0}, expect: 0},
		{name: "test 4", input: &Input{N: 2}, expect: 1},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := fib2(c.input.N)
			t.Logf("case %s success expect: %v, output: %v, with input: %+v", c.name, c.expect, out, c.input)
		})
	}
}

func TestFib3(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{N: 1}, expect: 1},
		{name: "test 2", input: &Input{N: 3}, expect: 2},
		{name: "test 3", input: &Input{N: 0}, expect: 0},
		{name: "test 4", input: &Input{N: 2}, expect: 1},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := fib3(c.input.N)
			t.Logf("case %s success expect: %v, output: %v, with input: %+v", c.name, c.expect, out, c.input)
		})
	}
}

func TestFib4(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{N: 1}, expect: 1},
		{name: "test 2", input: &Input{N: 3}, expect: 2},
		{name: "test 3", input: &Input{N: 0}, expect: 0},
		{name: "test 4", input: &Input{N: 2}, expect: 1},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := fib4(c.input.N)
			t.Logf("case %s success expect: %v, output: %v, with input: %+v", c.name, c.expect, out, c.input)
		})
	}
}
