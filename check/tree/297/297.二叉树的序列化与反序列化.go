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

/*
解法1 前序遍历
*/
// Serializes a tree to a single string.
func (this *Codec) serialize1(root *TreeNode) string {
	if root == nil {
		return ""
	}

	str := strings.Builder{} //比直接字符串相加效率更高
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			str.WriteString("#,")
			return
		}
		str.WriteString(strconv.Itoa(root.Val))
		str.WriteString(",")
		dfs(root.Left)
		dfs(root.Right)
	}

	dfs(root)
	// fmt.Printf("ret=%v\n", ret)
	return str.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize1(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	strArr := strings.Split(data, ",")
	strArr = strArr[:len(strArr)-1] //去掉逗号分割后最后一个空值
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

/*
解法2 后序遍历
*/
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	str := strings.Builder{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			str.WriteString("#,")
			return
		}
		dfs(root.Left)
		dfs(root.Right)
		str.WriteString(strconv.Itoa(root.Val))
		str.WriteString(",")
	}

	dfs(root)
	// fmt.Printf("ret=%v\n", ret)
	return str.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	strArr := strings.Split(data, ",")

	strArr = strArr[:len(strArr)-1] //去掉逗号分割后最后一个空值
	var buildTree func() *TreeNode
	buildTree = func() *TreeNode {
		n := len(strArr)
		if n == 0 {
			return nil
		}
		/*
			1. 后序遍历反序列化要从后往前,依次是root, right, left
		*/
		cur := strArr[n-1]
		strArr = strArr[:n-1]
		if cur == "#" {
			return nil
		}
		v, _ := strconv.Atoi(cur)
		//后序遍历反序列化要先求右孩子
		right := buildTree()
		left := buildTree()
		node := &TreeNode{
			Val:   v,
			Left:  left,
			Right: right,
		}
		// fmt.Printf("node=%+v\n", node)
		return node
	}

	// fmt.Printf("strArr=%+v\n", strArr)
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
