package main

import "sort"

/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start

/*
解法1 backtrack
*/
func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) < 1 {
		return nil
	}

	//先进行排序,让相同的元素靠在一起
	sort.Ints(candidates)

	ret := [][]int{}
	var dfs func(record []int, start, sum int)
	dfs = func(record []int, start, sum int) {
		if sum == target {
			ret = append(ret, append([]int{}, record...))
			return
		}
		if sum > target {
			return
		}
		for i := start; i < len(candidates); i++ {
			//剪枝逻辑，值相同的同一层的树枝，只遍历第一条
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			record = append(record, candidates[i])
			dfs(record, i+1, sum+candidates[i])
			record = record[:len(record)-1]
		}
		return
	}
	dfs(make([]int, 0, len(candidates)), 0, 0)

	return ret
}

// @lc code=end
