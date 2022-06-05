package main

import (
	. "leetcode/check/tree"
)

/*
 *
 *
 * 剑指 Offer 07. 重建二叉树
 */

// @lc code=start

/*
 解法1
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	//找到当前的根节点位置
	rootIdx := 0
	for rootIdx < len(inorder) && inorder[rootIdx] != preorder[0] {
		rootIdx++
	}
	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:1+rootIdx], inorder[:rootIdx]),
		Right: buildTree(preorder[1+rootIdx:], inorder[rootIdx+1:]),
	}
}

// @lc code=end
