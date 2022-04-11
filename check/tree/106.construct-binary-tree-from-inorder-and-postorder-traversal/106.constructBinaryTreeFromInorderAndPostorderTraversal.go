package main

import (
	. "leetcode/check/tree"
)

/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] Construct Binary Tree from Inorder and Postorder Traversal
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
 解法1: 递归
*/
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	/*
	   0 1 2 3 * 4 * 5 6 7 8 9 left*root*right inorder
	   0 1 2 3 * 5 6 7 8 9 * 4 left*right*root postorder
	*/
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	//根据中序遍历找到root的index
	index := 0
	for i, v := range inorder {
		if v == root.Val {
			index = i
			break
		}
	}
	root.Left = buildTree(inorder[:index], postorder[:index])
	root.Right = buildTree(inorder[index+1:], postorder[index:len(postorder)-1])
	return root
}

// @lc code=end
