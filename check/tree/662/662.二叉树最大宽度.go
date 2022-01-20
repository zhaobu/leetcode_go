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

/*
 解法1:层序遍历
*/
func widthOfBinaryTree1(root *TreeNode) int {
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

/*
 解法2:深度优先
*/
func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var (
		maxWidth  = 0
		dfs       func(node *TreeNode, deepth, index int)
		max       func(a, b int) int
		leftIndex = map[int]int{} //用于保存每层的最左边节点的index
	)
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dfs = func(node *TreeNode, deepth, index int) {
		if node == nil {
			return
		}
		if _, ok := leftIndex[deepth]; !ok {
			leftIndex[deepth] = index
		}
		maxWidth = max(maxWidth, index-leftIndex[deepth]+1)
		dfs(node.Left, deepth+1, 2*index+1)
		dfs(node.Right, deepth+1, 2*index+2)
	}
	dfs(root, 0, 0)
	return maxWidth
}
