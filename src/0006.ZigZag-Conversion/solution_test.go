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
	s       string
	numRows int
}

type Case struct {
	name   string
	input  *Input
	expect string
}

func TestConvert1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{s: "LEETCODEISHIRING", numRows: 3}, expect: "LCIRETOESIIGEDHN"},
		{name: "test 2", input: &Input{s: "LEETCODEISHIRING", numRows: 4}, expect: "LDREOEIIECIHNTSG"},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := Convert1(c.input.s, c.input.numRows)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}

func TestConvert2(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{s: "LEETCODEISHIRING", numRows: 3}, expect: "LCIRETOESIIGEDHN"},
		{name: "test 2", input: &Input{s: "LEETCODEISHIRING", numRows: 4}, expect: "LDREOEIIECIHNTSG"},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := Convert2(c.input.s, c.input.numRows)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}
