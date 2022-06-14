package main

import (
	. "leetcode/check/tree"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
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
 解法1 dfs
*/
func binaryTreePaths(root *TreeNode) []string {
	ret := []string{}
	var dfs func(root *TreeNode, record string)
	dfs = func(root *TreeNode, record string) {
		if root == nil {
			return
		}
		record += strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			ret = append(ret, record)
			return
		}
		record += "->"
		dfs(root.Left, record)
		dfs(root.Right, record)
	}

	dfs(root, "")

	return ret
}

// @lc code=end
