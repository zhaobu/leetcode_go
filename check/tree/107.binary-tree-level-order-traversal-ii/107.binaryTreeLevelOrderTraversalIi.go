package main

import (
	. "leetcode/check/tree"
)

/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层序遍历 II
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
func levelOrderBottom(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	levelNums := 1
	curLevel := []int{}
	for len(q) > 0 {
		head := q[0]
		q = q[1:]
		levelNums--
		curLevel = append(curLevel, head.Val)
		if head.Left != nil {
			q = append(q, head.Left)
		}
		if head.Right != nil {
			q = append(q, head.Right)
		}
		if levelNums == 0 {
			levelNums = len(q)
			curLevel = make([]int, 0, levelNums)
		}
	}
	//翻转
	for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
		ret[i], ret[j] = ret[j], ret[i]
	}
	return ret
}

// @lc code=end
