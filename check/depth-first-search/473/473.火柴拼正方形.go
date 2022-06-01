package main

import "sort"

/*
 * @lc app=leetcode.cn id=473 lang=golang
 *
 * [473] 火柴拼正方形
 */

// @lc code=start
/*
解法1 回溯
*/
func makesquare(matchsticks []int) bool {
	n := len(matchsticks)
	if n < 4 {
		return false
	}
	//求出正方形的边长
	board := 0
	for _, v := range matchsticks {
		board += v
	}
	if board%4 != 0 {
		return false
	}
	board /= 4
	// fmt.Printf("board=%v\n", board)
	//过滤有火柴单根已经超过边长的情况
	for _, v := range matchsticks {
		if v > board {
			return false
		}
	}

	//排序
	sort.Sort(sort.Reverse(sort.IntSlice(matchsticks)))

	//回溯出一种能拼成正方形的方法
	ret := false
	boards := [4]int{}

	var dfs func(i int)
	dfs = func(i int) {
		if ret {
			return
		}
		if i == n {
			ret = true
			return
		}
		for j := 0; j < 4; j++ {
			boards[j] += matchsticks[i]
			if boards[j] <= board {
				dfs(i + 1)
			}
			boards[j] -= matchsticks[i]
		}
	}

	dfs(0)

	// fmt.Printf("used=%+v,usedCount=%d \n", used, usedCount)
	return ret
}

// @lc code=end
