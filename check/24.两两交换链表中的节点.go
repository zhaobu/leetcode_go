/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 *
 * https://leetcode-cn.com/problems/swap-nodes-in-pairs/description/
 *
 * algorithms
 * Medium (70.43%)
 * Likes:    1164
 * Dislikes: 0
 * Total Accepted:    347.6K
 * Total Submissions: 493.5K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,3,4]
 * 输出：[2,1,4,3]
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：head = [1]
 * 输出：[1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中节点的数目在范围 [0, 100] 内
 * 0 <= Node.val <= 100
 *
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func swapPairs(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil { //节点为空或者只有一个节点直接返回
// 		return head
// 	}

// 	var (
// 		res  = head.Next             //最终返回的链表头结点肯定指向head.Next
// 		tail = &ListNode{Next: head} //记录已经交换好的部分的尾节点
// 	)
// 	for head != nil && head.Next != nil {
// 		next := head.Next       //保存下一个节点,用来本次交换
// 		thrid := head.Next.Next //保存下下个节点,用来记录下次交换的位置
// 		head.Next = thrid       //本次交换的第1个节点指向还未交换的部分
// 		next.Next = head        //本次交换的第2个节点指向第一个节点
// 		tail.Next = next        //把已经交换的部分和本次交换的连接
// 		tail = head             //已经交换的部分尾节点前移
// 		head = thrid            //准备进行下次交换
// 	}
// 	return res
// }

//递归
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil { //节点为空或者只有一个节点直接返回
		return head
	}
	newHead := head.Next //先保存下一个节点
	/*
		swapPairs(newHead.Next)看成是后面所有已经交换的部分
		1.先处理第1个节点,交换到2的位置和swapPairs(newHead.Next)连接
	*/
	head.Next = swapPairs(newHead.Next)
	// 2.再处理第2个节点,把他和第1个节点连接
	newHead.Next = head
	return newHead //返回第二个节点
}

// @lc code=end

