package main

/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
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
//解法1:递归实现
func preorderTraversal1(root *TreeNode) []int {
	var (
		res      []int
		preorder func(node *TreeNode)
	)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		if node.Left != nil {
			preorder(node.Left)
		}
		if node.Right != nil {
			preorder(node.Right)
		}
	}
	preorder(root)
	return res
}

/*
解法2:迭代,手动维护栈
思路:由于栈是“先进后出”的顺序，所以入栈时先将右子树入栈，这样使得前序遍历结果为 “根->左->右”的顺序。
*/
func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var (
		res   []int
		stack = []*TreeNode{root}
	)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		res = append(res, top.Val)
		stack = stack[:len(stack)-1]
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}
	return res
}

/*
解法3:迭代,手动维护栈
思路:访问一个节点,然后将节点入栈,访问左子树,知道访问到了最左边的叶子节点
然后，弹出一个栈顶元素,访问它的右孩子，直至栈为空
*/
func preorderTraversal3(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var (
		res   []int
		stack = []*TreeNode{}
		node  = root
	)
	for node != nil || len(stack) > 0 {
		for node != nil { //根节点和左孩子入栈
			res = append(res, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right //每弹出一个元素,就到达右孩子
		stack = stack[:len(stack)-1]
	}
	return res
}

/*
解法4:迭代,手动维护栈
思路是先访问每个节点,然后保存节点的右子树,然后在访问节点的左子树,
这样到达树的最左边的叶子节点时开始从栈中访问最后入栈的右子树
*/
func preorderTraversal(root *TreeNode) []int {
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
			res = append(res, node.Val)
			stack = append(stack, node) //访问一个节点时,用栈保存右子树部分
			node = node.Left
		} else {
			if len(stack) < 1 {
				return res
			}
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			node = node.Right
		}
	}
}

// @lc code=end
