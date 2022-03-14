package main

import (
	. "leetcode/check/tree"
)

/*
 * @lc app=leetcode.cn id=236 lang=golang
 *
 * [236] 二叉树的最近公共祖先
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
 1. 如果p,q同时存在于以root为根节点的二叉树中,最近公共祖先就是root
 2. 如果p,q都不存在于以root为根节点的二叉树中,返回nil
 3. 如果只有p存在于以root为根节点的二叉树中,返回p
 4. 如果只有q存在于以root为根节点的二叉树中,返回q
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}

	//在左子树中找公共父节点
	left := lowestCommonAncestor(root.Left, p, q)
	//在右子树中找公共父节点
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left == nil && right == nil {
		return nil
	}

	if left != nil {
		return left
	}
	return right
}

// @lc code=end
