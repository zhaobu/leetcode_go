package main

import (
	. "leetcode/check/tree"
)

/*
 * @lc app=leetcode.cn id=1022 lang=golang
 *
 * [1022] 从根到叶的二进制数之和
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
 后序遍历,注意可以用或的方式直接计算二进制数
*/
func sumRootToLeaf1(root *TreeNode) int {
	ret := 0

	var dfs func(node *TreeNode, val int)
	dfs = func(node *TreeNode, val int) {
		if node == nil {
			return
		}
		val = (val << 1) | node.Val
		dfs(node.Left, val)
		dfs(node.Right, val)
		if node.Left == nil && node.Right == nil {
			ret += val
		}
	}

	dfs(root, 0)
	return ret
}

/*
解法2 迭代
	     0
     1         2
   3    4    5    6
*/
func sumRootToLeaf(root *TreeNode) int {
	ret := 0

	stack := []*TreeNode{}
	cur := root
	pre := (*TreeNode)(nil)
	val := 0
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			val = (val << 1) | cur.Val
			stack = append(stack, cur)
			cur = cur.Left
		}
		top := stack[len(stack)-1]
		if top.Right == nil || top.Right == pre {
			if top.Right == nil && top.Left == nil { //不能写成top.Right == nil
				ret += val
			}
			// fmt.Printf("top.val=%d\n", top.Val)
			// fmt.Printf("前val=%b\n", val)
			// val = (val & (^top.Val)) >> 1
			val >>= 1
			// fmt.Printf("后val=%b\n\n", val)
			pre = top
			stack = stack[:len(stack)-1]
			cur = nil
		} else {
			cur = top.Right
		}
	}

	return ret
}

// @lc code=end
