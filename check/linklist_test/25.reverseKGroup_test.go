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
		pre     = &ListNode{Next: head} //已经翻转的链表尾结点
		resHead *ListNode               //要返回的已经翻转的链表
		// curHead *ListNode               //本次翻转的链表头结点
	)

	for head != nil {
		var newTail *ListNode //本次k个节点翻转后的尾结点
		pre.Next, newTail = swapList(head, k)
		if resHead == nil { //第一次翻转后的k个节点的head
			resHead = pre.Next
		}
		pre = newTail   //已经翻转的链表尾节点后移
		head = pre.Next //继续翻转剩下的节点
		fmt.Printf("resHead:%s\n", resHead.Print())
	}
	return resHead
}

/*
newHead:翻转后的部分的头节点
newTail:翻转后的部分的尾节点
*/
func swapList(head *ListNode, n int) (newHead, newTail *ListNode) {
	k := 0 //记录本次要反转的节点个数
	newHead = head
	for newHead != nil && k < n { //剩余链表节点个数小于n个或者超过n个都要结束
		k++
		newTail = newHead
		newHead = newHead.Next
	}
	if k < n { //剩余链表节点个数小于n,本次不翻转
		return head, newTail
	}

	//翻转n个链表
	newHead = nil
	newTail = head //翻转后,头结点就会变为尾节点
	for k > 0 {    //剩余节点个数>=n个时就翻转这n个节点
		curNext := head.Next //记录下一个节点
		head.Next = newHead  //当前节点指向已翻转部分的head
		newHead = head       //已翻转部分的head前移
		head = curNext       //翻转下一个节点
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
