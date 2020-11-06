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
	s1 string
	s2 string
}

type Case struct {
	name   string
	input  *Input
	expect bool
}

func TestCheckInclusion1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{s1: "ab", s2: "eidbaooo"}, expect: true},
		{name: "test 2", input: &Input{s1: "ab", s2: "eidboaoo"}, expect: false},
		{name: "test 3", input: &Input{s1: "abc", s2: "bbbca"}, expect: true},
		{name: "test 4", input: &Input{s1: "abcdxabcde", s2: "abcdeabcdx"}, expect: true},
	}
	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := checkInclusion1(c.input.s1, c.input.s2)
			t.Logf("case %s success expect: %v, output: %v, with input: %+v", c.name, c.expect, out, c.input)
		})
	}
}

func TestCheckInclusion2(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{s1: "ab", s2: "eidbaooo"}, expect: true},
		{name: "test 2", input: &Input{s1: "ab", s2: "eidboaoo"}, expect: false},
		{name: "test 3", input: &Input{s1: "abc", s2: "bbbca"}, expect: true},
		{name: "test 4", input: &Input{s1: "abcdxabcde", s2: "abcdeabcdx"}, expect: true},
	}
	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := checkInclusion2(c.input.s1, c.input.s2)
			t.Logf("case %s success expect: %v, output: %v, with input: %+v", c.name, c.expect, out, c.input)
		})
	}
}

func TestCheckInclusion3(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{s1: "ab", s2: "eidbaooo"}, expect: true},
		{name: "test 2", input: &Input{s1: "ab", s2: "eidboaoo"}, expect: false},
		{name: "test 3", input: &Input{s1: "abc", s2: "bbbca"}, expect: true},
		{name: "test 4", input: &Input{s1: "abcdxabcde", s2: "abcdeabcdx"}, expect: true},
	}
	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := checkInclusion3(c.input.s1, c.input.s2)
			t.Logf("case %s success expect: %v, output: %v, with input: %+v", c.name, c.expect, out, c.input)
		})
	}
}
