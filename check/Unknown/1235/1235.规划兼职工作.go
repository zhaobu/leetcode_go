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
func jobScheduling1(startTime []int, endTime []int, profit []int) int {
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
		/*
			1. 从i-1开始到0 逆序找到endtime不晚于jobs[i]的starttime的第一份工作
			2. 如果找不到前一份工作,但jobs[i]又必须做,则赚到的钱就是jobs[i].profit
			3. pre默认初始化为i,这样可以保证找不到时dp[pre]=0.
		*/
		pre := i
		for j := i - 1; j >= 0; j-- {
			if jobs[j].endTime <= jobs[i].startTime {
				pre = j
				break
			}
		}
		//从不做第i件工作和做第i件工作中找到最大值
		dp[i] = max(dp[i-1], dp[pre]+jobs[i].profit)
	}
	return dp[len(jobs)-1]
}

/*
解法2 解法1的优化,二分查找前一份工作
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

	//二分查找最后一个不大于target的元素
	var bSearch = func(i, j, target int) int {

		for i <= j {
			mid := i + (j-i)>>1
			if jobs[mid].endTime > target {
				j = mid - 1
			} else {
				if mid == j || jobs[mid+1].endTime > target {
					return mid
				}
				i = mid + 1
			}
		}

		return -1
	}

	for i := 1; i < len(jobs); i++ {
		/*
			1. 从i-1开始到0 逆序找到endtime不晚于jobs[i]的starttime的第一份工作
			2. 如果找不到前一份工作,但jobs[i]又必须做,则赚到的钱就是jobs[i].profit
			3. pre默认初始化为i,这样可以保证找不到时dp[pre]=0.
		*/
		pre := i
		k := bSearch(0, i-1, jobs[i].startTime)
		if k != -1 {
			pre = k
		}

		//从不做第i件工作和做第i件工作中找到最大值
		dp[i] = max(dp[i-1], dp[pre]+jobs[i].profit)
	}
	return dp[len(jobs)-1]
}

// @lc code=end
