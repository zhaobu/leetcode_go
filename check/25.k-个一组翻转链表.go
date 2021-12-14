/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
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
		pre = &ListNode{Next: head} //已经翻转的链表尾结点
		// curHead *ListNode               //本次翻转的链表头结点
		resHead *ListNode //要返回的已经翻转的链表
	)

	for head != nil {
		var newTail *ListNode
		pre.Next, newTail = swapList(head, k)
		if resHead == nil {
			resHead = pre.Next
		}
		pre = newTail
		head = pre.Next
	}
	return resHead
}

/*
newHead:翻转后的部分的头节点
newTail:翻转后的部分的尾节点
*/
func swapList(head *ListNode, n int) (newHead, newTail *ListNode) {
	k := 0
	newHead = head
	for newHead != nil {
		k++
		newHead = newHead.Next
		if k >= n {
			break
		}
	}
	if k < n {
		return head, nil
	}
	//翻转n个链表

	newHead = nil
	newTail = head //翻转后,头结点就会变为尾节点

	for head != nil {
		curNext := head.Next
		head.Next = newHead
		newHead = head
		head = curNext
	}
	newTail.Next = head //把翻转的链表和还未翻转的连接
	return
}

// @lc code=end

