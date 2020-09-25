package Solution

import (
	"leetcode/utils/zaplog"
	"reflect"
	"strconv"
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

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs bool
		expect bool
	}{
		{"TestCase", true, true},
		{"TestCase", true, true},
		{"TestCase", false, false},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			ret := Solution(c.inputs)
			if !reflect.DeepEqual(ret, c.expect) {
				zaplog.Fatalf("expected: %v, but got: %v, with inputs: %v", c.expect, ret, c.inputs)
			} else {
				zaplog.Infof("expected: %v, result: %v, with inputs: %v", c.expect, ret, c.inputs)
			}
		})
	}
}

//	压力测试
func BenchmarkSolution(b *testing.B) {
	return
}

//	使用案列
func ExampleSolution() {
	return
}
