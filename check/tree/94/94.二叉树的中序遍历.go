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

//方法2:迭代
func inorderTraversal2(root *TreeNode) (res []int) {
	if root == nil {
		return nil
	}

	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		//模拟所有左节点入栈
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		//栈顶出栈
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 打印出栈元素
		res = append(res, root.Val)
		//遍历当前节点的右子节点
		root = root.Right
	}
	return
}

//方法3:Morris 中序遍历
func inorderTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return nil
	}
	return
}

// @lc code=end
