package main

import (
	. "leetcode/check/tree"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=449 lang=golang
 *
 * [449] 序列化和反序列化二叉搜索树
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
解法1 通用方式
*/
// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	str := strings.Builder{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		str.WriteString(strconv.Itoa(root.Val))
		str.WriteString(",")
		dfs(root.Right)
	}

	dfs(root)
	return str.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	nums := strings.Split(data, ",")
	nums = nums[:len(nums)-1]

	var bst func(i, j int) *TreeNode
	bst = func(i, j int) *TreeNode {
		if i > j {
			return nil
		}
		mid := i + (j-1)>>1
		v, _ := strconv.Atoi(nums[mid])
		root := &TreeNode{
			Val:   v,
			Left:  bst(i, mid-1),
			Right: bst(mid+1, j),
		}
		return root
	}

	return bst(0, len(nums)-1)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
// @lc code=end
