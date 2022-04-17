package main

import (
	. "leetcode/check/tree"
)

/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
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
 解法1
*/
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ret := root.Val //存在节点值为负数的情况,不能初始化为0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	/*
		dfs 返回值为以root为根节点的树的最大贡献值
		1. 空节点的最大贡献值等于 0
		2. 非空节点的最大贡献值等于节点值与其子节点中的最大贡献值之和
		3. 对于叶节点而言，最大贡献值等于节点值
	*/
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		/*
			1. 计算最大贡献值时,根节点必选
			2. 如果子树存在贡献值为正数的情况,就加上左右子树中贡献值较大的那个
		*/
		cur := max(root.Val, root.Val+max(left, right))
		/*
			对于节点root来说,计算最大路径和共有四种情况:
			1. 只包含根节点
			2. 左子树最大路径+根节点
			3. 右子树最大路径+根节点
			4. 左子树最大路径+根节点+右子树最大路径
		*/
		curMax := max(cur, root.Val+left+right) //前3种情况就是求cur,还要和第4种情况比较
		ret = max(curMax, ret)
		return cur
	}

	dfs(root)
	return ret
}

// @lc code=end
