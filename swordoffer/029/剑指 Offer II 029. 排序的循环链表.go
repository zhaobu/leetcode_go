package main

import (
	. "leetcode/check"
)

/*
 *
 *
 * 剑指 Offer II 029. 排序的循环链表
 */

// @lc code=start

/**
* Definition for a Node.
* type Node struct {
	*     Val int
	*     Next *Node
	* }
*/

/*
 解法1 一次遍历
pre表示前一个节点,cur表示当前节点
观察循环升序数组的特点,pre和cur的关系只有2种情况
1. pre <= cur
此时只需要pre <= x <= cur 就可以让x插入到pre和cur之间
2. pre > cur
此种情况时表示cur就是整个数组中最小的元素,此时只需要
 x >= pre变成新的最大元素
 x < cur变成新的最小元素
这2种情况下都可以插入到pre和cur之间
*/
func insert(aNode *Node, x int) *Node {
	addNode := &Node{Val: x}
	if aNode == nil {
		//空节点
		addNode.Next = addNode
		return addNode
	}
	if aNode.Next == nil {
		//只有一个节点
		aNode.Next = addNode
		addNode.Next = aNode
		return aNode
	}
	pre := aNode       //记录当前节点cur的前一个节点
	cur := aNode.Next  //当前节点从给定的aNode的下一个开始
	start := pre       //标记遍历的起始节点,防止无限循环遍历
	for cur != start { //从aNode的Next开始,遍历到aNode结束

		if pre.Val <= x && x <= cur.Val {
			/*
				1. cur非最小节点比如说[3,4,4,6,1]来说cur=2,3,4,5都算这种情况
			*/
			break
		} else if pre.Val > cur.Val {
			/*
				1. cur是最小节点比如说[3,4,4,6,1]来说cur=1就是这种情况
			*/
			if pre.Val <= x || x <= cur.Val {
				break
			}
		}
		pre = cur
		cur = cur.Next
	}
	/*
		1. 遍历完一圈还未找到合适的插入位置比如说[3,4,4,6,1]来说,当x=5,aNode=6时
	*/
	pre.Next = addNode
	addNode.Next = cur
	return aNode
}

// @lc code=end
