package main

import (
	. "leetcode/check/tree"
)

/*
 * @lc app=leetcode.cn id=112 lang=golang
 *
 * [112] 路径总和
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
 解法1 递归dfs
*/
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	ret := false
	var dfs func(root *TreeNode, sum int)
	dfs = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}
		sum += root.Val
		//左右孩子都为nil的才叫算叶子节点
		if root.Left == nil && root.Right == nil && sum == targetSum {
			ret = true
		}
		dfs(root.Left, sum)
		dfs(root.Right, sum)
	}

	dfs(root, 0)

	return ret
}

// @lc code=end
