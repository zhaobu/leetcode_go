package main

import (
	. "leetcode/check/tree"
)

/*
 * @lc app=leetcode.cn id=515 lang=golang
 *
 * [515] 在每个树行中找最大值
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
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
 解法1 层序遍历
*/
func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	ret := []int{}

	queue := []*TreeNode{root}
	curCnt, curMax := 1, root.Val
	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]

		if head.Left != nil {
			queue = append(queue, head.Left)
		}
		if head.Right != nil {
			queue = append(queue, head.Right)
		}
		curCnt--
		if head.Val > curMax {
			curMax = head.Val
		}

		if curCnt == 0 {
			ret = append(ret, curMax)
			curCnt = len(queue)
			if curCnt > 0 {
				curMax = queue[0].Val
			}
		}
	}

	return ret
}

// @lc code=end
