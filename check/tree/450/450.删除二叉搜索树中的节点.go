package main

import (
	. "leetcode/check/tree"
)

/*
 * @lc app=leetcode.cn id=450 lang=golang
 *
 * [450] 删除二叉搜索树中的节点
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
解法1 迭代(删除后继节点)

情况1: 如果要删除的节点没有孩子节点,直接删除该节点,也就是删除父节点指向该节点的指针
情况2: 如果要删除的节点只有一个子节点（只有左子节点或者右子节点），只需要更新父节点中，
指向要删除节点的指针，让它指向要删除节点的子节点.
具体的就是:
1. 如果要删除的节点只有左孩子，只需要更新父节点中，指向要删除节点的指针，让它指向要删除节点的左孩子
2. 如果要删除的节点只有右孩子，只需要更新父节点中，指向要删除节点的指针，让它指向要删除节点的右孩子

情况3: 如果要删除的节点有两个子节点,需要找到这个节点的右子树中的最小节点(后继节点)，把它替换到要删除的节点上。
然后再删除掉这个最小节点，因为最小节点肯定没有左子节点（如果有左子结点，那就不是最小节点了）,
这时删除后继节点也就变成情况2了

情况3也可以写成找到左子树中最大的节点(前驱),然后后面步骤一样
*/
func deleteNode1(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	parent := (*TreeNode)(nil) //指向要删除节点的父节点
	delNode := root            //指向要删除的节点
	for delNode != nil && delNode.Val != key {
		parent = delNode
		if key > delNode.Val {
			delNode = delNode.Right
		} else {
			delNode = delNode.Left
		}
	}
	//没有找到要删除的节点
	if delNode == nil {
		return root
	}

	// 先处理情况3
	if delNode.Left != nil && delNode.Right != nil {
		//找到要删除节点的后继节点
		parent = delNode
		successor := delNode.Right
		for successor.Left != nil {
			parent = successor
			successor = successor.Left
		}
		//交换要删除的节点和后继节点上的值,相当于把后继节点和要删除的节点交换
		delNode.Val = successor.Val
		//变成了删除后继节点
		delNode = successor
	}

	/*
		1. 获取要删除节点的孩子节点
		2. 要删除的节点最多只有一个孩子节点
	*/
	child := (*TreeNode)(nil)
	if delNode.Left != nil {
		child = delNode.Left
	} else if delNode.Right != nil {
		child = delNode.Right
	}

	if parent == nil { //删除的是根节点并且根节点最多只有一个子节点
		return child
	}
	if parent.Left == delNode {
		//要删除的节点是父节点的左孩子
		parent.Left = child
	} else {
		//要删除的节点是父节点的右孩子
		parent.Right = child
	}

	return root
}

/*
解法2  迭代(删除前驱节点)
*/
func deleteNode2(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	parent := (*TreeNode)(nil) //指向要删除节点的父节点
	delNode := root            //指向要删除的节点
	for delNode != nil && delNode.Val != key {
		parent = delNode
		if key > delNode.Val {
			delNode = delNode.Right
		} else {
			delNode = delNode.Left
		}
	}
	//没有找到要删除的节点
	if delNode == nil {
		return root
	}

	// 先处理情况3
	if delNode.Left != nil && delNode.Right != nil {
		//找到要删除节点的前驱节点
		parent = delNode
		successor := delNode.Left
		for successor.Right != nil {
			parent = successor
			successor = successor.Right
		}
		//交换要删除的节点和后继节点上的值,相当于把后继节点和要删除的节点交换
		delNode.Val = successor.Val
		//变成了删除后继节点
		delNode = successor
	}

	/*
		1. 获取要删除节点的孩子节点
		2. 要删除的节点最多只有一个孩子节点
	*/
	child := (*TreeNode)(nil)
	if delNode.Left != nil {
		child = delNode.Left
	} else if delNode.Right != nil {
		child = delNode.Right
	}

	if parent == nil { //删除的是根节点并且根节点最多只有一个子节点
		return child
	} else {
		if parent.Left == delNode {
			//要删除的节点是父节点的左孩子
			parent.Left = child
		} else {
			//要删除的节点是父节点的右孩子
			parent.Right = child
		}
	}

	return root
}

/*
解法3 递归
递归函数返回的是删除以root为根的树上的值为key的节点后的树
*/
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		// 处理情况3,获得右子树最小的节点(后继节点)
		successor := root.Right
		for successor.Left != nil {
			successor = successor.Left
		}
		/*
			1. 删除当前节点后,当前节点的后继节点应该交换到当前节点的位置
			2. 所以可以先删除当前节点的后继节点,然后把后继节点交换到当前节点位置

		*/
		successor.Right = deleteNode(root.Right, successor.Val)
		successor.Left = root.Left
		root = successor
	}

	return root
}

// @lc code=end
