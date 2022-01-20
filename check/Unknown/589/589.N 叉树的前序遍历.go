/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N 叉树的前序遍历
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
func preorder1(root *Node) []int {
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
		ret = append(ret, node.Val)
		for _, n := range node.Children {
			bfs(n)
		}
	}
	bfs(root)
	return ret
}

/*
解法2:迭代, 出栈时访问元素,并且从后往前依次加入孩子节点
*/
func preorder(root *Node) []int {
	ret := []int{}
	if root == nil {
		return ret
	}

	stack := []*Node{root}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, top.Val)
		for i := len(top.Children) - 1; i >= 0; i-- {
			stack = append(stack, top.Children[i])
		}
	}
	return ret
}

// @lc code=end

