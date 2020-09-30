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
	digits []int
}

type Case struct {
	name   string
	input  *Input
	expect []int
}

func TestPlusOne1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{digits: []int{1, 2, 3}}, expect: []int{1, 2, 4}},
		{name: "test 2", input: &Input{digits: []int{4, 3, 2, 1}}, expect: []int{4, 3, 2, 2}},
		{name: "test 2", input: &Input{digits: []int{4, 9, 9, 9}}, expect: []int{5, 0, 0, 0}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := PlusOne1(c.input.digits)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}

func TestPlusOne2(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{digits: []int{1, 2, 3}}, expect: []int{1, 2, 4}},
		{name: "test 2", input: &Input{digits: []int{4, 3, 2, 1}}, expect: []int{4, 3, 2, 2}},
		{name: "test 2", input: &Input{digits: []int{4, 9, 9, 9}}, expect: []int{5, 0, 0, 0}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := PlusOne2(c.input.digits)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}

func TestPlusOne3(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{digits: []int{1, 2, 3}}, expect: []int{1, 2, 4}},
		{name: "test 2", input: &Input{digits: []int{4, 3, 2, 1}}, expect: []int{4, 3, 2, 2}},
		{name: "test 2", input: &Input{digits: []int{4, 9, 9, 9}}, expect: []int{5, 0, 0, 0}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := PlusOne3(c.input.digits)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}

func TestPlusOne4(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{digits: []int{1, 2, 3}}, expect: []int{1, 2, 4}},
		{name: "test 2", input: &Input{digits: []int{4, 3, 2, 1}}, expect: []int{4, 3, 2, 2}},
		{name: "test 2", input: &Input{digits: []int{4, 9, 9, 9}}, expect: []int{5, 0, 0, 0}},
		{name: "test 2", input: &Input{digits: []int{9, 9, 9}}, expect: []int{1, 0, 0, 0}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := PlusOne4(c.input.digits)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}
