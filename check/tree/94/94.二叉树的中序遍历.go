package main

/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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

//方法1:递归
func inorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var (
		res     []int
		inorder func(root *TreeNode)
	)

	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			inorder(root.Left)
		}
		res = append(res, root.Val)
		if root.Right != nil {
			inorder(root.Right)
		}
	}
	inorder(root)
	return res
}

//方法2:
func inorderTraversal(root *TreeNode) []int {
}

// @lc code=end
