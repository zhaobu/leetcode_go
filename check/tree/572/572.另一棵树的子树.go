package main

import (
	"fmt"
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
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return false
	}

	/*
		思路:
		1.二叉树用前序遍历方式序列化成字符串
		2.判断子树序列化后的字符串是否是root数序列化后的字符串的子串
		3.注意如果用前序遍历序列化,可能存在root的根结点的值是12,
		subRoot的根节点的值为2时, 判断错误的情况,所以序列化时在前面加上一个特殊字符,比如a
		就会变成 a12 和 a2 不会像等

	*/
	var strTree func(str string, root *TreeNode) string

	strTree = func(str string, root *TreeNode) string {
		// if root == nil {
		// 	return str
		// }
		str += fmt.Sprintf("a%d;", root.Val)
		if root.Left == nil {
			str += "#;"
		} else {
			str = strTree(str, root.Left)
		}

		if root.Right == nil {
			str += "#;"
		} else {
			str = strTree(str, root.Right)
		}
		return str
	}

	rootStr := strTree("", root)
	subRootStr := strTree("", subRoot)
	fmt.Printf("rootStr=%s\nsubRootStr=%s", rootStr, subRootStr)
	return strings.Contains(rootStr, subRootStr)
}

func isSubtree1(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return false
	}

	/*
		思路:
		1.二叉树用后序遍历方式序列化成字符串
		2.判断子树序列化后的字符串是否是root数序列化后的字符串的子串
	*/
	var strTree func(str string, root *TreeNode) string

	strTree = func(str string, root *TreeNode) string {
		// if root == nil {
		// 	return str
		// }
		if root.Left == nil {
			str += "#;"
		} else {
			str = strTree(str, root.Left)
		}

		if root.Right == nil {
			str += "#;"
		} else {
			str = strTree(str, root.Right)
		}
		str += fmt.Sprintf("%d;", root.Val)
		return str
	}

	rootStr := strTree("", root)
	subRootStr := strTree("", subRoot)
	// fmt.Printf("rootStr=%s\nsubRootStr=%s", rootStr, subRootStr)
	return strings.Contains(rootStr, subRootStr)
}

// @lc code=end
