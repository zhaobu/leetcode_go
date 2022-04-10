package main

import (
	. "leetcode/check"
)

/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
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
 解法1 二叉堆
*/
func mergeKLists1(lists []*ListNode) *ListNode {
	m := len(lists)
	if m == 0 {
		return nil
	}
	if m == 1 {
		return lists[0]
	}

	//构建大小为k的小顶堆
	heap := make([]*ListNode, 0, m)

	heapify := func(start int) {
		n := len(heap)
		for {
			minPos := start
			left, right := 2*start+1, 2*start+2
			if left < n && heap[left].Val <= heap[minPos].Val {
				minPos = left
			}
			if right < n && heap[right].Val <= heap[minPos].Val {
				minPos = right
			}
			if minPos == start {
				break
			}
			heap[start], heap[minPos] = heap[minPos], heap[start]
			start = minPos
		}
	}

	//初始化堆中元素
	for i := range lists {
		if lists[i] != nil {
			heap = append(heap, lists[i])
			// lists[i] = lists[i].Next
		}
	}
	//从最后一个非叶子节点开始自顶向下堆化
	for i := (len(heap) >> 1) - 1; i >= 0; i-- {
		heapify(i)
	}

	head := &ListNode{}
	last := head //已经合并好的链表的最后一个节点
	for len(heap) > 0 {
		top := heap[0]
		last.Next = top
		last = last.Next

		if top.Next != nil {
			//如果当前链表不为空,继续用当前节点下一个替换堆顶,相当于删除堆顶元素
			heap[0] = top.Next
		} else {
			//最后一个元素和对顶元素交换,然后删除最后一个元素
			heap[0], heap[len(heap)-1] = heap[len(heap)-1], heap[0]
			heap = heap[:len(heap)-1]
			//此处不能写成heap = heap[1:],必须是最后一个元素和堆顶交换,然后删除最后一个元素
		}
		heapify(0)
	}

	return head.Next
}

/*
解法2 二叉堆
另一种写法
*/
func mergeKLists(lists []*ListNode) *ListNode {
	m := len(lists)
	if m == 0 {
		return nil
	}
	if m == 1 {
		return lists[0]
	}

	//构建大小为k的小顶堆
	heap := make([]*ListNode, 0, m)

	heapify := func(start int) {
		n := len(heap)
		for {
			minPos := start
			left, right := 2*start+1, 2*start+2
			if left < n && heap[left].Val <= heap[minPos].Val {
				minPos = left
			}
			if right < n && heap[right].Val <= heap[minPos].Val {
				minPos = right
			}
			if minPos == start {
				break
			}
			heap[start], heap[minPos] = heap[minPos], heap[start]
			start = minPos
		}
	}

	insert := func(v *ListNode) {
		heap = append(heap, v)
		n := len(heap)
		i := n - 1             //从最后一个节点开始
		parent := (i - 1) >> 1 //父节点
		for parent >= 0 && heap[i].Val <= heap[parent].Val {
			heap[i], heap[parent] = heap[parent], heap[i]
			i = parent
			parent = (i - 1) >> 1
		}
	}

	//初始化堆中元素
	for i := range lists {
		if lists[i] != nil {
			insert(lists[i])
			// lists[i] = lists[i].Next
		}
	}

	head := &ListNode{}
	last := head //已经合并好的链表的最后一个节点
	for len(heap) > 0 {
		top := heap[0]
		last.Next = top
		last = last.Next

		if top.Next != nil {
			//如果当前链表不为空,继续用当前节点下一个替换堆顶,相当于删除堆顶元素
			heap[0] = top.Next
		} else {
			//最后一个元素和对顶元素交换,然后删除最后一个元素
			heap[0], heap[len(heap)-1] = heap[len(heap)-1], heap[0]
			heap = heap[:len(heap)-1]
			//此处不能写成heap = heap[1:],必须是最后一个元素和堆顶交换,然后删除最后一个元素
		}
		heapify(0)
	}

	return head.Next
}

// @lc code=end
