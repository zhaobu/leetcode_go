package main

import "sort"

/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	ret := [][]int{}
	if len(nums) < 1 {
		return ret
	}
	sort.Ints(nums)
	used := make([]bool, len(nums)) //用来标记那些元素已使用
	var dfs func(pickedNums []int)
	dfs = func(pickedNums []int) {
		if len(pickedNums) == len(nums) {
			ret = append(ret, append([]int{}, pickedNums...))
			return
		}
		for i, v := range nums {
			if used[i] || (i > 0 && nums[i] == nums[i-1] && used[i-1]) {
				continue
			}
			used[i] = true
			dfs(append(pickedNums, v))
			used[i] = false
		}
	}

	dfs(make([]int, 0, len(nums)))

	return ret
}

// @lc code=end
