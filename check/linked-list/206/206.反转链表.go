package main

import (
	. "leetcode/check"
)

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

func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var (
		pre  *ListNode //记录已经反转的链表
		next = head    //保存剩下未反转的链表
		cur  = head    //从第一个
	)
	for cur != nil {
		next = cur.Next //先保存剩下未反转的,防止丢失
		cur.Next = pre  //反转
		pre = cur       //记录最新的反转链表
		cur = next      //进行下一次反转
	}
	return pre //返回最新的反转链表
}

func reverseList2(head *ListNode) *ListNode {
	var newHead *ListNode //已经反转的链表
	for head != nil {
		next := head.Next
		head.Next = newHead
		newHead = head
		head = next
	}
	return newHead
}

// 递归
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var reverseListInt func(head, newHead *ListNode) *ListNode
	//head表示当前正在反转的节点,newHead表示已经反转的节点的头部
	reverseListInt = func(head, newHead *ListNode) *ListNode {
		if head == nil {
			return newHead
		}
		next := head.Next                 //先记录下一个要反转的节点
		head.Next = newHead               //当前节点进行反转
		return reverseListInt(next, head) //反转下一个节点
	}
	return reverseListInt(head, nil)
}

// @lc code=end
