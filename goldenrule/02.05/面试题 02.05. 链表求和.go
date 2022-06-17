package main

import (
	. "leetcode/check"
)

/*
题目:[面试题 02.05. 链表求和](https://leetcode.cn/problems/sum-lists-lcci/)

*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	help := &ListNode{}
	pre := help
	last := 0
	for l1 != nil || l2 != nil {
		cur := last
		if l1 != nil {
			cur += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			cur += l2.Val
			l2 = l2.Next
		}
		last = cur / 10
		pre.Next = &ListNode{Val: cur % 10}
		pre = pre.Next
	}
	if last > 0 {
		pre.Next = &ListNode{Val: last}
	}
	return help.Next
}
