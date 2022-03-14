package main

import (
	. "leetcode/check/tree"
)

/*
 * @lc app=leetcode.cn id=99 lang=golang
 *
 * [99] 恢复二叉搜索树
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
 解法1 递归中序遍历求解
*/
func recoverTree1(root *TreeNode) {
	if root == nil {
		return
	}
	/*
	   321
	   123

	   1324
	   1234
	*/

	/*
		1. 第一个要交换的元素是第1个逆序对中较大的元素
		2. 第二个要交换的元素是第2个逆序对中较小的元素
		2. 分别记录第一个和第二个要交换的元素,以及中序遍历上一个访问的节点
	*/
	var n1, n2, lastNode *TreeNode
	var travel func(node *TreeNode)

	travel = func(node *TreeNode) {
		if node == nil {
			return
		}
		travel(node.Left)
		// fmt.Printf("node=%+v\n", node.Val)
		// fmt.Printf("node=%+v, lastNode=%+v \n", node, lastNode)
		if lastNode != nil && node.Val < lastNode.Val {
			if n1 == nil {
				n1 = lastNode
			}
			/*
				存在只有一个逆序对的情况,所以n2始终记录为较小元素
			*/
			n2 = node
			// fmt.Printf("n1=%+v, n2= %+v\n",n1,n2)
		}
		lastNode = node
		travel(node.Right)

		return
	}

	travel(root)
	// fmt.Printf("n1=%+v, n2= %+v\n",n1,n2)
	n1.Val, n2.Val = n2.Val, n1.Val //交换2个节点的值
	return
}

/*
	解法2 迭代中序遍历求解
*/

func recoverTree2(root *TreeNode) {
	if root == nil {
		return
	}
	/*
	   321
	   123

	   1324
	   1234
	*/

	/*
		1. 第一个要交换的元素是第1个逆序对中较大的元素
		2. 第二个要交换的元素是第2个逆序对中较小的元素
		2. 分别记录第一个和第二个要交换的元素,以及中序遍历上一个访问的节点
	*/
	var n1, n2, lastNode *TreeNode
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if lastNode != nil && root.Val < lastNode.Val {
			if n1 == nil {
				n1 = lastNode
			}
			/*
				存在只有一个逆序对的情况,所以n2始终记录为较小元素
			*/
			n2 = root
			// fmt.Printf("n1=%+v, n2= %+v\n",n1,n2)
		}
		lastNode = root

		root = root.Right
	}

	n1.Val, n2.Val = n2.Val, n1.Val //交换2个节点的值
	return
}

/*
	解法2 mirrors 中序遍历求解
*/

func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}
	/*
	   321
	   123

	   1324
	   1234
	*/

	/*
		1. 第一个要交换的元素是第1个逆序对中较大的元素
		2. 第二个要交换的元素是第2个逆序对中较小的元素
		2. 分别记录第一个和第二个要交换的元素,以及中序遍历上一个访问的节点
	*/
	var n1, n2, lastNode *TreeNode

	var check = func(root *TreeNode) {
		if lastNode != nil && root.Val < lastNode.Val {
			if n1 == nil {
				n1 = lastNode
			}
			/*
				存在只有一个逆序对的情况,所以n2始终记录为较小元素
			*/
			n2 = root
			// fmt.Printf("n1=%+v, n2= %+v\n",n1,n2)
		}
		// fmt.Printf("n1=%+v, n2= %+v\n", n1, n2)
		lastNode = root
	}

	for root != nil {
		if root.Left != nil {
			// 找到前驱节点
			preNode := root.Left
			for preNode.Right != nil && preNode.Right != root {
				preNode = preNode.Right
			}
			if preNode.Right == nil {
				preNode.Right = root
				root = root.Left
			} else if preNode.Right == root {
				preNode.Right = nil
				check(root)
				root = root.Right
			}
		} else {
			check(root)
			root = root.Right
		}
	}

	n1.Val, n2.Val = n2.Val, n1.Val //交换2个节点的值
	return
}

// @lc code=end
