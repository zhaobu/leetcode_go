package main

import (
	. "leetcode/check/tree"
)

/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
解法1 层序遍历
等价于层序遍历时,每层的最后一个元素
*/
func rightSideView1(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	count := 1
	ret := []int{}

	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]
		count--
		if head.Left != nil {
			queue = append(queue, head.Left)
		}

		if head.Right != nil {
			queue = append(queue, head.Right)
		}

		if count == 0 {
			ret = append(ret, head.Val)
			count = len(queue)
		}
	}

	return ret
}

/*
解法2 dfs
1. 按照 根,右,左的顺序dfs遍历数
2. 记录最大深度,每次到达一个新的最大深度时的第一个节点就是最右边的值
*/
func rightSideView2(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	ret := []int{}
	maxDepth := -1

	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if depth > maxDepth {
			ret = append(ret, root.Val)
			maxDepth = depth
		}
		dfs(root.Right, depth+1)
		dfs(root.Left, depth+1)
	}

	dfs(root, 0)
	return ret
}

/*
解法3 栈
1. 记录每个节点的深度,按照根左右顺序入栈,这样出栈时,右孩子总是第一个出栈
2. 记录最大深度,每次第一次出现节点的深度>当前记录的最大深度时就记录该节点的值
*/
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	ret := []int{}
	maxDepth := -1
	type Node struct {
		node  *TreeNode
		depth int
	}
	stack := []*Node{{
		node:  root,
		depth: 0,
	}}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		if top.depth > maxDepth {
			maxDepth = top.depth
			ret = append(ret, top.node.Val)
		}
		stack = stack[:len(stack)-1]

		if top.node.Left != nil {
			stack = append(stack, &Node{
				node:  top.node.Left,
				depth: top.depth + 1,
			})
		}
		if top.node.Right != nil {
			stack = append(stack, &Node{
				node:  top.node.Right,
				depth: top.depth + 1,
			})
		}
	}
	return ret
}

// @lc code=end
