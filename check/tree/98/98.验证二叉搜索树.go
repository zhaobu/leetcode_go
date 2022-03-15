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
 解法1 中序遍历
*/
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var preNode *TreeNode
	isBST := true
	var bst func(root *TreeNode)
	bst = func(root *TreeNode) {
		if root == nil {
			return
		}
		bst(root.Left)
		// fmt.Printf("root=%+v\n", root.Val)
		if preNode != nil && root.Val <= preNode.Val {
			isBST = false
		}
		preNode = root
		bst(root.Right)
		return
	}
	bst(root)
	return isBST
}

// @lc code=end
