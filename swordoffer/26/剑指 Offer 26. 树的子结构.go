package main

import (
	. "leetcode/check/tree"
)

/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * 剑指 Offer 26. 树的子结构
 */

// @lc code=start

/*
 解法1 dfs
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return false
	}

	//判断root2是否是root1的子结构
	var check func(root1, root2 *TreeNode) bool
	check = func(root1, root2 *TreeNode) bool {
		if root1 == nil && root2 != nil {
			return false
		}
		if root2 == nil {
			return true
		}
		if root1.Val != root2.Val {
			return false
		}
		left := check(root1.Left, root2.Left)
		right := check(root1.Right, root2.Right)
		return left && right
	}

	ret := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if ret != 0 || node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		if node.Val == B.Val {
			if check(node, B) {
				ret = 1
			} else {
				ret = -1
			}
		}
	}

	dfs(A)

	return ret == 1
}

// @lc code=end
