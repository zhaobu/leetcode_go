package main

import (
	. "leetcode/check/tree"
)

/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第K小的元素
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
 注意如果不加是否已经找到的find标记,不能提前退出递归,
 会造成2次到达k
*/
func kthSmallest1(root *TreeNode, k int) int {
	ret := 0
	index := 1
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		// fmt.Printf("root=%+v, index=%d \n", root.Val, index)
		if index == k {
			// fmt.Printf("find root=%+v, index=%d \n", root.Val, index)
			ret = root.Val
			//此处不能提前退出,会造成2次index=k的情况
		}
		index++
		inorder(root.Right)
	}

	inorder(root)
	return ret
}

/*
解法2 递归优化版
*/
func kthSmallest2(root *TreeNode, k int) int {
	find := false //增加find变量,可以在找到后就不用继续后面的递归
	ret := 0
	index := 1
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		// fmt.Printf("root=%+v, index=%d \n", root.Val, index)
		if !find && index == k {
			// fmt.Printf("find root=%+v, index=%d \n", root.Val, index)
			ret = root.Val
			find = true
			return
		}
		index++
		inorder(root.Right)
	}

	inorder(root)
	return ret
}

/*
解法3 更简洁写法
总结就是k--或者k++要放在前面写,如果先判断k然后再--,就会造成2次一样的k
*/
func kthSmallest(root *TreeNode, k int) int {

	ret := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		k--
		if k == 0 {
			ret = node.Val
			return
		}
		dfs(node.Right)
	}

	dfs(root)

	return ret
}

// @lc code=end
