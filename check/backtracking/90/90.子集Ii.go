package main

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */

// @lc code=start

/*
解法1 backtrack
参考 https://labuladong.gitee.io/algo/4/30/108/
*/
func subsetsWithDup1(nums []int) [][]int {
	if len(nums) < 1 {
		return nil
	}
	// 先排序nums数组
	sort.Ints(nums)

	ret := [][]int{}
	var dfs func(record []int, start int)
	dfs = func(record []int, start int) {

		ret = append(ret, append([]int{}, record...))
		for i := start; i < len(nums); i++ {
			//在[start, len(nums)) 元素不能重复
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			record = append(record, nums[i])
			dfs(record, i+1)
			record = record[:len(record)-1]
		}
	}
	dfs(make([]int, 0, len(nums)), 0)

	return ret
}

/*
解法2 枚举2^n -1
官方解法1 https://leetcode-cn.com/problems/subsets-ii/solution/zi-ji-ii-by-leetcode-solution-7inq/
*/
func subsetsWithDup(nums []int) [][]int {
	if len(nums) < 1 {
		return nil
	}
	// 先排序nums数组
	sort.Ints(nums)

	n := len(nums)
	ret := [][]int{}
	record := make([]int, 0, n)
	m := 1 << n
flag:
	for i := 0; i < m; i++ {
		record = make([]int, 0, n)
		for j, v := range nums {
			if i&(1<<j) > 0 {
				/*
					对于当前选择的数 x，若前面有与其相同的数 y，且没有选择 y，
					此时包含 x 的子集，必然会出现在包含 y 的所有子集中
				*/
				if j > 0 && i&(1<<(j-1)) == 0 && nums[j-1] == v {
					continue flag
				}
				record = append(record, v)
			}
		}
		ret = append(ret, record)
	}

	return ret
}

// @lc code=end
