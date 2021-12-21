package main

/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	if nums == nil || len(nums) < 1 {
		return
	}

	insertPos := 0 //下次元素移动到该位置
	for k, v := range nums {
		if v != 0 {
			nums[insertPos], nums[k] = nums[k], nums[insertPos]
			insertPos++
		}
	}

	return
}

// 10001203
// @lc code=end

// https://www.yuque.com/docs/share/5e71189b-b21c-479f-90de-6f46e555eab8?# 《283.移动零》
