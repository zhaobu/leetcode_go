package main

import "sort"

/*
 * @lc app=leetcode.cn id=1235 lang=golang
 *
 * [1235] 规划兼职工作
 */

// @lc code=start
/*
解法1 排序+动态规划
*/
func jobScheduling(startTime []int, endTime []int, profit []int) int {
	type job struct {
		startTime int
		endTime   int
		profit    int
	}
	jobs := make([]*job, 0, len(startTime))
	for i, _ := range startTime {
		jobs = append(jobs, &job{
			startTime: startTime[i],
			endTime:   endTime[i],
			profit:    profit[i],
		})
	}

	sort.Slice(jobs, func(i, j int) bool { return jobs[i].endTime <= jobs[j].endTime })

	//dp[i] 表示第i件工作结束后获得的最大报酬
	dp := make([]int, len(jobs))
	dp[0] = jobs[0].profit
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 1; i < len(jobs); i++ {
		//如果选择做改建工作
		pre := i
		for j := i - 1; j >= 0; j-- {
			if jobs[j].endTime <= jobs[i].startTime {
				pre = j
				break
			}
		}
		dp[i] = max(dp[i-1], dp[pre]+jobs[i].profit)
	}
	return dp[len(jobs)-1]
}

// @lc code=end
