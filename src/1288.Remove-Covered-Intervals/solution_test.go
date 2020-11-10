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
	intervals [][]int
}

type Case struct {
	name   string
	input  *Input
	expect int
}

func TestRob1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{intervals: [][]int{{1, 4}, {3, 6}, {2, 8}}}, expect: 2}}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := removeCoveredIntervals1(c.input.intervals)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}

func TestRob2(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{intervals: [][]int{{1, 4}, {3, 6}, {2, 8}}}, expect: 2}}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := removeCoveredIntervals2(c.input.intervals)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}
