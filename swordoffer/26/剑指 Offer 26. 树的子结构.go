package main

import (
	. "leetcode/check/tree"
)

/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * 剑指 Offer 26. 树的子结构
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
*/
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return false
	}

	//判断root2是否是root1的子结构
	var check func(root1, root2 *TreeNode) bool
	check = func(root1, root2 *TreeNode) bool {
		if root1 == nil && root2 != nil {
			return false
		}
		if root2 == nil {
			return true
		}
		if root1.Val != root2.Val {
			return false
		}
		left := check(root1.Left, root2.Left)
		right := check(root1.Right, root2.Right)
		return left && right
	}

	/*
	   1. 初始化ret为0, 表示在A中没找到和B的根节点相同值的节点
	   2. 当 ret=1 时, 表示已经找到了一个B是A的子结构的情况
	   3. 当 ret=-1时, 表示根节点和B相同,但是不是子结构
	   4. A树中可能存在节点值重复的情况
	*/
	ret := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if ret == 1 || node == nil {
			return
		}
		/*
		   1. 不能后序遍历
		   2. 类似
		       [4,2,3,4,5,6,7,8,9]
		       [4,8,9]
		       这样的情况,A树上节点有重复值
		*/
		if node.Val == B.Val {
			// fmt.Printf("root1=[%d,%d,%d], root2=%d \n", node.Val, node.Left.Val, node.Right.Val, B.Val)
			if check(node, B) {
				ret = 1
			} else {
				ret = -1
			}
		}
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(A)
	// fmt.Printf("ret=%d\n", ret)
	return ret == 1
}

// @lc code=end
