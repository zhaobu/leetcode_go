package _6_linkedlist

/**
 * Definition for singly-linked list.
type ListNode struct {
	next  *ListNode
	value interface{}
}
*/
func isPalindrome(head *ListNode) bool {
	var slow *ListNode = head
	var fast *ListNode = head
	var prev *ListNode = nil
	var temp *ListNode = nil

	if head == nil || head.next == nil {
		return true
	}

	for fast != nil && fast.next != nil {
		fast = fast.next.next
		temp = slow.next
		slow.next = prev
		prev = slow
		slow = temp
	} // 快的先跑完,同时反转了一半链表,剪短

	/*
	   运行到此处时:
	   1->2->3->4->5->6: prew=3->2->1, slow=4->5->6, fast=nil
	   1->2->3->4->5: prew=2->1, slow=3->4->5, fast=5

	*/

	if fast != nil { //基数个字符,此时slow指向最中间的节点
		slow = slow.next // 处理余数,跨过中位数
		// prev 增加中 2->1->nil
	}

	var l1 *ListNode = prev
	var l2 *ListNode = slow

	for l1 != nil && l2 != nil && l1.value == l2.value {
		l1 = l1.next
		l2 = l2.next
	}

	return (l1 == nil && l2 == nil)

}
