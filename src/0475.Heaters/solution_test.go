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
	houses  []int
	heaters []int
}

type Case struct {
	name   string
	input  *Input
	expect int
}

func TestFindRadius1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{houses: []int{1, 2, 3}, heaters: []int{2}}, expect: 1},
		{name: "test 2", input: &Input{houses: []int{1, 2, 3, 4}, heaters: []int{1, 4}}, expect: 1},
		{name: "test 3", input: &Input{houses: []int{1, 5}, heaters: []int{2}}, expect: 3},
	}
	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := FindRadius1(c.input.houses, c.input.heaters)
			t.Logf("case %s success expect: %v, output: %v, with input: %+v", c.name, c.expect, out, c.input)
		})
	}
}
