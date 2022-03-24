package main

/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) < 1 {
		return nil
	}

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
			record = append(record, candidates[i])
			/*
				1. start保证了当前层只能从[start,len(candidates))之间选元素
				2. 当前分支的下一层也是只能从[start,len(candidates))之间选元素
			*/
			dfs(record, i, sum+candidates[i])
			record = record[:len(record)-1]
		}
		return
	}
	dfs(make([]int, 0, len(candidates)), 0, 0)

	return ret
}

// @lc code=end
