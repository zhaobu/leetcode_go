package Solution

import (
	"fmt"
)

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Print() string {
	res := ""
	for i := l; i != nil; i = i.Next {
		res += fmt.Sprintf("val:%d,next:%x |", i.Val, i.Next)
	}
	return res
}

func (l *ListNode) Copy() *ListNode {
	res := &ListNode{}
	cur := res

	//防止循环链表进入死循环
	vals := make(map[*ListNode]int)
	for i := l; i != nil; i = i.Next {
		_, ok := vals[i]
		if ok {
			break
		}
		vals[i] = 1
		cur.Next = &ListNode{Val: i.Val}
		cur = cur.Next
	}
	return res.Next
}

/*
[leetcode官方 方法一：哈希表](https://leetcode-cn.com/problems/linked-list-cycle-ii/solution/huan-xing-lian-biao-ii-by-leetcode-solution/)
*/
func DetectCycle1(head *ListNode) *ListNode {
	record := make(map[*ListNode]int)
	for i := 0; head != nil; i++ {
		if _, ok := record[head]; ok {
			return head
		}
		record[head] = i
		head = head.Next
	}
	return nil
}

/*
[leetcode官方 方法二：快慢指针](https://leetcode-cn.com/problems/linked-list-cycle-ii/solution/huan-xing-lian-biao-ii-by-leetcode-solution/)
*/
func DetectCycle2(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil {
		if fast.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
		//快慢指针相遇
		if slow == fast {
			meet := fast
			for meet != head {
				head = head.Next
				meet = meet.Next
			}
			return meet
		}
	}
	return nil
}
