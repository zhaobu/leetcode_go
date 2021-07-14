package _6_linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
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

	if fast != nil {
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
