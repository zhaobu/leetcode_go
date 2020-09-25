package Solution

import (
	"leetcode/utils/zaplog"
	"math/rand"
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
	Nums   []int
	Target int
}

type Case struct {
	name   string
	input  *Input
	expect []int
}

func TestTwoSum1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{Nums: []int{2, 7, 11, 15}, Target: 9}, expect: []int{0, 1}},
		{name: "test 2", input: &Input{Nums: []int{3, 2, 4}, Target: 6}, expect: []int{1, 2}},
		{name: "test 3", input: &Input{Nums: []int{7, 6, 5, 3, 2, 1, 4, 9, 10}, Target: 17}, expect: []int{0, 8}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			output := TwoSum1(c.input.Nums, c.input.Target)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, output, c.input)
		})
	}
}

func TestTwoSum2(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{Nums: []int{2, 7, 11, 15}, Target: 9}, expect: []int{0, 1}},
		{name: "test 2", input: &Input{Nums: []int{3, 2, 4}, Target: 6}, expect: []int{1, 2}},
		{name: "test 3", input: &Input{Nums: []int{7, 6, 5, 3, 2, 1, 4, 9, 10}, Target: 17}, expect: []int{0, 8}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			output := TwoSum2(c.input.Nums, c.input.Target)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, output, c.input)
		})
	}
}

func TestTwoSum3(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{Nums: []int{2, 7, 11, 15}, Target: 9}, expect: []int{0, 1}},
		{name: "test 2", input: &Input{Nums: []int{3, 2, 4}, Target: 6}, expect: []int{1, 2}},
		{name: "test 3", input: &Input{Nums: []int{7, 6, 5, 3, 2, 1, 4, 9, 10}, Target: 17}, expect: []int{0, 8}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			output := TwoSum3(c.input.Nums, c.input.Target)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, output, c.input)
		})
	}
}

const N = 100

func BenchmarkTwoSum1(b *testing.B) {
	nums := []int{}
	for i := 0; i < N; i++ {
		nums = append(nums, rand.Int())
	}
	nums = append(nums, 7, 2)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSum1(nums, 9)
	}
}

func BenchmarkTwoSum2(b *testing.B) {
	nums := []int{}
	for i := 0; i < N; i++ {
		nums = append(nums, rand.Int())
	}
	nums = append(nums, 7, 2)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSum2(nums, 9)
	}
}
