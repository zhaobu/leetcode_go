package main

import (
	. "leetcode/check"
)

/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
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
 解法1 先找到left和right然后反转
*/
func reverseBetween1(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}
	help := &ListNode{}
	leftPre, cur := help, head                                //分别表示left前一个位置的节点和当前正在遍历的节点
	leftNode, rightNode := (*ListNode)(nil), (*ListNode)(nil) //要反转的left和right位置的节点
	n := 1                                                    //记录访问到的位置
	//找到要反转的开始节点和结束节点位置
	for cur != nil && (leftNode == nil || rightNode == nil) { //如果left和right位置都找到了就停止
		if n == left { //找到了left
			leftNode = cur
		} else if n < left { //访问left节点之前才需要记录上一个节点
			leftPre.Next = cur
			leftPre = leftPre.Next
		}
		if n == right { //找到了right
			rightNode = cur
		}
		cur = cur.Next
		n++
	}

	//反转链表
	curRevese := leftNode       //当前正在反转的节点
	rightNext := rightNode.Next //right节点右边部分
	pre := rightNext            // 保证反转后的left到right部分链接right右边的部分
	for curRevese != rightNext {
		next := curRevese.Next
		curRevese.Next = pre
		pre = curRevese
		curRevese = next
	}

	//left前的链表连接反转后的left到right部分
	leftPre.Next = pre
	return help.Next
}

/*
解法2 遍历的时候就直接反转
*/
func reverseBetween2(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}
	help := &ListNode{}
	n := 1
	pre, cur := help, head
	preReverse := (*ListNode)(nil) //用来辅助反转链表
	leftNode := (*ListNode)(nil)   //left位置对应的节点,保存下来用来最后和right右边部分连接
	for cur != nil {
		if n < left {
			pre.Next = cur
			pre = pre.Next
			cur = cur.Next
		} else {
			if n == left {
				leftNode = cur
			}
			next := cur.Next
			cur.Next = preReverse
			preReverse = cur
			cur = next
		}

		if n == right { //把三部分连接起来
			pre.Next = preReverse
			leftNode.Next = cur
			break
		}
		n++
	}

	return help.Next
}

/*
解法3 遍历的时候就直接反转
更简洁写法(官方解法2)
*/
func reverseBetween3(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}
	help := &ListNode{}
	n := 1
	pre, cur := help, head
	//提前把pre指向head,因为有可能left=1
	pre.Next = head
	//先找到left节点的前一个
	for cur != nil && n < left {
		pre.Next = cur
		pre = pre.Next
		cur = cur.Next
		n++
	}

	//从left开始反转
	for cur != nil && n < right {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
		n++
	}

	return help.Next
}

// @lc code=end
