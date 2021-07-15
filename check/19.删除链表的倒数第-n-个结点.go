/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}
*/

// func removeNthFromEnd(head *ListNode, n int) *ListNode {
// 	if head == nil {
// 		return nil
// 	}
// 	cur := head
// 	//遍历,用map保存所有节点
// 	listMap := make(map[int]*ListNode, 10)
// 	for i := 0; cur != nil; i++ {
// 		listMap[i] = cur
// 		cur = cur.Next
// 	}
// 	if n > len(listMap) {
// 		return nil
// 	} else if n == len(listMap) { //删除第一个节点
// 		head = head.Next
// 	} else { //删除非第一个节点
// 		listMap[len(listMap)-n-1].Next = listMap[len(listMap)-n].Next
// 	}
// 	return head
// }

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	//使用哨兵节点避免为删除head为单节点还有删除第一个节点等情况做特殊处理
	sentry := &ListNode{Next: head}
	//快指针先走n+1步
	fast := sentry
	for i := 0; i <= n && fast != nil; i++ {
		fast = fast.Next
	}
	//快慢指针同时走
	slow := sentry
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	//当快指针到达终点时,快慢指针相差n+1,也就是slow->next指向要删除的节点
	slow.Next = slow.Next.Next
	return sentry.Next
}

// @lc code=end

