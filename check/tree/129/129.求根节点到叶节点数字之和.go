package main

import (
	. "leetcode/check/tree"
)

/*
 * @lc app=leetcode.cn id=129 lang=golang
 *
 * [129] 求根节点到叶节点数字之和
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
 解法1 dfs
1. 只要没有子节点就算找到了一条路径
*/
func sumNumbers1(root *TreeNode) int {
	if root == nil {
		/*
			如果放在这里来计算sum,会存在以下几个问题
			1. 如果下面没有if root.Right != nil和if root.Left != nil的判断,sum会计算2次.至少要判断一次
			2. 如果是只有2层2个节点的二叉树,会把根节点也计算一遍,如果要判断,要加一个depth参数
		*/
		return 0
	}
	ret := 0
	var dfs func(root *TreeNode, num int)
	dfs = func(root *TreeNode, num int) {
		if root == nil {
			return
		}
		num = num*10 + root.Val
		if root.Left == nil && root.Right == nil { //只要没有子节点说明已经到了叶子节点
			ret += num
			return
		} else {
			if root.Left != nil {
				dfs(root.Left, num)
			}
			if root.Right != nil {
				dfs(root.Right, num)
			}
		}
	}

	dfs(root, 0)
	return ret
}

/*
 解法2 层序遍历
*/
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ret := 0
	queue1 := []*TreeNode{root}
	queue2 := []int{root.Val}

	for len(queue1) > 0 {
		head1 := queue1[0]
		queue1 = queue1[1:]
		head2 := queue2[0]
		queue2 = queue2[1:]
		if head1.Left == nil && head1.Right == nil {
			ret += head2
			continue
		}
		if head1.Left != nil {
			queue1 = append(queue1, head1.Left)
			queue2 = append(queue2, head2*10+head1.Left.Val)
		}
		if head1.Right != nil {
			queue1 = append(queue1, head1.Right)
			queue2 = append(queue2, head2*10+head1.Right.Val)
		}
	}
	return ret
}

// @lc code=end
