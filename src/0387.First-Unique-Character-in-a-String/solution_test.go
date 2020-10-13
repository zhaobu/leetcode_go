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

func TestFirstUniqChar1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{s: "leetcode"}, expect: 0},
		{name: "test 2", input: &Input{s: "loveleetcode"}, expect: 2},
		{name: "test 3", input: &Input{s: "cc"}, expect: -1},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := FirstUniqChar1(c.input.s)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}

func TestFirstUniqChar2(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{s: "leetcode"}, expect: 0},
		{name: "test 2", input: &Input{s: "loveleetcode"}, expect: 2},
		{name: "test 3", input: &Input{s: "cc"}, expect: -1},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := FirstUniqChar2(c.input.s)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}
