package main

import (
	. "leetcode/check"
)

/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
解法1 栈保存
*/
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	help := &ListNode{Next: head}
	stack := []*ListNode{help}
	cur := head
	for cur != nil {
		stack = append(stack, cur)
		cur = cur.Next
	}
	pre := stack[len(stack)-n-1]
	pre.Next = pre.Next.Next
	return help.Next
}

/*
解法2 双指针
help, 1,2,3,4,5
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	help := &ListNode{Next: head}
	fast := help
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	pre := help
	for fast != nil {
		pre = pre.Next
		fast = fast.Next
	}
	pre.Next = pre.Next.Next
	return help.Next
}

// @lc code=end
