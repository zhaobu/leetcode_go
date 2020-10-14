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
	A string
	B string
}

type Case struct {
	name   string
	input  *Input
	expect bool
}

func TestRotateString1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{A: "abcde", B: "cdeab"}, expect: true},
		{name: "test 2", input: &Input{A: "abcde", B: "abced"}, expect: false},
		{name: "test 3", input: &Input{A: "", B: ""}, expect: true},
		{name: "test 4", input: &Input{A: "bbbacddceeb", B: "ceebbbbacdd"}, expect: true},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := RotateString1(c.input.A, c.input.B)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}

func TestRotateString2(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{A: "abcde", B: "cdeab"}, expect: true},
		{name: "test 2", input: &Input{A: "abcde", B: "abced"}, expect: false},
		{name: "test 3", input: &Input{A: "", B: ""}, expect: true},
		{name: "test 4", input: &Input{A: "bbbacddceeb", B: "ceebbbbacdd"}, expect: true},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := RotateString2(c.input.A, c.input.B)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}
