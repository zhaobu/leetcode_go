package main

import (
	. "leetcode/check"
)

/*
 *
 *
 * 剑指 Offer 06. 从尾到头打印链表
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
 解法1 双指针
*/
func reversePrint(head *ListNode) []int {
	ret := []int{}
	for head != nil {
		ret = append(ret, head.Val)
		head = head.Next
	}
	for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
		ret[i], ret[j] = ret[j], ret[i]
	}
	return ret
}

// @lc code=end
