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
	head *ListNode
}

func (i *Input) copy() *Input {
	return &Input{
		head: i.head.Copy(),
	}
}

func (i *Input) Print() string {
	return fmt.Sprintf("head:%s", i.head.Print())
}

type Case struct {
	name   string
	input  *Input
	expect *ListNode
}

func getRandList(ReverseList bool, len int) *ListNode {

	head := UnmarshalListByRand(100, len)
	//	生成没有环的链表
	if !ReverseList {
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

func TestReverseList1(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1",
			input: &Input{
				head: UnmarshalListBySlice([]int{2, 4, 3, 1}),
			},
			expect: UnmarshalListBySlice([]int{1, 3, 4, 2}),
		},
		{name: "test 2",
			input: &Input{
				head: UnmarshalListBySlice([]int{1, 2, 3, 4}),
			},
			expect: UnmarshalListBySlice([]int{4, 3, 2, 1}),
		},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			copy := c.input.copy()
			out := ReverseList1(c.input.head)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect.Print(), out.Print(), copy.Print())
		})
	}
}

func TestReverseList2(t *testing.T) {
	//	测试用例
	cases := []*Case{
		{name: "test 1",
			input: &Input{
				head: UnmarshalListBySlice([]int{2, 4, 3, 1}),
			},
			expect: UnmarshalListBySlice([]int{1, 3, 4, 2}),
		},
		{name: "test 2",
			input: &Input{
				head: UnmarshalListBySlice([]int{1, 2, 3, 4}),
			},
			expect: UnmarshalListBySlice([]int{4, 3, 2, 1}),
		},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			copy := c.input.copy()
			out := ReverseList2(c.input.head)
			t.Logf("success expect: %v, output: %v, with input: %+v", c.expect.Print(), out.Print(), copy.Print())
		})
	}
}
