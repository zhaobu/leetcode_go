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
	prices []int
	fee    int
}

type Case struct {
	name   string
	input  *Input
	expect int
}

func TestMaxProfit1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{prices: []int{1, 3, 2, 8, 4, 9}, fee: 2}, expect: 8},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			output := MaxProfit1(c.input.prices, c.input.fee)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, output, c.input)
		})
	}
}