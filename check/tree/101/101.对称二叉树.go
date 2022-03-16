package main

import (
	. "leetcode/check/tree"
)

/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
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
 解法1 递归
 root节点点的左子树和右子树对称
 用n1表示左子树的根节点,n2表示右子树的根节点
 可以从一颗数的第2层看第三层
 n1的左子节点和n2的右子节点,n1的右子节点和n2的左子节点都相等
*/
func isSymmetric1(root *TreeNode) bool {
	var check func(n1, n2 *TreeNode) bool

	check = func(n1, n2 *TreeNode) bool {
		if n1 == nil || n2 == nil {
			return false
		}
		if n1 == nil && n2 == nil {
			return true
		}
		if n1.Val != n2.Val {
			return false
		}

		return check(n1.Left, n2.Right) && check(n1.Right, n2.Left)
	}

	return check(root, root)
}

/*
解法2 层序遍历
1. 每一层判断是否对称
2. 注意空节点也要加入
*/
func isSymmetric2(root *TreeNode) bool {

	if root == nil {
		return true
	}
	//判断是否对称
	check := func(queue []*TreeNode) bool {
		for i, j := 0, len(queue)-1; i < j; i, j = i+1, j-1 {
			if queue[i] != nil && queue[j] != nil {
				if queue[i].Val != queue[j].Val {
					return false
				}
			} else if queue[i] != queue[j] {
				return false
			}
		}
		return true
	}

	queue := []*TreeNode{root}
	levelNum := 1
	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]
		if head != nil {
			if head.Left != nil {
				queue = append(queue, head.Left)
			} else {
				queue = append(queue, nil)
			}
			if head.Right != nil {
				queue = append(queue, head.Right)
			} else {
				queue = append(queue, nil)
			}
		}
		levelNum--
		if levelNum == 0 {
			if !check(queue) {
				return false
			}
			levelNum = len(queue)
		}
	}
	return true
}

/*
解法3 层序遍历
1. 官方解法2
2. 队列中存在2颗一样的树 first 和 second , first 从左往右加入, second 从右往左加入
3. 每次加入时是轮流加入,不是一颗树加入完后再加入另一颗
4. 每次出队列时,出2个元素.判断他们是否相等
*/

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := []*TreeNode{root, root}
	levelNum := 2
	for len(queue) > 0 {
		first, second := queue[0], queue[1]
		queue = queue[2:]
		if first == nil && second == nil {
			continue
		}
		if first == nil || second == nil {
			return false
		}

		if first.Val != second.Val {
			return false
		}
		queue = append(queue, first.Left)
		queue = append(queue, second.Right)
		queue = append(queue, first.Right)
		queue = append(queue, second.Left)
		levelNum -= 2
		if levelNum == 0 {
			levelNum = len(queue)
		}
	}
	return true
}

// @lc code=end
