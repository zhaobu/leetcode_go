package main

import (
	. "leetcode/check/tree"
)

/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
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
解法1:深度优先,递归
*/
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var (
		max func(node *TreeNode) int
	)
	max = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := max(node.Left) + 1
		right := max(node.Right) + 1
		if left > right {
			return left
		}
		return right
	}

	return max(root)
}

/*
解法2:广度优先,层序遍历求高度
*/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	q := []*TreeNode{root}
	height := 0
	levelNums := 1

	for len(q) > 0 {
		head := q[0]
		q = q[1:]
		levelNums--
		if head.Left != nil {
			q = append(q, head.Left)
		}
		if head.Right != nil {
			q = append(q, head.Right)
		}
		if levelNums == 0 {
			height++
			levelNums = len(q)
		}
	}

	return height
}

// @lc code=end
