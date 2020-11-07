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
	k      int
}

type Case struct {
	name   string
	input  *Input
	expect int
}

func TestMaxProfit1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{prices: []int{2, 4, 1}, k: 2}, expect: 2},
		{name: "test 2", input: &Input{prices: []int{3, 2, 6, 5, 0, 3}, k: 2}, expect: 7},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			output := MaxProfit1(c.input.k, c.input.prices)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, output, c.input)
		})
	}
}

func TestMaxProfit2(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{prices: []int{2, 4, 1}, k: 2}, expect: 2},
		{name: "test 2", input: &Input{prices: []int{3, 2, 6, 5, 0, 3}, k: 2}, expect: 7},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			output := MaxProfit2(c.input.k, c.input.prices)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, output, c.input)
		})
	}
}
