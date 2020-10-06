package Solution

import "fmt"

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

/*
 */
func RemoveNthFromEnd1(head *ListNode, n int) *ListNode {
	result := &ListNode{Val: 0, Next: head}
	cur := result
	var pre *ListNode = nil
	count := 0
	for ; head != nil; head = head.Next {
		if count >= n-1 {
			pre = cur
			cur = cur.Next
		}
		count++
	}
	pre.Next = cur.Next
	return result.Next
}
