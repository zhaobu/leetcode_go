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
	haystack string
	needle   string
}

type Case struct {
	name   string
	input  *Input
	expect int
}

func TestStrStr1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{haystack: "hello", needle: "ll"}, expect: 2},
		{name: "test 2", input: &Input{haystack: "aaaaa", needle: "bba"}, expect: -1},
		{name: "test 2", input: &Input{haystack: "mississippi", needle: "pi"}, expect: 9},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := StrStr1(c.input.haystack, c.input.needle)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}

func TestStrStr2(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{haystack: "mississippi", needle: "pi"}, expect: 9},
		{name: "test 2", input: &Input{haystack: "hello", needle: "ll"}, expect: 2},
		{name: "test 3", input: &Input{haystack: "aaaaa", needle: "bba"}, expect: -1},
		{name: "test 4", input: &Input{haystack: "", needle: "a"}, expect: -1},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := StrStr2(c.input.haystack, c.input.needle)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}

func TestStrStr3(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{haystack: "mississippi", needle: "pi"}, expect: 9},
		{name: "test 2", input: &Input{haystack: "hello", needle: "ll"}, expect: 2},
		{name: "test 3", input: &Input{haystack: "aaaaa", needle: "bba"}, expect: -1},
		{name: "test 4", input: &Input{haystack: "", needle: "a"}, expect: -1},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := StrStr3(c.input.haystack, c.input.needle)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}

func TestStrStr4(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{haystack: "mississippi", needle: "pi"}, expect: 9},
		{name: "test 2", input: &Input{haystack: "hello", needle: "ll"}, expect: 2},
		{name: "test 3", input: &Input{haystack: "aaaaa", needle: "bba"}, expect: -1},
		{name: "test 4", input: &Input{haystack: "", needle: "a"}, expect: -1},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := StrStr4(c.input.haystack, c.input.needle)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}
