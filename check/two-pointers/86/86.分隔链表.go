/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	left, right := &ListNode{}, &ListNode{} //left串比x小的部分, right串>=x的部分
	cur, p1, p2 := head, left, right
	for cur != nil {
		if cur.Val < x {
			p1.Next = cur
			p1 = cur
		} else {
			p2.Next = cur
			p2 = cur
		}
		last := cur
		cur = cur.Next
		last.Next = nil
	}
	p1.Next = right.Next

	return left.Next
}

// @lc code=end

