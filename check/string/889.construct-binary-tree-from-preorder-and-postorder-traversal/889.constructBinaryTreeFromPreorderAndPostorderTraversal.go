package main

/*
 * @lc app=leetcode.cn id=889 lang=golang
 *
 * [889] Construct Binary Tree from Preorder and Postorder Traversal
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
 解法1:递归
*/
func constructFromPrePost1(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	if len(preorder) == 1 {
		return root
	}

	/*
	  preorder: 0 * 1 2 3 4 5 * 6 7 8 9  root*left*right
	 postorder: 1 2 3 4 5 * 6 7 8 9 * 0  left*right*root

	 1.root节点在前序遍历中为 preorder[0],在后序遍历中为 postorder[len(postorder)-1]
	 2.root节点的左孩子在前序遍历中为 preorder[1],假设左子树节点个数为L,则在后序遍历中为
	 postorder[L-1]
	*/
	var (
		leftRootVal   = preorder[1] //root节点左孩子的值
		leftRootindex = 0           //root节点左孩子在后序遍历数组中的下标
	)
	for i, v := range postorder {
		if v == leftRootVal {
			leftRootindex = i
			break
		}
	}
	root.Left = constructFromPrePost(preorder[1:2+leftRootindex], postorder[:1+leftRootindex])
	root.Right = constructFromPrePost(preorder[2+leftRootindex:], postorder[1+leftRootindex:len(postorder)-1])
	return root
}

/*
 解法2:递归 用map优化
*/
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	hashTable := make(map[int]int)
	// 在后序遍历序列postorder中建立元素和位置的映射关系
	for i, v := range postorder {
		hashTable[v] = i
	}
	// l1,r1,l2,r2分别代表前序遍历序列和后序遍历序列的首尾位置
	var buildTree func(int, int, int, int) *TreeNode
	buildTree = func(l1, r1, l2, r2 int) *TreeNode {
		// 递归终止条件
		if l1 > r1 || l2 > r2 {
			return nil
		}
		// l1初始值为0，那么preorder[l1]一定是对应根节点的。
		root := &TreeNode{Val: preorder[l1]}
		if l1 == r1 {
			return root
		}
		// 先序遍历序列preorder是根左右,preorder[l1]对应根节点，那么preorder[l1+1]一定对应左子树根节点
		// 确定左子树根节点在后序遍历序列postorder中的位置
		pos := hashTable[preorder[l1+1]]
		// 求root节点左子树的长度(后序遍历序列postorder是左右根，那么左子树长度即为pos-l2+1, l2初始值为0)
		leftLength := pos - l2 + 1
		// 确定左子树范围
		root.Left = buildTree(l1+1, l1+leftLength, l2, pos)
		// 确定右子树范围，右子树范围可以根据左子树范围推出
		root.Right = buildTree(l1+leftLength+1, r1, pos+1, r2-1)
		return root
	}
	return buildTree(0, len(preorder)-1, 0, len(postorder)-1)
}

// @lc code=end
