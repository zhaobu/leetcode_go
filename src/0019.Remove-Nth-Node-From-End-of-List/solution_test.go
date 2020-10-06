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
	head *ListNode
	n    int
}

func (i *Input) copy() *Input {
	result := &ListNode{}
	var (
		pre     = result
		newNode *ListNode
	)
	for j := i.head; j != nil; j = j.Next {
		newNode = &ListNode{Val: j.Val}
		pre.Next = newNode
		pre = newNode
	}
	return &Input{
		n:    i.n,
		head: result.Next,
	}
}

type Case struct {
	name   string
	input  *Input
	expect *ListNode
}

func TestRemoveNthFromEnd1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1",
			input:  &Input{head: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}}, n: 1},
			expect: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}},
		},
		{name: "test 2",
			input:  &Input{head: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}}, n: 2},
			expect: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: nil}}}}},
		},
		{name: "test 3",
			input:  &Input{head: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}}, n: 6},
			expect: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}},
		},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			copy := c.input.copy()
			out := RemoveNthFromEnd1(c.input.head, c.input.n)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect.Print(), out.Print(), copy.head.Print())
		})
	}
}
