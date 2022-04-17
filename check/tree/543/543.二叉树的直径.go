package main

import (
	. "leetcode/check/tree"
)

/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
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
 解法1 后序遍历
 注意最大直径不一定经过根节点
*/
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ret := 0
	var maxDepth func(root *TreeNode) int
	maxDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := maxDepth(root.Left)
		right := maxDepth(root.Right)
		cur := left + right
		// fmt.Printf("left=%d,right=%d\n",left,right)
		if cur > ret {
			ret = cur
		}
		if left > right {
			return left + 1
		}
		return right + 1
	}

	maxDepth(root)
	return ret
}

// @lc code=end
