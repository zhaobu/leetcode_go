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
	coins  []int
	amount int
}

type Case struct {
	name   string
	input  *Input
	expect int
}

func TestCoinChange1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{coins: []int{1, 2, 5}, amount: 11}, expect: 3},
		{name: "test 2", input: &Input{coins: []int{2}, amount: 3}, expect: -1},
		{name: "test 3", input: &Input{coins: []int{1}, amount: 0}, expect: 0},
		{name: "test 4", input: &Input{coins: []int{1}, amount: 1}, expect: 1},
		{name: "test 5", input: &Input{coins: []int{1}, amount: 2}, expect: 2}}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := coinChange1(c.input.coins, c.input.amount)
			t.Logf("case %s success expect: %v, output: %v, with input: %+v", c.name, c.expect, out, c.input)
		})
	}
}

func TestCoinChange2(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{coins: []int{1, 2, 5}, amount: 11}, expect: 3},
		{name: "test 2", input: &Input{coins: []int{2}, amount: 3}, expect: -1},
		{name: "test 3", input: &Input{coins: []int{1}, amount: 0}, expect: 0},
		{name: "test 4", input: &Input{coins: []int{1}, amount: 1}, expect: 1},
		{name: "test 5", input: &Input{coins: []int{1}, amount: 2}, expect: 2}}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := coinChange2(c.input.coins, c.input.amount)
			t.Logf("case %s success expect: %v, output: %v, with input: %+v", c.name, c.expect, out, c.input)
		})
	}
}

func TestCoinChange3(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{coins: []int{1, 2, 5}, amount: 11}, expect: 3},
		{name: "test 2", input: &Input{coins: []int{2}, amount: 3}, expect: -1},
		{name: "test 3", input: &Input{coins: []int{1}, amount: 0}, expect: 0},
		{name: "test 4", input: &Input{coins: []int{1}, amount: 1}, expect: 1},
		{name: "test 5", input: &Input{coins: []int{1}, amount: 2}, expect: 2}}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := coinChange3(c.input.coins, c.input.amount)
			t.Logf("case %s success expect: %v, output: %v, with input: %+v", c.name, c.expect, out, c.input)
		})
	}
}
