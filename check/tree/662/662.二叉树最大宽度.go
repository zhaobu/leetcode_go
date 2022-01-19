package main

/*
 * @lc app=leetcode.cn id=662 lang=golang
 *
 * [662] 二叉树最大宽度
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
func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var (
		curLevel  = []*TreeeNode{root} //当前层的所有节点
		levelNums = 1
		nextLevel []*TreeNode //下一层的所有节点
	)
	root.Val = 0
	for len(curLevel) > 0 {
		//遍历当前层的所有节点
		nextLevel = make([]*TreeeNode, 0, len(curLevel))
		for i := 0; i < len(curLevel); i++ {
			if curLevel[i] 
		}
	}
}
