package main

/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] Sort Colors
 */

// @lc code=start
func sortColors(nums []int) {
	if len(nums) < 1 {
		return
	}
	leftIdx := 0              //已排好序的最后一个0的位置
	rightIdx := len(nums) - 1 //已排好序的第一个2的位置
	i := 0                    //当前遍历到的位置
	for i <= rightIdx {
		if nums[i] == 0 {
			nums[leftIdx], nums[i] = nums[i], nums[leftIdx]
			leftIdx++
			i++
		} else if nums[i] == 1 {
			i++
		} else {
			nums[rightIdx], nums[i] = nums[i], nums[rightIdx]
			rightIdx--
		}
	}
	return
}

// @lc code=end
