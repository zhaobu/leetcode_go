package main

import (
	. "leetcode/check"
)

/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
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
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	help := &ListNode{}
	preHead, cur := help, head
	for cur != nil {
		curStart := cur //保存当前节点为第一个节点
		//尝试找到最后一个val与cur相同的最后一个节点
		for cur.Next != nil && cur.Next.Val == cur.Val {
			cur = cur.Next
		}
		//如果curStart == cur 说明和curStart值相同的节点只有一个
		if curStart == cur {
			preHead.Next = curStart
			preHead = preHead.Next
		}
		//继续查找下一个
		cur = cur.Next
	}
	//遍历完链表后必须断开最后一个节点
	preHead.Next = nil
	return help.Next
}

// @lc code=end
