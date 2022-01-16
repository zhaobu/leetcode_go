package main

/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
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
//解法1:递归
func postorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var (
		res       []int
		postorder func(node *TreeNode)
	)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			postorder(node.Left)
		}
		if node.Right != nil {
			postorder(node.Right)
		}
		res = append(res, node.Val)
	}
	postorder(root)
	return res
}

/*
解法2:迭代
*/
func postorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var (
		res     []int
		node    = root
		stack   []*TreeNode
		preNode *TreeNode //记录上次访问的节点,用来判断是不是第2次出栈,也就是第三次到达节点
	)

	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		//获取栈顶元素
		node = stack[len(stack)-1]
		/*
			后序遍历是当第三次到达时才访问
			只有右边没有节点或者是上一个访问的节点是当前节点的右节点时,才表示是第三次到达
		*/
		if node.Right == nil || node.Right == preNode {
			res = append(res, node.Val)
			preNode = node
			node = nil
			stack = stack[:len(stack)-1]
		} else {
			//如果不是第三次到达,也不用出栈,就先访问右子树
			node = node.Right
		}
	}
	return res
}

/*
 方法3:迭代,上一种写法的另一种形式
*/
func postorderTraversal3(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var (
		res     []int
		node    = root
		stack   []*TreeNode
		preNode *TreeNode //上次访问的节点

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
			if node.Right == nil || node.Right == preNode {
				res = append(res, node.Val)
				preNode = node
				node = nil
				stack = stack[:len(stack)-1]
			} else {
				node = node.Right
			}
		}
	}
}

// @lc code=end
