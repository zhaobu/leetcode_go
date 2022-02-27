package main

/*
 * @lc app=leetcode.cn id=654 lang=golang
 *
 * [654] 最大二叉树
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
 解法1
 使用递归方式创建所有节点
 注意切片截取时,截取开始的下标的取值范围是[0,len(nums)], 而不是[0,len(nums)-1]
 超出[0,len(nums)]都会报错
*/
func constructMaximumBinaryTree1(nums []int) *TreeNode {
	var root *TreeNode
	if len(nums) < 1 {
		return root
	}
	var getRoot func(nums []int) *TreeNode

	getRoot = func(nums []int) *TreeNode {
		if nums == nil || len(nums) == 0 {
			return nil
		}
		maxIdx := 0
		for i, v := range nums {
			if v > nums[maxIdx] {
				maxIdx = i
			}
		}
		// fmt.Printf("maxIdx=%d, nums=%+v,nums[:maxIdx]=%+v, nums[maxIdx+1:]=%+v \n", maxIdx, nums, nums[:maxIdx], nums[maxIdx+1:])
		root := &TreeNode{Val: nums[maxIdx]}
		root.Left = getRoot(nums[:maxIdx])
		root.Right = getRoot(nums[maxIdx+1:])
		return root
	}

	root = getRoot(nums)
	return root
}

// @lc code=end
