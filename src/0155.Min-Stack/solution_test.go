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
	nums []int
}

type Case struct {
	name   string
	input  *Input
	expect int
}

func TestFindMin1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{nums: []int{1, 3, 5}}, expect: 1},
		{name: "test 2", input: &Input{nums: []int{2, 2, 2, 0, 1}}, expect: 0},
		{name: "test 3", input: &Input{nums: []int{4, 5, 6, 7, 0, 1, 2}}, expect: 0},
		{name: "test 4", input: &Input{nums: []int{4, 5, 6, 7, 0, 1, 2}}, expect: 0},
		{name: "test 5", input: &Input{nums: []int{4}}, expect: 4},
		{name: "test 6", input: &Input{nums: []int{1, 2, 3}}, expect: 1},
		{name: "test 7", input: &Input{nums: []int{1, 1}}, expect: 1},
		{name: "test 8", input: &Input{nums: []int{3, 3, 1, 3}}, expect: 1},
		{name: "test 9", input: &Input{nums: []int{1, 3, 3}}, expect: 1},
	}
	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := FindMin1(c.input.nums)
			t.Logf("case %s success expect: %v, output: %v, with input: %+v", c.name, c.expect, out, c.input)
		})
	}
}
