package main

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

// 方法4:Morris 中序遍历
// func inorderTraversal4(root *TreeNode) []int {
// 	if root == nil {
// 		return nil
// 	}
// 	var (
// 		res   []int
// 		node  = root
// 		stack []*TreeNode
// 	)
// 	return
// }

// @lc code=end
