package main

import (
	. "leetcode/check/tree"
)

/*
 * @lc app=leetcode.cn id=687 lang=golang
 *
 * [687] 最长同值路径
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
解法1 后序遍历
*/
func longestUnivaluePath1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ret := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	/*
		返回以root为根节点的树的必须经过root节点的最长的同值路径的长度
	*/
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)

		leftMax, rightMax := 1, 1                          //分别表示根节点+左右子树最长路径长度
		if root.Left != nil && root.Left.Val == root.Val { //必须根节点和左孩子节点相等才可能走和左子树构成同值路径
			leftMax += left
		}
		if root.Right != nil && root.Right.Val == root.Val { //必须根节点和右孩子节点相等才可能走和左子树构成同值路径
			rightMax += right
		}
		curMax := max(leftMax, rightMax)
		/*
		 必须经过root节点的同值路径有4种情况:
		 1. 只经过根节点
		 2. 根节点和左孩子值相同
		 3. 根节点和右孩子值相同
		 4. 根节点和左右孩子都相同
		*/
		curMaxUnivalue := curMax
		if leftMax > 1 && rightMax > 1 && root.Val == root.Left.Val { //考虑第4种情况
			curMaxUnivalue = leftMax + right
		}
		ret = max(ret, curMaxUnivalue-1) //所有的值都是求的路径上点的个数,实际上要求的是边的个数,所以应该减1
		return curMax
	}

	dfs(root)
	return ret
}

/*
解法2 后序遍历
方法1的更精简写法
*/
func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ret := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)

		leftMax, rightMax := 0, 0 //相比于方法1,初始化为0可以更方便的判断第4种情况
		if root.Left != nil && root.Left.Val == root.Val {
			leftMax = left + 1
		}
		if root.Right != nil && root.Right.Val == root.Val {
			rightMax = right + 1
		}
		ret = max(ret, leftMax+rightMax)
		return max(leftMax, rightMax)
	}

	dfs(root)
	return ret
}

// @lc code=end
