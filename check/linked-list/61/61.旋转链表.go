package main

import (
	. "leetcode/check"
)

/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	//计算链表长度
	length := 0
	cur := head
	last := head
	for cur != nil {
		last = cur
		cur = cur.Next
		length++
	}
	k %= length
	if k == 0 {
		return head
	}

	//遍历到第length-k个节点

	cur = head
	count := length - k - 1
	for i := 0; i < count; i++ {
		cur = cur.Next
	}
	//第一部分
	ret := cur.Next

	//断开第二部分最后和第一部分开始
	cur.Next = nil

	//第一部分最后和第二部分开始连接
	last.Next = head

	return ret
}

// @lc code=end
