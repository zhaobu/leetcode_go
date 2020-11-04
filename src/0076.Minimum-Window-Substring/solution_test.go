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
	t string
}

type Case struct {
	name   string
	input  *Input
	expect string
}

func TestMinWindow1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{s: "ADOBECODEBANC", t: "ABC"}, expect: "BANC"},
		{name: "test 2", input: &Input{s: "ADOBANCDEBBAC", t: "ABC"}, expect: "BAC"},
		{name: "test 3", input: &Input{s: "ADOBANCDEBBAC", t: "ABBC"}, expect: "BBAC"},
		{name: "test 4", input: &Input{s: "ADZXYOBANCDE", t: "XYZ"}, expect: "ZXY"},
		{name: "test 5", input: &Input{s: "a", t: "a"}, expect: "a"},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := minWindow1(c.input.s, c.input.t)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}
