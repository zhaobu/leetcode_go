package Solution

import (
	"fmt"
	"math/rand"
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

func getRandList(AddTwoNumbers bool, len int) *ListNode {

	head := UnmarshalListByRand(100, len)
	//	生成没有环的链表
	if !AddTwoNumbers {
		return head
	}

	tmp := head

	// 生成有环的链表，让第index个及节点指向头结点
	index := rand.Intn(len)
	for i := 0; i < index; i++ {
		tmp = tmp.Next
	}

	tmp.Next = head
	return head
}

// 随机初始化链表
func UnmarshalListByRand(max_num int, len int) *ListNode {
	head := &ListNode{Val: -1, Next: nil}
	tmp := head

	for i := 0; i < len; i++ {
		tmp.Next = &ListNode{Val: rand.Intn(max_num), Next: nil}
		tmp = tmp.Next
	}
	return head.Next
}

//	根据数组反序列化链表
func UnmarshalListBySlice(nums []int) *ListNode {
	head := &ListNode{Val: -1, Next: nil}
	tmp := head
	for _, v := range nums {
		tmp.Next = &ListNode{Val: v, Next: nil}
		tmp = tmp.Next
	}
	return head.Next
}

func TestAddTwoNumbers1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1",
			input: &Input{
				l1: UnmarshalListBySlice([]int{2, 4, 3}),
				l2: UnmarshalListBySlice([]int{5, 6, 4}),
			},
			expect: UnmarshalListBySlice([]int{7, 0, 8}),
		},
		{name: "test 2",
			input: &Input{
				l1: UnmarshalListBySlice([]int{9, 5, 9}),
				l2: UnmarshalListBySlice([]int{4, 6, 9}),
			},
			expect: UnmarshalListBySlice([]int{3, 2, 9, 1}),
		},
		{name: "test 3",
			input: &Input{
				l1: UnmarshalListBySlice([]int{9, 5, 9}),
				l2: UnmarshalListBySlice([]int{6, 9}),
			},
			expect: UnmarshalListBySlice([]int{5, 5, 0, 1}),
		},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			copy := c.input.copy()
			out := AddTwoNumbers1(c.input.l1, c.input.l2)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect.Print(), out.Print(), copy.Print())
		})
	}
}

func TestAddTwoNumbers2(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1",
			input: &Input{
				l1: UnmarshalListBySlice([]int{2, 4, 3}),
				l2: UnmarshalListBySlice([]int{5, 6, 4}),
			},
			expect: UnmarshalListBySlice([]int{7, 0, 8}),
		},
		{name: "test 2",
			input: &Input{
				l1: UnmarshalListBySlice([]int{9, 5, 9}),
				l2: UnmarshalListBySlice([]int{4, 6, 9}),
			},
			expect: UnmarshalListBySlice([]int{3, 2, 9, 1}),
		},
		{name: "test 3",
			input: &Input{
				l1: UnmarshalListBySlice([]int{9, 5, 9}),
				l2: UnmarshalListBySlice([]int{6, 9}),
			},
			expect: UnmarshalListBySlice([]int{5, 5, 0, 1}),
		},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			copy := c.input.copy()
			out := AddTwoNumbers2(c.input.l1, c.input.l2)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect.Print(), out.Print(), copy.Print())
		})
	}
}
