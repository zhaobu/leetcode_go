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
	expect bool
}

func TestIsPalindrome1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{s: "A man, a plan, a canal: Panama"}, expect: true},
		{name: "test 2", input: &Input{s: "race a car"}, expect: false},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := IsPalindrome1(c.input.s)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}
