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
	piles []int
	H     int
}

type Case struct {
	name   string
	input  *Input
	expect int
}

func TestLengthOfLastWord1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{piles: []int{3, 6, 7, 11}, H: 8}, expect: 4},
		{name: "test 2", input: &Input{piles: []int{30, 11, 23, 4, 20}, H: 5}, expect: 30},
		{name: "test 3", input: &Input{piles: []int{30, 11, 23, 4, 20}, H: 6}, expect: 23},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := MinEatingSpeed1(c.input.piles, c.input.H)
			t.Logf("case %s success expect: %v, output: %v, with input: %+v", c.name, c.expect, out, c.input)
		})
	}
}

func TestLengthOfLastWord2(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{piles: []int{3, 6, 7, 11}, H: 8}, expect: 4},
		{name: "test 2", input: &Input{piles: []int{30, 11, 23, 4, 20}, H: 5}, expect: 30},
		{name: "test 3", input: &Input{piles: []int{30, 11, 23, 4, 20}, H: 6}, expect: 23},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := MinEatingSpeed2(c.input.piles, c.input.H)
			t.Logf("case %s success expect: %v, output: %v, with input: %+v", c.name, c.expect, out, c.input)
		})
	}
}
