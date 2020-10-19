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
	x int
}

type Case struct {
	name   string
	input  *Input
	expect int
}

func TestLengthOfLastWord1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{x: 4}, expect: 2},
		{name: "test 2", input: &Input{x: 8}, expect: 2},
		{name: "test 3", input: &Input{x: 0}, expect: 0},
		{name: "test 4", input: &Input{x: 1}, expect: 1},
		{name: "test 5", input: &Input{x: 6}, expect: 2},
		{name: "test 6", input: &Input{x: 9}, expect: 3},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := MySqrt1(c.input.x)
			t.Logf("case %s success expect: %v, output: %v, with input: %+v", c.name, c.expect, out, c.input)
		})
	}
}

func TestLengthOfLastWord2(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{x: 4}, expect: 2},
		{name: "test 2", input: &Input{x: 8}, expect: 2},
		{name: "test 3", input: &Input{x: 0}, expect: 0},
		{name: "test 4", input: &Input{x: 1}, expect: 1},
		{name: "test 5", input: &Input{x: 6}, expect: 2},
		{name: "test 6", input: &Input{x: 9}, expect: 3},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := MySqrt2(c.input.x)
			t.Logf("case %s success expect: %v, output: %v, with input: %+v", c.name, c.expect, out, c.input)
		})
	}
}