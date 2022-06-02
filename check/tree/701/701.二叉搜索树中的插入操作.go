package main

import (
	. "leetcode/check/tree"
)

/*
 * @lc app=leetcode.cn id=701 lang=golang
 *
 * [701] 二叉搜索树中的插入操作
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
func insertIntoBST1(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if val > node.Val {
			if node.Right == nil {
				node.Right = &TreeNode{Val: val}
			} else {
				dfs(node.Right)
			}
		} else {
			if node.Left == nil {
				node.Left = &TreeNode{Val: val}
			} else {
				dfs(node.Left)
			}
		}
	}

	dfs(root)
	return root
}

/*
解法2 迭代
*/
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	cur := root
	for cur != nil {
		if val > cur.Val {
			if cur.Right == nil {
				cur.Right = &TreeNode{Val: val}
				break
			}
			cur = cur.Right
		} else {
			if cur.Left == nil {
				cur.Left = &TreeNode{Val: val}
				break
			}
			cur = cur.Left
		}
	}

	return root
}

// @lc code=end
