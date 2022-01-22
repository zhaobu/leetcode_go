/*
 * @lc app=leetcode.cn id=559 lang=golang
 *
 * [559] Maximum Depth of N-ary Tree
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
解法1: dfs
*/
func maxDepth1(root *Node) int {
	if root == nil {
		return 0
	}
	var (
		retDepth = 0
		dfs      func(node *Node, depth int)
	)
	dfs = func(node *Node, depth int) {
		if node == nil {
			return
		}
		if depth > retDepth {
			retDepth = depth
		}
		for _, n := range node.Children {
			dfs(n, depth+1)
		}
	}

	dfs(root, 1)
	return retDepth
}

/*
解法2: bfs
*/
func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	var (
		queue     = []*Node{root}
		depth     = 0
		nextLevel []*Node
	)
	for len(queue) > 0 {
		nextLevel = make([]*Node, 0, len(queue)*3)
		for _, n := range queue {
			//可以在此访问节点
			nextLevel = append(nextLevel, n.Children...)
		}
		depth++
		queue = nextLevel
	}

	return depth
}

// @lc code=end

