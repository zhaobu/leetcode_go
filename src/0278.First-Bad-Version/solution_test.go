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
	n     int
	start int
}

func (it *Input) Init() []int {
	res := make([]int, it.n)
	for i, _ := range res {
		//start 表示从版本start开始出错
		if i+1 >= it.start {
			res[i] = 1
		}
	}
	return res
}

type Case struct {
	name   string
	input  *Input
	expect int
}

func TestLengthOfLastWord1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{n: 5, start: 4}, expect: 4},
		{name: "test 2", input: &Input{n: 4, start: 2}, expect: 2},
		{name: "test 3", input: &Input{n: 7, start: 1}, expect: 1},
		{name: "test 4", input: &Input{n: 11, start: 11}, expect: 11},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			Versions = c.input.Init()
			out := FirstBadVersion1(c.input.n)
			t.Logf("case %s success expect: %v, output: %v, with input: %+v", c.name, c.expect, out, c.input)
		})
	}
}

func TestLengthOfLastWord2(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{n: 5, start: 4}, expect: 4},
		{name: "test 2", input: &Input{n: 4, start: 2}, expect: 2},
		{name: "test 3", input: &Input{n: 7, start: 1}, expect: 1},
		{name: "test 4", input: &Input{n: 11, start: 11}, expect: 11},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			Versions = c.input.Init()
			out := FirstBadVersion2(c.input.n)
			t.Logf("case %s success expect: %v, output: %v, with input: %+v", c.name, c.expect, out, c.input)
		})
	}
}
