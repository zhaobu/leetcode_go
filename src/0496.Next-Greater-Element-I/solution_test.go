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
	nums1 []int
	nums2 []int
}

type Case struct {
	name   string
	input  *Input
	expect []int
}

func TestNextGreaterElement1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{nums1: []int{10, 5, 7, 2}, nums2: []int{10, 5, 8, 7, 9, 2}}, expect: []int{-1, 8, 9, -1}},
		{name: "test 2", input: &Input{nums1: []int{4, 1, 2}, nums2: []int{1, 3, 4, 2}}, expect: []int{-1, 3, -1}},
		{name: "test 3", input: &Input{nums1: []int{2, 4}, nums2: []int{1, 2, 3, 4}}, expect: []int{3, -1}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := nextGreaterElement1(c.input.nums1, c.input.nums2)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}

func TestNextGreaterElement2(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{nums1: []int{10, 5, 7, 9}, nums2: []int{10, 5, 7, 6, 8, 9}}, expect: []int{-1, 7, 8, -1}},
		{name: "test 2", input: &Input{nums1: []int{4, 1, 2}, nums2: []int{1, 3, 4, 2}}, expect: []int{-1, 3, -1}},
		{name: "test 3", input: &Input{nums1: []int{2, 4}, nums2: []int{1, 2, 3, 4}}, expect: []int{3, -1}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := nextGreaterElement2(c.input.nums1, c.input.nums2)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}

func TestNextGreaterElement3(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{nums1: []int{10, 5, 7, 2}, nums2: []int{10, 5, 8, 7, 9, 2}}, expect: []int{-1, 8, 9, -1}},
		{name: "test 2", input: &Input{nums1: []int{4, 1, 2}, nums2: []int{1, 3, 4, 2}}, expect: []int{-1, 3, -1}},
		{name: "test 3", input: &Input{nums1: []int{2, 4}, nums2: []int{1, 2, 3, 4}}, expect: []int{3, -1}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := nextGreaterElement3(c.input.nums1, c.input.nums2)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, out, c.input)
		})
	}
}
