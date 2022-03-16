package main

import (
	. "leetcode/check/tree"
)

/*
 * @lc app=leetcode.cn id=108 lang=golang
 *
 * [108] 将有序数组转换为二叉搜索树
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
1. 找到树的根节点,找到树的左孩子,找到树的右孩子,然后把左右孩子分别和根节点连接
2. 因为是BST,数组有序,所以根节点每次都是数组的中间数
*/

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) < 1 {
		return nil
	}

	var genBST func(i, j int) *TreeNode
	genBST = func(i, j int) *TreeNode {
		if i > j {
			return nil
		}
		mid := i + (j-i)>>1
		// fmt.Printf("mid=%+v\n", mid)
		node := &TreeNode{Val: nums[mid]}
		node.Left = genBST(i, mid-1)
		node.Right = genBST(mid+1, j)
		return node
	}

	return genBST(0, len(nums)-1)
}

// @lc code=end
