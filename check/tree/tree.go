package tree

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
	Str   string
}

//不带null节点的参与运算的树
func BuildTree(str string) (root *TreeNode) {
	if len(str) < 3 {
		return
	}
	arr := strings.Split(str[1:len(str)-1], ",")
	/*
		  	  0
		  1      2
		3   4  5    6
	*/

	var dfs func(i int) *TreeNode
	dfs = func(i int) *TreeNode {
		if i >= len(arr) || arr[i] == "null" {
			return nil
		}
		nodeVal, _ := strconv.Atoi(arr[i])
		return &TreeNode{Val: nodeVal,
			Left:  dfs(2*i + 1),
			Right: dfs(2*i + 2),
			Str:   arr[i],
		}
	}
	root = dfs(0)
	return root
}

//带null节点的打印的树
func BuildStrTree(str string) (root *TreeNode) {
	if len(str) < 3 {
		return
	}
	arr := strings.Split(str[1:len(str)-1], ",")
	/*
		  	  0
		  1      2
		3   4  5    6
	*/

	var dfs func(i int) *TreeNode
	dfs = func(i int) *TreeNode {
		if i >= len(arr) {
			return nil
		}
		nodeVal, _ := strconv.Atoi(arr[i])
		return &TreeNode{Val: nodeVal,
			Left:  dfs(2*i + 1),
			Right: dfs(2*i + 2),
			Str:   arr[i],
		}
	}
	root = dfs(0)
	return root
}

func (root *TreeNode) String() (str string) {
	str += "["
	if root != nil {
		queue := []*TreeNode{root}
		for len(queue) > 0 {
			head := queue[0]
			queue = queue[1:]
			if head.Str == "null" {
				str += head.Str + ","
			} else {
				str += strconv.Itoa(head.Val) + ","
			}
			if head.Left != nil {
				queue = append(queue, head.Left)
			}
			if head.Right != nil {
				queue = append(queue, head.Right)
			}
		}
		str = str[:len(str)-1]
	}
	str += "]"
	return str
}
