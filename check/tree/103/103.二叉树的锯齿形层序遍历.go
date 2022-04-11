package main

import (
	. "leetcode/check/tree"
)

/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层序遍历
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
 解法1 层序遍历
1. 在存储每层结果时用flag判断下标计算方式模拟是从左向右还是从右向左
2. 比用双端队列,时间复杂度或空间复杂度上更好
3. 不需要反转
*/
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	ret := [][]int{}
	queue := []*TreeNode{root}
	flag := true                 //true表示要从右往左输出,false表示从左往右输出
	count := 1                   //当前层剩下的节点个数
	curRet := make([]int, count) //当前层的遍历结果

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		if top.Right != nil {
			queue = append(queue, top.Right)
		}
		if top.Left != nil {
			queue = append(queue, top.Left)
		}

		if flag { //从右往左输出
			curRet[count-1] = top.Val
		} else { //从左往右输出
			curRet[len(curRet)-count] = top.Val
		}

		count--
		if count == 0 {
			flag = !flag
			count = len(queue)
			ret = append(ret, curRet)
			curRet = make([]int, count)
		}
	}

	return ret
}

// @lc code=end
