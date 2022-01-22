package main

/*
 * @lc app=leetcode.cn id=590 lang=golang
 *
 * [590] N 叉树的后序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

/*
解法1:递归
*/
func postorder1(root *Node) []int {
	ret := []int{}
	if root == nil {
		return ret
	}
	var (
		bfs func(node *Node)
	)
	bfs = func(node *Node) {
		if node == nil {
			return
		}
		for _, n := range node.Children {
			bfs(n)
		}
		ret = append(ret, node.Val)
	}
	bfs(root)
	return ret
}

/*
解法2:迭代,复用二叉树的后续遍历思想
*/
func postorder2(root *Node) []int {
	ret := []int{}
	if root == nil {
		return ret
	}

	stack := []*Node{root}
	preNode := (*Node)(nil) //记录上一个访问的节点

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		childrenNums := len(top.Children)
		//当前栈顶没有子节点,或者最后一个子节点已经访问过则出栈,并且进行访问
		if childrenNums == 0 || top.Children[childrenNums-1] == preNode {
			ret = append(ret, top.Val)
			preNode = top
			stack = stack[:len(stack)-1]
		} else {
			//否则将子节点加入栈中
			for i := childrenNums - 1; i >= 0; i-- {
				stack = append(stack, top.Children[i])
			}
		}
	}

	return ret
}

// @lc code=end
