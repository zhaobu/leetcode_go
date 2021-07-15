/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var (
		cur *ListNode = head
		pre *ListNode
		tmp *ListNode
	)
	for cur != nil {
		tmp = cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

// @lc code=end

