package main

/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
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
func buildTree1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	index := 0 //rootVal在中序遍历数组中的索引

	for i, v := range inorder {
		if v == root.Val {
			index = i
			break
		}
	}
	/*
	 	0 *1 2 3 4 5 * 6 7 8 9 root * left * right
	   	1 2 3 4 5 * 0 * 6 7 8 9  left * root * right
	*/
	root.Left = buildTree(preorder[1:index+1], inorder[0:index])
	root.Right = buildTree(preorder[index+1:len(preorder)], inorder[index+1:len(inorder)])
	return root
}

/*
 解法2: 迭代
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	return nil
}

// @lc code=end
