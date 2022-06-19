package main

import (
	. "leetcode/check/tree"
)

/*
 * @lc app=leetcode.cn id=508 lang=golang
 *
 * [508] 出现次数最多的子树元素和
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

func findFrequentTreeSum(root *TreeNode) []int {
	record := map[int]int{}
	maxCnt := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		cur := left + right + node.Val
		record[cur]++
		if record[cur] > maxCnt {
			maxCnt = record[cur]
		}
		return cur
	}

	dfs(root)

	ret := []int{}
	for num, cnt := range record {
		if cnt == maxCnt {
			ret = append(ret, num)
		}
	}

	return ret
}

// @lc code=end
