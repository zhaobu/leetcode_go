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
[leetcode官方 方法一：哈希表](https://leetcode-cn.com/problems/linked-list-cycle/solution/huan-xing-lian-biao-by-leetcode/)
*/
func HasCycle1(head *ListNode) bool {
	vals := make(map[*ListNode]int)
	for head != nil {
		_, ok := vals[head]
		if ok {
			return true
		}
		vals[head] = 1
		head = head.Next
	}
	return false
}

/*
[leetcode官方 方法二：双指针](https://leetcode-cn.com/problems/linked-list-cycle/solution/huan-xing-lian-biao-by-leetcode/)
*/
func HasCycle2(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast := head.Next
	for fast != nil && head != nil && fast.Next != nil {
		if head == fast {
			return true
		}
		head = head.Next
		fast = fast.Next.Next
	}
	return false
}
