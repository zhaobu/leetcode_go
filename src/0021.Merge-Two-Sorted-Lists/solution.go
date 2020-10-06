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

func (l *ListNode) Copy() *ListNode {
	res := &ListNode{}
	cur := res
	for i := l; i != nil; i = i.Next {
		cur.Next = &ListNode{Val: i.Val}
		cur = cur.Next
	}
	return res.Next
}

/*
 */
func MergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	result := &ListNode{}
	head := result
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
	}
	if l1 != nil {
		head.Next = l1
	} else if l2 != nil {
		head.Next = l2
	}
	return result.Next
}
