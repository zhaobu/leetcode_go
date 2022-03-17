package main

import (
	"fmt"
	. "leetcode/check/tree"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=297 lang=golang
 *
 * [297] 二叉树的序列化与反序列化
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

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	ret := ""
	if root == nil {
		return ret
	}

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			ret += "#,"
			return
		}
		ret += fmt.Sprintf("%d,", root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}

	dfs(root)
	// fmt.Printf("ret=%v\n", ret)
	return ret
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	strArr := strings.Split(data, ",")
	var buildTree func() *TreeNode
	buildTree = func() *TreeNode {
		cur := strArr[0]
		strArr = strArr[1:]
		if cur == "#" {
			return nil
		}
		v, _ := strconv.Atoi(cur)
		left := buildTree()
		right := buildTree()
		node := &TreeNode{
			Val:   v,
			Left:  left,
			Right: right,
		}
		fmt.Printf("node=%+v\n", node)
		return node
	}

	fmt.Printf("strArr=%+v\n", strArr)
	return buildTree()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
// @lc code=end
