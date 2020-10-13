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
	s []byte
}

func (i *Input) Copy() []byte {
	return append([]byte{}, i.s...)
}

type Case struct {
	name   string
	input  *Input
	expect []byte
}

func TestReverseString1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1", input: &Input{s: []byte{'h', 'e', 'l', 'l', 'o'}}, expect: []byte{'o', 'l', 'l', 'e', 'h'}},
		{name: "test 2", input: &Input{s: []byte{'H', 'a', 'n', 'n', 'a', 'h'}}, expect: []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
		{name: "test 3", input: &Input{s: []byte{'H', 'a', 'b', 'c', 'd', 'e'}}, expect: []byte{'e', 'd', 'c', 'b', 'a', 'H'}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ReverseString1(c.input.s)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect, c.input.s, c.input.Copy())
		})
	}
}
