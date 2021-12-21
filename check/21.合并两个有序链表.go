package check

/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var (
		head = &ListNode{}
		res  = head
	)
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			head.Next = list1
			head = list1
			list1 = list1.Next
		} else {
			head.Next = list2
			head = list2
			list2 = list2.Next
		}
	}
	if list1 != nil {
		head.Next = list1
	} else if list2 != nil {
		head.Next = list2
	}
	return res.Next
}

// @lc code=end
