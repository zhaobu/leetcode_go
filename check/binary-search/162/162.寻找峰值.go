package main

/*
 * @lc app=leetcode.cn id=162 lang=golang
 *
 * [162] 寻找峰值
 */

// @lc code=start

/*
解法1 直接求得第一个波峰
求第一个nums[i] > nums[i+1]的值,所以在区间[0,i]上必然是递增的
*/
func findPeakElement(nums []int) int {
	n := len(nums)
	if n < 2 { //n必定>=1
		return 0
	}

	for i := range nums {
		if i+1 < n && nums[i] > nums[i+1] {
			return i
		}
	}
	return n - 1
}

// @lc code=end
