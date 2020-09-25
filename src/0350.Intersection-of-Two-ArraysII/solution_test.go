package Solution

import (
	"leetcode/utils/zaplog"
	"testing"
)

func init() {
	zaplog.InitLog(&zaplog.Config{
		Logdir:   "./",
		LogName:  "main.log",
		Loglevel: "debug",
		Debug:    true},
	)
}

//定义结构
type Input struct {
	Nums1 []int
	Nums2 []int
}

type Case struct {
	name   string
	input  *Input
	expect []int
}

func TestIntersect1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{Nums1: []int{1, 2, 2, 1}, Nums2: []int{2, 2}}, expect: []int{2, 2}},
		{name: "test 1", input: &Input{Nums1: []int{4, 9, 5, 4, 4}, Nums2: []int{9, 4, 9, 8, 4, 6}}, expect: []int{9, 4, 4}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			output := Intersect1(c.input.Nums1, c.input.Nums2)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, output, c.input)
		})
	}
}

func TestIntersect2(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{Nums1: []int{1, 2, 2, 1}, Nums2: []int{2, 2}}, expect: []int{2, 2}},
		{name: "test 1", input: &Input{Nums1: []int{4, 9, 5, 4, 4}, Nums2: []int{9, 4, 9, 8, 4, 6}}, expect: []int{9, 4, 4}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			output := Intersect2(c.input.Nums1, c.input.Nums2)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, output, c.input)
		})
	}
}
