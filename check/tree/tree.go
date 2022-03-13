package tree

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(str string) (root *TreeNode) {
	if str == "" || len(str) < 1 {
		return
	}
	arr := []byte(str)
	/*
		  	  0
		  1      2
		3   4  5    6
	*/

	levelNode := make([]*TreeNode, len(arr))
	for i := 0; 2*i+2 < len(arr); i++ {
		if arr[i] == '#' {
			continue
		}
		nodeVal, _ := strconv.Atoi(string(arr[i]))
		if levelNode[i] == nil {
			levelNode[i] = &TreeNode{Val: nodeVal}
		}
		if arr[2*i+1] != '#' {
			nodeVal, _ = strconv.Atoi(string(arr[2*i+1]))
			levelNode[2*i+1] = &TreeNode{Val: nodeVal}
			levelNode[i].Left = levelNode[2*i+1]
		}
		if arr[2*i+2] != '#' {
			nodeVal, _ = strconv.Atoi(string(arr[2*i+2]))
			levelNode[2*i+2] = &TreeNode{Val: nodeVal}
			levelNode[i].Right = levelNode[2*i+2]
		}
	}

	return levelNode[0]
}
