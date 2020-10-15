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
	s string
}

type Case struct {
	name   string
	input  *Input
	expect int
}

func TestLengthOfLastWord1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{s: "Hello World"}, expect: 5},
		{name: "test 2", input: &Input{s: "Hello World  "}, expect: 5},
		{name: "test 3", input: &Input{s: "Hello Wor ld  "}, expect: 2},
		{name: "test 4", input: &Input{s: " Hello"}, expect: 5},
		{name: "test 5", input: &Input{s: "Hello "}, expect: 5},
		{name: "test 6", input: &Input{s: "Hello"}, expect: 5},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := LengthOfLastWord1(c.input.s)
			t.Logf("case %s success expect: %v, output: %v, with input: %+v", c.name, c.expect, out, c.input)
		})
	}
}
