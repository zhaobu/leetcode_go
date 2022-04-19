package main

import (
	. "leetcode/check"
)

/*
 * @lc app=leetcode.cn id=147 lang=golang
 *
 * [147] 对链表进行插入排序
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
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	help := &ListNode{Next: head}
	pre := head                                        //当前排序节点的前驱
	for cur := head.Next; cur != nil; cur = cur.Next { //从head.Next开始进行排序
		if pre.Val > cur.Val {
			//先找到当前节点应该要插入的位置的前驱
			p := help //从help开始找到第一个下一个节点>cur的节点也就是cur要插入的位置
			for ; p.Next != cur && p.Next.Val <= cur.Val; p = p.Next {
			}
			//断开当前节点,当前节点的前驱指向后继
			pre.Next = cur.Next
			//移动当前节点到恰当位置
			cur.Next = p.Next
			p.Next = cur
		}
		pre = cur
	}
	return help.Next
}

// @lc code=end
