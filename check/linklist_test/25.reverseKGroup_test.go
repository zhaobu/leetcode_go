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

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 1 {
		return head
	}

	var (
		pre = &ListNode{Next: head} //已经翻转的链表尾结点
		// curHead *ListNode               //本次翻转的链表头结点
		resHead *ListNode //要返回的已经翻转的链表
	)

	for head != nil {
		var newTail *ListNode
		pre.Next, newTail = swapList(head, k)
		if resHead == nil {
			resHead = pre.Next
		}
		pre = newTail
		head = pre.Next
		fmt.Printf("resHead:%s\n", resHead.Print())
	}
	return resHead
}

/*
newHead:翻转后的部分的头节点
newTail:翻转后的部分的尾节点
*/
func swapList(head *ListNode, n int) (newHead, newTail *ListNode) {
	k := 0
	newHead = head
	for newHead != nil {
		k++
		newHead = newHead.Next
		if k >= n {
			break
		}
	}
	if k < n {
		return head, nil
	}
	//翻转n个链表

	newHead = nil
	newTail = head //翻转后,头结点就会变为尾节点

	for head != nil && k > 0 {
		curNext := head.Next
		head.Next = newHead
		newHead = head
		head = curNext
		k--
	}
	newTail.Next = head //把翻转的链表和还未翻转的连接
	return
}

type ReverseKGroup struct {
	input *Input
	name  string
}

func (m *ReverseKGroup) GetInput() interface{} {
	return m.input
}

func (m *ReverseKGroup) GetName() string {
	return m.name
}

// func (m *ReverseKGroup) Copy() CaseIface {
// 	new := &ReverseKGroup{
// 		name:  m.name,
// 		input: m.input.copy(),
// 	}
// 	return new
// }

func (m *ReverseKGroup) Print() string {
	return fmt.Sprintf("name:%s,input:%s", m.name, m.input.Print())
}

func Test1(t *testing.T) {

	//	测试用例
	cases := []CaseIface{
		&ReverseKGroup{name: "test 1", input: &Input{head: UnmarshalListBySlice([]int{1, 2, 3, 4, 5})}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.GetName(), func(t *testing.T) {
			t.Logf("success  input:%+v", c.GetInput().(*Input).head.Print())
			out := reverseKGroup(c.GetInput().(*Input).head, 2)
			t.Logf("success  output:%+v\n", out.Print())
		})
	}
}
