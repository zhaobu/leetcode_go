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
		curLevel  = []*TreeNode{root} //当前层的所有节点
		nextLevel []*TreeNode         //下一层的所有节点
		maxWidth  = 1                 //最大宽度
		max       func(a, b int) int
	)
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	root.Val = 0 //用节点的值来保存index
	for len(curLevel) > 0 {
		//遍历当前层的所有节点
		nextLevel = make([]*TreeNode, 0, 2*len(curLevel))
		//访问当前层所有节点
		maxWidth = max(maxWidth, curLevel[len(curLevel)-1].Val-curLevel[0].Val+1)
		for i := 0; i < len(curLevel); i++ {
			if curLevel[i].Left != nil {
				curLevel[i].Left.Val = 2*curLevel[i].Val + 1
				nextLevel = append(nextLevel, curLevel[i].Left)
			}
			if curLevel[i].Right != nil {
				curLevel[i].Right.Val = 2*curLevel[i].Val + 2
				nextLevel = append(nextLevel, curLevel[i].Right)
			}
		}
		curLevel = nextLevel
	}
	return maxWidth
}
