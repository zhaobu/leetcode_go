package main

import (
	. "leetcode/check/tree"
)

/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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
 解法1 递归
*/
func isValidBST1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var bst func(root, min, max *TreeNode) bool
	bst = func(root, min, max *TreeNode) bool {
		if root == nil {
			return true
		}
		if min != nil && root.Val <= min.Val {
			return false
		}
		if max != nil && root.Val >= max.Val {
			return false
		}
		/*
			1. 左孩子全部节点都小于root,右孩子全部节点都大于root
			2. min和max必须通过参数往下传递
		*/
		return bst(root.Left, min, root) && bst(root.Right, root, max)
	}

	return bst(root, nil, nil)
}

/*
 解法2 中序遍历递归
*/
func isValidBST2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var preNode *TreeNode
	var bst func(root *TreeNode) bool
	bst = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		if !bst(root.Left) {
			return false
		}
		// fmt.Printf("root=%+v\n", root.Val)
		if preNode != nil && root.Val <= preNode.Val {
			return false
		}
		preNode = root
		return bst(root.Right)
	}

	return bst(root)
}

/*
 解法3 中序遍历迭代
*/
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var preNode *TreeNode
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if preNode != nil && root.Val <= preNode.Val {
			return false
		}
		preNode = root
		root = root.Right
	}

	return true
}

// @lc code=end
