package main

/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] Symmetric Tree
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
func isSymmetric(root *TreeNode) bool {

	queue := []*TreeNode{root}
	nodeNil := &TreeNode{}
	for len(queue) > 0 {

		//判断当前层是否对称
		for i, j := 0, len(queue)-1; i < j; i, j = i+1, j-1 {
			if queue[i] == nodeNil && queue[j] != nodeNil ||
				queue[j] == nodeNil && queue[i] != nodeNil {
				return false
			}
			if queue[i].Val != queue[j].Val {
				return false
			}
		}
		//生成下一层的节点
		nextLevel := make([]*TreeNode, 0, len(queue)*2)
		for _, node := range queue {
			if node == nodeNil {
				continue
			}
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			} else {
				nextLevel = append(nextLevel, nodeNil)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			} else {
				nextLevel = append(nextLevel, nodeNil)
			}
		}
		queue = nextLevel
	}
	return true
}

// @lc code=end
