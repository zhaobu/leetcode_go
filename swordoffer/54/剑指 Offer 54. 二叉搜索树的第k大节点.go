package main

import (
	. "leetcode/check/tree"
)

/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * 剑指 Offer 54. 二叉搜索树的第k大节点
 */

// @lc code=start

/*
 解法1
 动态规划
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
解法1 迭代中序遍历
注意顺序是 右,中,左
*/
func kthLargest(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Right
		}
		root = stack[len(stack)-1]
		k--
		if k == 0 {
			break
		}
		stack = stack[:len(stack)-1]
		root = root.Left
	}

	return root.Val
}

// @lc code=end
