package main

import (
	. "leetcode/check/tree"
)

/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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
func levelOrder1(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var (
		queue     = []*TreeNode{root}
		ret       [][]int
		levelNums = 1   //每一层的节点个数
		curValues []int //当前层的节点值
	)
	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]
		levelNums--                             //每遍历一个节点,当前层节点个数-1
		curValues = append(curValues, head.Val) //访问节点
		if head.Left != nil {
			queue = append(queue, head.Left)
		}
		if head.Right != nil {
			queue = append(queue, head.Right)
		}
		if levelNums == 0 {
			levelNums = len(queue)
			ret = append(ret, curValues)
			curValues = make([]int, 0, levelNums)
		}
	}
	return ret
}

func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	curLevelQueue := []*TreeNode{root} //当前层的所有节点
	for i := 0; len(curLevelQueue) > 0; i++ {
		ret = append(ret, make([]int, 0, len(curLevelQueue)))
		nextLevelQueue := []*TreeNode{} //下一层的所有节点
		//访问当前层的每一个节点,并把下一层的节点加入到队列
		for j := 0; j < len(curLevelQueue); j++ {
			node := curLevelQueue[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				nextLevelQueue = append(nextLevelQueue, node.Left)
			}
			if node.Right != nil {
				nextLevelQueue = append(nextLevelQueue, node.Right)
			}
		}
		curLevelQueue = nextLevelQueue //接着处理下一层
	}
	return ret
}

// @lc code=end
