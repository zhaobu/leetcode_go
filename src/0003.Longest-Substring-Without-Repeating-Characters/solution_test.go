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

func TestFindAnagrams1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{s: "abcabcbb"}, expect: 3},
		{name: "test 2", input: &Input{s: "bbbbb"}, expect: 1},
		{name: "test 3", input: &Input{s: "pwwkew"}, expect: 3},
		{name: "test 4", input: &Input{s: " "}, expect: 1},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := lengthOfLongestSubstring1(c.input.s)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}

func TestFindAnagrams2(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{s: "abcabcbb"}, expect: 3},
		{name: "test 2", input: &Input{s: "bbbbb"}, expect: 1},
		{name: "test 3", input: &Input{s: "pwwkew"}, expect: 3},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := lengthOfLongestSubstring2(c.input.s)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}
