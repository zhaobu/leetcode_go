package Solution

import (
	"fmt"
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
	l1 *ListNode
	l2 *ListNode
}

func (i *Input) copy() *Input {
	return &Input{
		l1: i.l1.Copy(),
		l2: i.l2.Copy(),
	}
}

func (i *Input) Print() string {
	return fmt.Sprintf("l1:%s,l2:%s", i.l1.Print(), i.l2.Print())
}

type Case struct {
	name   string
	input  *Input
	expect *ListNode
}

func TestMergeTwoLists1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1",
			input: &Input{
				l1: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}},
				l2: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}},
			},
			expect: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: nil}}}}}},
		},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			copy := c.input.copy()
			out := MergeTwoLists1(c.input.l1, c.input.l2)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect.Print(), out.Print(), copy.Print())
		})
	}
}
