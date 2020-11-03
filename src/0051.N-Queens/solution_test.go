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
	n int
}

type Case struct {
	name   string
	input  *Input
	expect [][]string
}

func TestPermute1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{n: 4}, expect: [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}}},
	}
	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := solveNQueens1(c.input.n)
			t.Logf("case %s success expect: %v, output: %v, with input: %+v", c.name, c.expect, out, c.input)
		})
	}
}
