package main

/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
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
//解法1:递归
func postorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var (
		res       []int
		postorder func(node *TreeNode)
	)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			postorder(node.Left)
		}
		if node.Right != nil {
			postorder(node.Right)
		}
		res = append(res, node.Val)
	}
	postorder(root)
	return res
}

/*
解法2:
*/
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var (
		res     []int
		node    = root
		stack   []*TreeNode
		preNode *TreeNode //上次访问的节点
	)

	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Right == nil || node.Right == preNode {
			res = append(res, node.Val)
			preNode = node
			node = nil
		} else {
			stack = append(stack, node)
			node = node.Right
		}
	}
	return res
}

// @lc code=end
