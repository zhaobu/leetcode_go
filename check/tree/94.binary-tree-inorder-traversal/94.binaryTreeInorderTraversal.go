package main

import (
	. "leetcode/check/tree"
)

/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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

//方法1:递归
func inorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var (
		res     []int
		inorder func(node *TreeNode)
	)

	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			inorder(node.Left)
		}
		res = append(res, node.Val)
		if node.Right != nil {
			inorder(node.Right)
		}
	}
	inorder(root)
	return res
}

//方法2:迭代
func inorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var (
		res   []int
		node  = root
		stack []*TreeNode
	)
	for node != nil || len(stack) > 0 {
		//模拟所有左节点入栈
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		//获取栈顶元素
		node = stack[len(stack)-1]
		//中序遍历是第二次到达时访问
		res = append(res, node.Val)
		//遍历当前节点的右子节点
		node = node.Right
		//栈顶元素出栈
		stack = stack[:len(stack)-1]
	}
	return res
}

/*
 方法3:迭代
 思路:从根节点开始,一直遍历左子树,并依次入栈,直到最左边叶子节点没有左孩子时开始出栈
第一个出栈的就是最左边叶子节点,

*/
func inorderTraversal3(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var (
		res   []int
		node  = root
		stack []*TreeNode
	)
	for {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
		} else {
			if len(stack) < 1 {
				return res
			}

			node = stack[len(stack)-1]
			res = append(res, node.Val)
			node = node.Right
			stack = stack[:len(stack)-1]
		}
	}
}

/*
方法4:Morris 中序遍历
中序遍历的结果中每个ret[i]的前驱节点就是ret[i-1]
所以在遍历到每个节点时,先把该节点的前驱节点的right指向自己
表示遍历完前驱节点后下一个要遍历的节点就是自己
*/
func inorderTraversal(root *TreeNode) []int {
	ret := []int{}
	if root == nil {
		return ret
	}

	for root != nil {
		if root.Left != nil {
			//找到当前节点的前驱节点
			preNode := root.Left
			/*
			 必须加上preNode.Right != root,因为第二次找前驱节点时
			 前驱节点的right已经指向自己
			*/
			for preNode.Right != nil && preNode.Right != root {
				preNode = preNode.Right
			}

			if preNode.Right == nil { //如果前驱节点的right为空,就指向自己.并处理左节点
				preNode.Right = root
				root = root.Left
			} else if preNode.Right == root { //如果发现前驱节点的right已经指向自己,相当于是第二次访问,直接打印
				preNode.Right = nil
				ret = append(ret, root.Val)
				root = root.Right
			}
		} else {
			/*
				1. 如果节点的左孩子为空表示到了叶子结点,就直接打印
				2. 此时rigth已经指向中序遍历的下一个节点.所以直接遍历right
			*/
			ret = append(ret, root.Val)
			root = root.Right
		}
	}
	return ret
}

// @lc code=end
