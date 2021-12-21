package check

/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func detectCycle(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return nil
// 	}
// 	var (
// 		first  = head //慢指针,每次走1步
// 		second = head //快指针,每次走2步
// 		/*
// 			当快慢指针相遇时,help指针从head开始
// 			出发,必定会与慢指针相遇于环的起始点
// 		*/
// 		help = head
// 	)
// 	for second.Next != nil && second.Next.Next != nil {
// 		first = first.Next
// 		second = second.Next.Next
// 		if first == second {
// 			for first != help {
// 				first = first.Next
// 				help = help.Next
// 			}

// 			return help
// 		}
// 	}
// 	return nil
// }

// hashmap保存
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	record := make(map[*ListNode]bool, 5)
	for head != nil {
		if record[head] {
			return head
		}
		record[head] = true
		head = head.Next
	}
	return nil
}

// @lc code=end
