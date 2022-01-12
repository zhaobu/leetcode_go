package main

/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
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
func preorderTraversal1(root *TreeNode) (res []int) {
	if root == nil {
		return nil
	}
	var preorder func(root *TreeNode)
	preorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		if root.Left != nil {
			preorder(root.Left)
		}
		if root.Right != nil {
			preorder(root.Right)
		}
	}
	preorder(root)
	return res
}

//解法2:迭代
func preorderTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		//模拟左节点入栈
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return res
}

// @lc code=end
