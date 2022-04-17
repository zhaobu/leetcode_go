package main

import (
	"fmt"
	. "leetcode/check"
)

/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
解法1 归并排序自顶向下

*/
func sortList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	merge := func(start, end *ListNode) *ListNode {
		if start == nil {
			return end
		}
		if end == nil {
			return start
		}
		help := &ListNode{}
		cur := help
		p1, p2 := start, end

		for p1 != nil && p2 != nil {
			if p1.Val <= p2.Val {
				cur.Next = p1
				p1 = p1.Next
			} else {
				cur.Next = p2
				p2 = p2.Next
			}
			cur = cur.Next
		}
		if p1 != nil {
			cur.Next = p1
		} else if p2 != nil {
			cur.Next = p2
		}
		return help.Next
	}

	var sort func(start *ListNode) *ListNode
	sort = func(start *ListNode) *ListNode {
		if start == nil || start.Next == nil {
			return start
		}

		//找到中心节点
		slow, fast := start, start.Next
		for fast != nil && fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
		}
		/*
			1. 如果节点为奇数个slow为最中间那个节点
			2. 如果节点为偶数个slow为最中间2个节点的第一个
			3. slow把链表分为前后2部分,前面一部分最后一个节点为slow,后面一部分第一个节点为slow.Next
		*/
		mid := slow.Next //先保存后半部分的头节点
		slow.Next = nil  //把前后2部分断开
		left := sort(start)
		right := sort(mid)
		return merge(left, right)
	}

	return sort(head)
}

/*
解法2 归并排序自底向上
*/
func sortList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	//合并2个有序链表
	merge := func(start, end *ListNode) *ListNode {
		if start == nil {
			return end
		}
		if end == nil {
			return start
		}
		help := &ListNode{}
		cur := help
		p1, p2 := start, end

		for p1 != nil && p2 != nil {
			if p1.Val <= p2.Val {
				cur.Next = p1
				p1 = p1.Next
			} else {
				cur.Next = p2
				p2 = p2.Next
			}
			cur = cur.Next
		}
		if p1 != nil {
			cur.Next = p1
		} else if p2 != nil {
			cur.Next = p2
		}
		return help.Next
	}

	//先求得链表的长度
	m := 0
	for cur := head; cur != nil; cur = cur.Next {
		m++
	}

	help := &ListNode{Next: head}
	for n := 1; n < m; n <<= 1 {
		/*
			1. 每一轮都是从链表头节点开始排序,只不过是每一轮排完序后,下一轮间隔节点的个数是本轮的2倍
			2. cur每次初始化时必须为help.Next,不能为head,因为排序后头节点可能会变
		*/
		pre, cur := help, help.Next

		for cur != nil {
			head1 := cur                 //待合并的第1部分头节点
			pre1 := &ListNode{Next: cur} //第1部分最后一个节点的前一个节点
			for i := 0; i < n && cur != nil; i++ {
				pre1.Next = cur
				pre1 = cur
				cur = cur.Next
			}
			pre1.Next = nil //第1部分和后面的断开

			head2 := cur                 //待合并的第2部分头节点
			pre2 := &ListNode{Next: cur} //第2部分最后一个节点的前一个节点
			for i := 0; i < n && cur != nil; i++ {
				pre2.Next = cur
				pre2 = cur
				cur = cur.Next
			}
			pre2.Next = nil //第2部分和后面的断开
			// fmt.Printf("head1:%+v,head2:%+v\n", head1, head2)

			//合并第1和第2部分
			pre.Next = merge(head1, head2)

			//把pre指向已合并部分的最后一个节点
			for pre.Next != nil {
				pre = pre.Next
			}
			// fmt.Printf("head1:%+v,head2:%+v,help=%+v\n", head1, head2, help)
		}

	}
	return help.Next
}

/*
解法3 快速排序交换值
*/
func sortList3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	parition := func(head, tail *ListNode) *ListNode {
		pivot := head.Val //保存分区点
		// fmt.Printf("head=%+v,pivot=%d;", head, pivot)
		/*
			1. left表示最后一个<pivot的节点
			2. 节点区间是左闭右开的,这样方便后面的区间划分为[head,left), [lef.Next,tail)
		*/
		left, cur := head, head.Next
		for cur != tail {
			if cur.Val < pivot {
				left = left.Next
				left.Val, cur.Val = cur.Val, left.Val
			}
			cur = cur.Next
		}
		head.Val, left.Val = left.Val, head.Val //执行这一行代码后,[head,left)<pivot, [left,tail)>=pivot
		return left
	}

	var quickSort func(head, tail *ListNode)
	quickSort = func(head, tail *ListNode) {
		if head == nil || head == tail {
			return
		}
		mid := parition(head, tail)
		// fmt.Printf("head=%+v.left=%d\n", head, left.Val)
		quickSort(head, mid)
		quickSort(mid.Next, tail)
		return
	}

	quickSort(head, nil)
	return head
}

/*
解法3 快速排序指针交换
*/
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var quickSort func(head *ListNode) *ListNode
	quickSort = func(head *ListNode) *ListNode {
		if head == nil || head.Next == nil {
			return head
		}
		fmt.Printf("head=%+v;", head)

		help1, help2 := &ListNode{}, &ListNode{}
		p1, p2, pivot := help1, help2, head
		for cur := pivot; cur != nil; cur = cur.Next {
			if cur.Val <= pivot.Val {
				p1.Next = cur
				p1 = p1.Next
			} else {
				p2.Next = cur
				p2 = p2.Next
			}
		}
		p1.Next = nil
		p2.Next = nil

		fmt.Printf("help1=%+v,help2=%+v,pivot=%+v\n", help1.Next, help2.Next, pivot)
		left := quickSort(help1.Next.Next)
		right := quickSort(help2.Next)

		//把3段连接起来
		cur := left
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = right
		return left
	}

	return quickSort(head)
}

// @lc code=end
