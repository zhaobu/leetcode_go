package main

import (
	. "leetcode/check"
)

/*
 * @lc app=leetcode.cn id=328 lang=golang
 *
 * [328] 奇偶链表
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
 解法1
*/
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	help1, help2 := &ListNode{}, &ListNode{}
	p1, p2 := help1, help2
	i := 1
	cur := head
	for cur != nil {
		if i&1 == 1 {
			p1.Next = cur
			p1 = p1.Next
		} else {
			p2.Next = cur
			p2 = p2.Next
		}
		next := cur.Next
		cur.Next = nil
		cur = next
		i++
	}
	//奇数部分尾部和偶数部分头部连接
	p1.Next = help2.Next

	return help1.Next
}

// @lc code=end
