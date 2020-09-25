package Solution

import (
	"leetcode/utils/zaplog"
	"testing"
)

func init() {
	zaplog.InitLog(&zaplog.Config{
		Logdir:   "./",
		LogName:  "main.log",
		Loglevel: "debug",
		Debug:    true},
	)
}

//定义结构
type Input struct {
	Strs []string
}

type Case struct {
	name   string
	input  *Input
	expect string
}

func TestLongestCommonPrefix1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{Strs: []string{"flower", "flow", "flight"}}, expect: "fl"},
		{name: "test 1", input: &Input{Strs: []string{"a"}}, expect: "a"},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			output := LongestCommonPrefix1(c.input.Strs)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, output, c.input)
		})
	}
}
