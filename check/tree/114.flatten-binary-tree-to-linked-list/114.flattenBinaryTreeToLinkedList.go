package main

/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] Flatten Binary Tree to Linked List
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
 解法1: 前序遍历后再展开 dfs
*/
func flatten1(root *TreeNode) {
	if root == nil {
		return
	}

	var (
		nodeList []*TreeNode //存储前序遍历的结果
		dfs      func(node *TreeNode)
	)
	//前序遍历二叉树
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		nodeList = append(nodeList, node)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	//将二叉树展开为链表
	preNode := nodeList[0]
	for i := 1; i < len(nodeList); i++ {
		preNode.Left = nil
		preNode.Right = nodeList[i]
		preNode = nodeList[i]
	}
	return
}

/*
 解法2: 前序遍历后再展开 迭代写法,其实和解法1是一种,只是遍历二叉树方式不一样
*/
func flatten2(root *TreeNode) {
	if root == nil {
		return
	}

	var (
		nodeList []*TreeNode //存储前序遍历的结果
		stack    []*TreeNode
		node     = root
	)
	//前序遍历二叉树
	for node != nil || len(stack) > 0 {
		for node != nil {
			nodeList = append(nodeList, node)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node = node.Right
	}
	//将二叉树展开为链表
	preNode := nodeList[0]
	for i := 1; i < len(nodeList); i++ {
		preNode.Left = nil
		preNode.Right = nodeList[i]
		preNode = nodeList[i]
	}
	return
}

/*
 解法3: 使用全部进栈后出栈才访问的迭代方式进行前序遍历,遍历过程中就处理为链表
*/
func flatten3(root *TreeNode) {
	if root == nil {
		return
	}

	var (
		stack   = []*TreeNode{root}
		preNode *TreeNode //上一个访问的节点
	)
	//前序遍历二叉树
	for len(stack) > 0 {
		curNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if curNode.Right != nil {
			stack = append(stack, curNode.Right)
		}
		if curNode.Left != nil {
			stack = append(stack, curNode.Left)
		}
		if preNode != nil {
			preNode.Left = nil
			preNode.Right = curNode
		}
		preNode = curNode
	}
	return
}

/*
 解法4: 使用全部进栈后出栈才访问的迭代方式进行前序遍历,遍历过程中就处理为链表
*/
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	var (
		stack   = []*TreeNode{root}
		preNode *TreeNode //上一个访问的节点
	)
	//前序遍历二叉树
	for len(stack) > 0 {
		curNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if curNode.Right != nil {
			stack = append(stack, curNode.Right)
		}
		if curNode.Left != nil {
			stack = append(stack, curNode.Left)
		}
		if preNode != nil {
			preNode.Left = nil
			preNode.Right = curNode
		}
		preNode = curNode
	}
	return
}

// @lc code=end
