package main

import (
	. "leetcode/check/tree"
)

/*
 * @lc app=leetcode.cn id=513 lang=golang
 *
 * [513] 找树左下角的值
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
func findBottomLeftValue1(root *TreeNode) int {
	maxDepth := 0
	ret := 0
	var dfs func(root *TreeNode, curDepth int)
	dfs = func(root *TreeNode, curDepth int) {
		if root == nil {
			return
		}
		if curDepth > maxDepth {
			ret = root.Val
			maxDepth = curDepth
		}
		dfs(root.Left, curDepth+1)
		dfs(root.Right, curDepth+1)
	}
	dfs(root, 1)

	return ret
}

/*
解法2 bfs
*/
func findBottomLeftValue(root *TreeNode) int {
	queue := []*TreeNode{root}
	ret := root.Val
	levelCnt := 1
	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]
		if head.Left != nil {
			queue = append(queue, head.Left)
		}
		if head.Right != nil {
			queue = append(queue, head.Right)
		}
		levelCnt--
		if levelCnt == 0 {
			//层序遍历当遍历完每层最后一个节点,确定下一层个数时记录下一层的第一个节点
			levelCnt = len(queue)
			if levelCnt > 0 {
				ret = queue[0].Val
			}
		}
	}
	return ret
}

// @lc code=end
