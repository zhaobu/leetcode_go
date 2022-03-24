package main

import (
	. "leetcode/check/tree"
)

/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
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
解法1 递归 dfs
*/
func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}

	ret := [][]int{}
	var dfs func(root *TreeNode, sum int, record []int)
	dfs = func(root *TreeNode, sum int, record []int) {
		if root == nil {
			return
		}
		sum += root.Val
		record = append(record, root.Val)
		//左右孩子都为nil的才叫算叶子节点
		if root.Left == nil && root.Right == nil && sum == targetSum {
			ret = append(ret, append([]int{}, record...))
		}
		dfs(root.Left, sum, record)
		dfs(root.Right, sum, record)
	}

	dfs(root, 0, []int{})

	return ret
}

// @lc code=end
