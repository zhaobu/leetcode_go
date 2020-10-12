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
	grid [][]int
}

type Case struct {
	name   string
	input  *Input
	expect int
}

func TestMinPathSum1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{grid: [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}}, expect: 7},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := MinPathSum1(c.input.grid)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}
