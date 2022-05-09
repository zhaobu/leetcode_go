package main

import (
	. "leetcode/check/tree"
)

/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
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
*/
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	ret := true
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	/*
	   递归函数返回node为根节点的树的高度
	*/
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		if !ret { //如果已经判断出结果了,就尽快结束所有递归
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		if left-right > 1 || left-right < -1 {
			ret = false
		}
		return max(left, right) + 1
	}

	dfs(root)

	return ret
}

// @lc code=end
