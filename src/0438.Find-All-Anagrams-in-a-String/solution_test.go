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
	p string
}

type Case struct {
	name   string
	input  *Input
	expect []int
}

func TestFindAnagrams1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{s: "cbaebabacd", p: "abc"}, expect: []int{0, 6}},
		{name: "test 2", input: &Input{s: "zyxcabixacb", p: "abc"}, expect: []int{3, 8}},
		{name: "test 3", input: &Input{s: "abab", p: "ab"}, expect: []int{0, 1, 2}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := findAnagrams1(c.input.s, c.input.p)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}

func TestFindAnagrams2(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{s: "cbaebabacd", p: "abc"}, expect: []int{0, 6}},
		{name: "test 2", input: &Input{s: "zyxcabixacb", p: "abc"}, expect: []int{3, 8}},
		{name: "test 3", input: &Input{s: "abab", p: "ab"}, expect: []int{0, 1, 2}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := findAnagrams2(c.input.s, c.input.p)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}
