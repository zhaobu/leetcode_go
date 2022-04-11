package main

import (
	"fmt"
	. "leetcode/check/tree"
	"strings"
)

/*
 * @lc app=leetcode.cn id=572 lang=golang
 *
 * [572] 另一棵树的子树
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
 思路:
一个树是另一个树的子树 则
1.要么这两个树相等
2.要么这个树是左树的子树
3.要么这个树是右树的子树
*/
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return subRoot == nil
	}

	var isSameTree func(r1, r2 *TreeNode) bool
	isSameTree = func(r1, r2 *TreeNode) bool {
		if r1 == nil && r2 == nil {
			return true
		}

		return r1 != nil && r2 != nil && r1.Val == r2.Val && isSameTree(r1.Left, r2.Left) && isSameTree(r1.Right, r2.Right)
	}

	return isSameTree(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

/*
思路:
1.二叉树用前序遍历方式序列化成字符串
2.判断子树序列化后的字符串是否是root数序列化后的字符串的子串
3.注意如果用前序遍历序列化,可能存在root的根结点的值是12,
subRoot的根节点的值为2时, 判断错误的情况,所以序列化时在前面加上一个特殊字符,比如*
就会变成 *12 和 *2 ,判断子串时就不相等
*/
func isSubtree2(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return true
	}

	var strTree func(str *string, root *TreeNode)

	strTree = func(str *string, root *TreeNode) {
		if root == nil {
			*str += "#;"
			return
		}
		*str += fmt.Sprintf("*%d;", root.Val)
		strTree(str, root.Left)
		strTree(str, root.Right)
		return
	}

	var rootStr, subRootStr string
	strTree(&rootStr, root)
	strTree(&subRootStr, subRoot)

	fmt.Printf("rootStr=%s\nsubRootStr=%s", rootStr, subRootStr)
	return strings.Contains(rootStr, subRootStr)
}

/*
思路:
1.二叉树用后序遍历方式序列化成字符串
2.判断子树序列化后的字符串是否是root数序列化后的字符串的子串
*/
func isSubtree1(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return true
	}

	var strTree func(str *string, root *TreeNode)

	strTree = func(str *string, root *TreeNode) {
		if root == nil {
			*str += "#;"
			return
		}
		strTree(str, root.Left)
		strTree(str, root.Right)
		*str += fmt.Sprintf("*%d;", root.Val)
		return
	}

	var rootStr, subRootStr string
	strTree(&rootStr, root)
	strTree(&subRootStr, subRoot)
	// fmt.Printf("rootStr=%s\nsubRootStr=%s", rootStr, subRootStr)
	return strings.Contains(rootStr, subRootStr)
}

// @lc code=end
