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
func inorderTraversal2(root *TreeNode) (res []int) {
	if root == nil {
		return nil
	}

	var (
		stack []*TreeNode
		node  = root
	)
	for node != nil || len(stack) > 0 {
		//模拟所有左节点入栈
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		//栈顶出栈
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 打印出栈元素
		res = append(res, node.Val)
		//遍历当前节点的右子节点
		node = node.Right
	}
	return
}

/*
 方法3:迭代
 思路:从根节点开始,一直遍历左子树,并依次入栈,直到最左边叶子节点没有左孩子时开始出栈
第一个出栈的就是最左边叶子节点,

*/
func inorderTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return nil
	}
	var (
		stack []*TreeNode
		node  = root
	)
	for {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
		} else {
			if len(stack) < 1 {
				return res
			}
			//访问栈顶元素
			node = stack[len(stack)-1]
			res = append(res, node.Val)

			//栈顶元素出栈
			stack = stack[:len(stack)-1]

			//下一个访问的是栈顶元素的右子树
			node = node.Right
		}
	}
}

// 方法4:Morris 中序遍历
func inorderTraversal4(root *TreeNode) (res []int) {
	if root == nil {
		return nil
	}
	return
}

// @lc code=end
