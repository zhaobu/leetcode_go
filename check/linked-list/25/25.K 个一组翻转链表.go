package main

/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 *
 * https://leetcode-cn.com/problems/reverse-nodes-in-k-group/description/
 *
 * algorithms
 * Hard (65.81%)
 * Likes:    1397
 * Dislikes: 0
 * Total Accepted:    253.7K
 * Total Submissions: 385.3K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
 *
 * k 是一个正整数，它的值小于或等于链表的长度。
 *
 * 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
 *
 * 进阶：
 *
 *
 * 你可以设计一个只使用常数额外空间的算法来解决此问题吗？
 * 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,3,4,5], k = 2
 * 输出：[2,1,4,3,5]
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = [1,2,3,4,5], k = 3
 * 输出：[3,2,1,4,5]
 *
 *
 * 示例 3：
 *
 *
 * 输入：head = [1,2,3,4,5], k = 1
 * 输出：[1,2,3,4,5]
 *
 *
 * 示例 4：
 *
 *
 * 输入：head = [1], k = 1
 * 输出：[1]
 *
 *
 *
 *
 *
 * 提示：
 *
 *
 * 列表中节点的数量在范围 sz 内
 * 1
 * 0
 * 1
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
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 1 {
		return head
	}

	var (
		pre     = &ListNode{Next: head} //已经翻转的链表尾结点
		resHead *ListNode               //要返回的已经翻转的链表
		// curHead *ListNode               //本次翻转的链表头结点
	)

	for head != nil {
		var newTail *ListNode //本次k个节点翻转后的尾结点
		pre.Next, newTail = reverseOnce(head, k)
		if resHead == nil { //第一次翻转后的k个节点的head
			resHead = pre.Next
		}
		pre = newTail   //已经翻转的链表尾节点后移
		head = pre.Next //继续翻转剩下的节点
	}
	return resHead
}

/*
newHead:翻转后的部分的头节点
newTail:翻转后的部分的尾节点
*/
func reverseOnce(head *ListNode, n int) (newHead, newTail *ListNode) {
	k := 0 //记录本次要反转的节点个数
	newHead = head
	for k < n { //剩余链表节点个数>=n个
		if newHead == nil { //剩余链表节点个数<n个
			return head, newTail
		}
		k++
		newTail = newHead
		newHead = newHead.Next
	}

	//翻转n个链表
	newHead = nil
	newTail = head //翻转后,头结点就会变为尾节点
	for k > 0 {    //剩余节点个数>=n个时就翻转这n个节点
		curNext := head.Next //记录下一个节点
		head.Next = newHead  //当前节点指向已翻转部分的head
		newHead = head       //已翻转部分的head前移
		head = curNext       //翻转下一个节点
		k--                  //剩余待反转的节点-1
	}
	newTail.Next = head //把翻转的链表和还未翻转的连接
	return
}

// 递归
func reverseKGroup1(head *ListNode, k int) *ListNode {
	node, cnt := head, 0 //node用来遍历,看剩下节点个数是否>=k
	for cnt < k {        //从0到k-1一共k个节点
		if node == nil { //如果不足k个要提前返回
			return head
		}
		node = node.Next
		cnt++
	}

	/*
		1.走到这里cnt必定=k
		2.下一次要反转的就是从node开始的k个节点
		3.cnt从0到k,也就表明head从指向第0个节点到指向了第k个节点,也就是下一次要翻转的head
	*/
	prev := reverseKGroup(node, k) //翻转剩下的部分,返回值指向已经翻转的部分的head
	for cnt > 0 {                  //cnt从k减小到1,恰好反转k个节点
		next := head.Next //记录下一个节点
		head.Next = prev  //当前节点指向已翻转部分的head
		prev = head       //已翻转部分的head前移
		head = next       //翻转下一个节点
		cnt--             //剩余待反转的节点-1
	}

	return prev
}

// @lc code=end
