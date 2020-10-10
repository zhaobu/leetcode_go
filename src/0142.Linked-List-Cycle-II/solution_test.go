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
	expect int
}

// index >-1时添加一个环
func getSliceList(nums []int, index int) (*ListNode, int) {
	if index >= len(nums) {
		panic("index >= len(nums)")
	}
	head := UnmarshalListBySlice(nums)
	//	生成没有环的链表
	if index == -1 {
		return head, -1
	}

	// 生成有环的链表，让尾节点指向序号为index的节点
	tmp := head  //index的节点
	tail := head //尾部节点
	for i := 0; ; i++ {
		if i < index {
			tmp = tmp.Next
		}
		if tail.Next == nil {
			break
		}
		tail = tail.Next
	}
	tail.Next = tmp
	return head, index
}

// 给list添加一个环
func getRandList(hasCycle bool, len int) (*ListNode, int) {

	head := UnmarshalListByRand(100, len)
	//	生成没有环的链表
	if !hasCycle {
		return head, -1
	}

	// 生成有环的链表，让尾节点指向序号为index的节点
	index := rand.Intn(len)
	tmp := head  //index的节点
	tail := head //尾部节点
	for i := 0; ; i++ {
		if i < index {
			tmp = tmp.Next
		}
		if tail.Next == nil {
			break
		}
		tail = tail.Next
	}
	tail.Next = tmp
	return head, index
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

func TestDetectCycle1(t *testing.T) {
	//	测试用例
	head1, expect1 := getSliceList([]int{1, 2, -2, 9, 8}, 3)
	head2, expect2 := getSliceList([]int{1, 2, -2, 9, 8}, -1)
	head3, expect3 := getSliceList([]int{1, 2, -2, 9, 8}, 0)
	cases := []*Case{
		{name: "test 1",
			input: &Input{
				head: head1,
			},
			expect: expect1,
		},
		{name: "test 2",
			input: &Input{
				head: head2,
			},
			expect: expect2,
		},
		{name: "test 3",
			input: &Input{
				head: head3,
			},
			expect: expect3,
		},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			copy := c.input.copy()
			out := DetectCycle1(c.input.head)
			t.Logf("success expect: %v, output: %v,with input: %+v", c.expect, out, copy.Print())
		})
	}
}

func TestDetectCycle2(t *testing.T) {
	//	测试用例
	head1, expect1 := getSliceList([]int{1, 2, -2, 9, 8}, 3)
	head2, expect2 := getSliceList([]int{1, 2, -2, 9, 8}, -1)
	head3, expect3 := getSliceList([]int{1, 2, -2, 9, 8}, 0)
	cases := []*Case{
		{name: "test 1",
			input: &Input{
				head: head1,
			},
			expect: expect1,
		},
		{name: "test 2",
			input: &Input{
				head: head2,
			},
			expect: expect2,
		},
		{name: "test 3",
			input: &Input{
				head: head3,
			},
			expect: expect3,
		},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			copy := c.input.copy()
			out := DetectCycle2(c.input.head)
			t.Logf("success expect: %v, output: %v,with input: %+v", c.expect, out, copy.Print())
		})
	}
}
