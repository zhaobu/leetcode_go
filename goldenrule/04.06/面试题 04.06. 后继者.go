package main

import (
	"fmt"
	. "leetcode/check/tree"
)

/*
 * [面试题 面试题 04.06. 后继者](https://leetcode.cn/problems/successor-lcci/)
 */

/*
 解法1: 动态规划,完全不优化版,timeout
 背包九讲里面的完全背包类似.
 dp定义: dp[i][j]表示前i种硬币凑齐j分的方法总数
 basecase:
如果待凑的是0分,则只有一个解就是所有硬币个数都为0  dp[i][0]=1
如果硬币种类数为0, 则凑齐任意分的方法总数都为0 dp[0][i]=0
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
解法1 中序遍历
*/
func inorderSuccessor1(root *TreeNode, p *TreeNode) *TreeNode {
	var (
		ret *TreeNode
		pre *TreeNode
	)

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if ret != nil || root == nil {
			return
		}
		dfs(root.Left)
		if pre == p {
			ret = root
		}
		pre = root
		dfs(root.Right)
	}

	dfs(root)

	return ret
}

/*
解法2 二叉搜素数性质
*/
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	ret := p.Right
	//如果p的右子树不为空,后继就在右子树上
	for ret != nil && ret.Left != nil {
		ret = ret.Left
	}
	if ret != nil {
		fmt.Printf("1")
		return ret
	}
	/*
		1. 如果p的右子树为空,就要找到p的第一个不小于p的祖先节点
	*/
	cur := root
	for cur != nil {
		if cur.Val <= p.Val {
			cur = cur.Right
		} else {
			/*
				1. 如果cur > p,则后继可能就是cur或者在cur的左子树上
				2. 如果cur就是后继,说明cur的左子树不存在>p的节点.
				3. 如果cur的后继在左子树上,也需要往左遍历
			*/
			ret = cur
			cur = cur.Left
		}
	}
	return ret
}
