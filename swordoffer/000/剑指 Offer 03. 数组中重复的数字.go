package main

import (
	. "leetcode/check"
)

/*
 *
 *
 * 剑指 Offer 03. 数组中重复的数字
 */

// @lc code=start

/*
 解法1 双指针
*/
func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast := head
	for i := 1; i <= k; i++ {
		fast = fast.Next
	}

	slow := head
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	/*
	   1,2,3,4,5,6,7
	*/
	return slow
}

// @lc code=end
