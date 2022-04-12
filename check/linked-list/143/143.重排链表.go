package main

import (
	. "leetcode/check"
)

/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
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
解法1 链表反转
*/
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	//快慢指针找到中点
	slow, fast := head, head.Next
	slowPre := (*ListNode)(nil) //记录慢指针的前一个节点
	for fast != nil && fast.Next != nil {
		slowPre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil { //偶数个节点
		slowPre = slow
		slow = slow.Next
	}
	slowPre.Next = nil //断开前后部分链表

	//反转后半部分
	cur, pre := slow, (*ListNode)(nil)
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	//前后部分链表重组
	help := &ListNode{}
	p1, p2 := head, pre
	for p1 != nil {
		//连接p1
		help.Next = p1
		p1 = p1.Next

		//连接p2
		help.Next.Next = p2
		p2 = p2.Next

		//继续下一轮
		help = help.Next.Next
	}
}

// @lc code=end
