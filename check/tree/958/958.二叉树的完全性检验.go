package main

import (
	. "leetcode/check/tree"
)

/*
 * @lc app=leetcode.cn id=958 lang=golang
 *
 * [958] 二叉树的完全性检验
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
func isCompleteTree(root *TreeNode) bool {
	queue := []*TreeNode{root}

	flag := false //是否已经全部节点入队过
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		if flag {
			//全部节点入队后,后面的节点必定为叶子节点
			if top.Left != nil || top.Right != nil {
				return false
			}
		} else {
			//完全二叉树第一个空节点必须是右孩子,且此时左孩子不能为空
			if top.Left == nil && top.Right != nil {
				return false
			}
			//第一次遇到空节点,说明全部节点已经入队过
			if top.Right == nil {
				flag = true
			}
		}
		if top.Left != nil {
			queue = append(queue, top.Left)
		}
		if top.Right != nil {
			queue = append(queue, top.Right)
		}
	}

	return true
}

// @lc code=end
