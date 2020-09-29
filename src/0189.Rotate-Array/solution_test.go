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
	k    int
}

func (i *Input) Copy() *Input {
	res := &Input{
		nums: make([]int, len(i.nums)),
		k:    i.k,
	}
	copy(res.nums, i.nums)
	return res
}

type Case struct {
	name   string
	input  *Input
	expect []int
}

func TestRotate1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 3}, expect: []int{5, 6, 7, 1, 2, 3, 4}},
		{name: "test 2", input: &Input{nums: []int{-1, -100, 3, 99}, k: 2}, expect: []int{3, 99, -1, -100}},
		{name: "test 2", input: &Input{nums: []int{1, 2}, k: 3}, expect: []int{2, 1}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := c.input.Copy()
			Rotate1(c.input.nums, c.input.k)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, c.input.nums, input)
		})
	}
}

func TestRotate2(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 3}, expect: []int{5, 6, 7, 1, 2, 3, 4}},
		{name: "test 2", input: &Input{nums: []int{-1, -100, 3, 99}, k: 2}, expect: []int{3, 99, -1, -100}},
		{name: "test 2", input: &Input{nums: []int{1, 2}, k: 3}, expect: []int{2, 1}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := c.input.Copy()
			Rotate2(c.input.nums, c.input.k)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, c.input.nums, input)
		})
	}
}

func TestRotate3(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 3}, expect: []int{5, 6, 7, 1, 2, 3, 4}},
		{name: "test 2", input: &Input{nums: []int{-1, -100, 3, 99}, k: 2}, expect: []int{3, 99, -1, -100}},
		{name: "test 2", input: &Input{nums: []int{1, 2}, k: 3}, expect: []int{2, 1}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := c.input.Copy()
			Rotate3(c.input.nums, c.input.k)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, c.input.nums, input)
		})
	}
}

func TestRotate4(t *testing.T) {
	//	测试用例
	cases := []*Case{
		// {name: "test 1", input: &Input{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 1}, expect: []int{7, 1, 2, 3, 4, 5, 6}},
		// {name: "test 1", input: &Input{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 2}, expect: []int{6, 7, 1, 2, 3, 4, 5}},
		// {name: "test 1", input: &Input{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 3}, expect: []int{5, 6, 7, 1, 2, 3, 4}},
		// {name: "test 1", input: &Input{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 4}, expect: []int{4, 5, 6, 7, 1, 2, 3}},
		// {name: "test 1", input: &Input{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 5}, expect: []int{3, 4, 5, 6, 7, 1, 2}},
		// {name: "test 1", input: &Input{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 6}, expect: []int{2, 3, 4, 5, 6, 7, 1}},
		// {name: "test 1", input: &Input{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 7}, expect: []int{1, 2, 3, 4, 5, 6, 7}},
		// {name: "test 2", input: &Input{nums: []int{-1, -100, 3, 99, 88, 66}, k: 1}, expect: []int{66, -1, -100, 3, 99, 88}},
		{name: "test 2", input: &Input{nums: []int{-1, -100, 3, 99, 88, 66}, k: 2}, expect: []int{88, 66, -1, -100, 3, 99}},
		{name: "test 2", input: &Input{nums: []int{-1, -100, 3, 99, 88, 66}, k: 3}, expect: []int{99, 88, 66, -1, -100, 3}},
		{name: "test 2", input: &Input{nums: []int{-1, -100, 3, 99, 88, 66}, k: 4}, expect: []int{3, 99, 88, 66, -1, -100}},
		{name: "test 2", input: &Input{nums: []int{-1, -100, 3, 99, 88, 66}, k: 5}, expect: []int{-100, 3, 99, 88, 66, -1}},
		{name: "test 2", input: &Input{nums: []int{-1, -100, 3, 99, 88, 66}, k: 6}, expect: []int{-1, -100, 3, 99, 88, 66}},
		{name: "test 2", input: &Input{nums: []int{1, 2}, k: 3}, expect: []int{2, 1}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := c.input.Copy()
			Rotate4(c.input.nums, c.input.k)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, c.input.nums, input)
		})
	}
}
