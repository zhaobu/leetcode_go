package main

import (
	. "leetcode/check"
)

/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
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
解法1 快慢指针
1. 先用快慢指针找到中间位置
2. 反转后半部分
3. 比较反转链表和前半部分
*/
func isPalindrome1(head *ListNode) bool {
	if head.Next == nil {
		return true
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	/*
	   循环结束时
	   如果是偶数个节点: slow指向中间2个节点的第一个, fast != nil
	   如果是奇数个节点: slow指向中间节点, fast == nil
	*/
	// 从slow的下一个节点开始反转剩下的链表
	cur, last := slow.Next, (*ListNode)(nil)
	for cur != nil {
		next := cur.Next //保存当前节点的下一个节点
		cur.Next = last  //反转当前节点
		last = cur       //更新已反转节点的头节点
		cur = next       //继续翻转下一个节点
	}

	/*
	   1. 反转结束试,反转链表的头节点为last
	   2. 反转链表的长度<=前一部分未反转的
	*/
	for last != nil {
		if last.Val != head.Val {
			return false
		}
		last = last.Next
		head = head.Next
	}

	return true
}

/*
解法2 数组
*/
func isPalindrome(head *ListNode) bool {
	if head.Next == nil {
		return true
	}

	//复制到数组
	ret := make([]int, 0, 50000) //根据题目条件预估一下节点数量
	for head != nil {
		ret = append(ret, head.Val)
		head = head.Next
	}
	for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
		if ret[i] != ret[j] {
			return false
		}
	}

	return true
}

// @lc code=end
