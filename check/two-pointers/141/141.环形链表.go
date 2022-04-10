package main

import (
	. "leetcode/check"
)

/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//快慢指针
func hasCycle1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	var (
		first  = head //慢指针,每次走1步
		second = head //快指针,每次走2步
	)
	for second.Next != nil && second.Next.Next != nil {
		first = first.Next
		second = second.Next.Next
		if first == second {
			return true
		}
	}
	return false
}

//hash表
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	record := make(map[*ListNode]bool) //记录所有访问过的节点地址
	for head != nil {
		if record[head] {
			return true
		}
		record[head] = true
		head = head.Next
	}
	return false
}

// @lc code=end
