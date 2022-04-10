package main

/*
 * @lc app=leetcode.cn id=203 lang=golang
 *
 * [203] 移除链表元素
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
 解法1: 利用辅助结点
*/
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	help := &ListNode{Next: head} //辅助节点
	cur, pre := head, help
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}
	return help.Next
}

// @lc code=end
