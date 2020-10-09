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

 */
func AddTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: -1, Next: nil}
	tmp := head
	//记录l1的值,l2的值,本次相加之和,本次相加之和的进位数,本次相加之和的个位数字
	n1, n2, sum, a1, a2 := 0, 0, 0, 0, 0
	for l1 != nil || l2 != nil {
		n1, n2 = 0, 0
		//记录l1的值
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		//记录l2的值
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		//计算l1+l2+进位
		sum = n1 + n2 + a1
		a1 = sum / 10
		a2 = sum % 10
		tmp.Next = &ListNode{Val: a2, Next: nil}
		tmp = tmp.Next
	}
	//如果最后一次相加有进位,就记录
	if a1 > 0 {
		tmp.Next = &ListNode{Val: a1, Next: nil}
	}
	return head.Next
}

func AddTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: -1, Next: nil}
	tmp := head
	sum := 0
	for l1 != nil || l2 != nil || sum != 0 {
		//记录l1的值
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		//记录l2的值
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		tmp.Next = &ListNode{Val: sum % 10, Next: nil}
		tmp = tmp.Next
		sum = sum / 10
	}
	return head.Next
}
