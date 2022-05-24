package main

import (
	. "leetcode/check/tree"
)

/*
 * @lc app=leetcode.cn id=965 lang=golang
 *
 * [965] 单值二叉树
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
func isUnivalTree1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && root.Val != root.Left.Val {
		return false
	}
	if root.Right != nil && root.Val != root.Right.Val {
		return false
	}
	return isUnivalTree(root.Left) && isUnivalTree(root.Right)
}

/*
解法2 迭代遍历(后序遍历实现)
        0
    1       2
  3   4   5   6
*/
func isUnivalTree(root *TreeNode) bool {

	var (
		stack = []*TreeNode{}
		cur   = root
		pre   *TreeNode //上一个访问的节点
	)
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		top := stack[len(stack)-1]
		if top.Right == nil || top.Right == pre {
			pre = top
			cur = nil
			stack = stack[:len(stack)-1]
			if top.Val != root.Val {
				return false
			}
		} else {
			cur = top.Right
		}
	}
	return true
}

// @lc code=end
