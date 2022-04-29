package main

import (
	. "leetcode/check/tree"
)

/*
 * @lc app=leetcode.cn id=337 lang=golang
 *
 * [337] 打家劫舍 III
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
解法1 动态规划
*/
func rob1(root *TreeNode) int {
	dp0 := map[*TreeNode]int{} //表示不选中当前节点时, 以当前节点为根节点的树的最大金额
	dp1 := map[*TreeNode]int{} //表示选中当前节点时, 以当前节点为根节点的树的最大金额

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)

		/*
			1. 当node被选中时,左右孩子都不能被选中
			2. 当node不被选中时,左右孩子都有选或者不选2种情况,应该分别取左右孩子选和不选当中较大值,然后相加

		*/
		dp1[node] = node.Val + dp0[node.Left] + dp0[node.Right]
		dp0[node] = max(dp1[node.Left], dp0[node.Left]) + max(dp1[node.Right], dp0[node.Right])
	}

	dfs(root)

	return max(dp1[root], dp0[root])
}

/*
解法2 动态规划
解法1的优化
*/
func rob(root *TreeNode) int {

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	/*
		ret[0] 表示当前节点不选
		ret[1] 表示当前节点选
	*/
	var dfs func(node *TreeNode) (ret [2]int)
	dfs = func(node *TreeNode) (ret [2]int) {
		if node == nil {
			return
		}
		left := dfs(node.Left)
		right := dfs(node.Right)

		//当node不被选中时,左右孩子都有选或者不选2种情况,应该分别取左右孩子选和不选当中较大值,然后相加
		ret[0] = max(left[0], left[1]) + max(right[0], right[1])
		//当node被选中时,左右孩子都不能被选中
		ret[1] = node.Val + left[0] + right[0]
		return
	}

	ret := dfs(root)

	return max(ret[0], ret[1])
}

// @lc code=end
