package Solution

import (
	"fmt"
)

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Print() string {
	res := ""
	for i := l; i != nil; i = i.Next {
		res += fmt.Sprintf("val:%d,next:%x |", i.Val, i.Next)
	}
	return res
}

func (l *ListNode) Copy() *ListNode {
	res := &ListNode{}
	cur := res

	//防止循环链表进入死循环
	vals := make(map[*ListNode]int)
	for i := l; i != nil; i = i.Next {
		_, ok := vals[i]
		if ok {
			break
		}
		vals[i] = 1
		cur.Next = &ListNode{Val: i.Val}
		cur = cur.Next
	}
	return res.Next
}

/*
[迭代法](https://leetcode.wang/leetcode-206-Reverse-Linked-List.html)
*/
func ReverseList1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	pre := (*ListNode)(nil)
	cur := head

	for cur != nil {
		head = head.Next
		cur.Next = pre
		pre = cur
		cur = head
	}

	return pre
}

/*
[递归法](https://leetcode.wang/leetcode-206-Reverse-Linked-List.html)
*/
func ReverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//把剩下部分反转
	newHead := ReverseList2(head.Next)
	//把head节点添加到反转后的链表的尾部
	head.Next.Next = head
	head.Next = nil
	return newHead
}
