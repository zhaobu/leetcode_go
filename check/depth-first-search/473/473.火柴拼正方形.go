package main

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=473 lang=golang
 *
 * [473] 火柴拼正方形
 */

// @lc code=start
/*
解法1 回溯
*/
func makesquare1(matchsticks []int) bool {
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

	/*
		1. dfs表示当前正在防止第i根火柴
		2. 每根火柴都有4种放置方式
		3. 因为前面已经过滤掉总和不能被4整除的情况,
		以及有单根火柴长度已经大于边长的情况,所以只要能放完所有火柴
		一定就是正方形
	*/
	var dfs func(i int)
	dfs = func(i int) {
		if ret {
			return
		}
		if i == n {
			ret = true
			// fmt.Printf("boards=%+v\n", boards)
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

/*
解法2 动态规划
*/
func makesquare2(matchsticks []int) bool {
	totalLen := 0
	for _, l := range matchsticks {
		totalLen += l
	}
	if totalLen%4 != 0 {
		return false
	}

	tLen := totalLen / 4
	dp := make([]int, 1<<len(matchsticks))
	for i := 1; i < len(dp); i++ {
		// fmt.Printf("%b\n", i)
		dp[i] = -1
	}
	for s := 1; s < len(dp); s++ {
		for k, v := range matchsticks {
			if s>>k&1 == 0 {
				continue
			}
			s1 := s &^ (1 << k)
			if dp[s1] >= 0 && dp[s1]+v <= tLen {
				dp[s] = (dp[s1] + v) % tLen
				break
			}
		}
	}
	return dp[len(dp)-1] == 0
}

/*
解法3 二进制枚举
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

	s := make([]int, 1<<n)
	/*
		1.预处理,s[mask]表示mask这种枚举状态下所有被选中的元素总和
		2. mask=0时表示所有的火柴都不选,所以mask至少>=1
	*/
	for mask := 1; mask < 1<<n; mask++ {
		//方法1:直接求和
		/*
			for i, v := range matchsticks {
				if mask&(1<<i) > 0 {
					s[mask] += v
				}
			}
		*/
		//方法2: 动态规划,找到最低位的1,由前面转移而来
		i := 0
		for (mask & (1 << i)) == 0 {
			i++
		}
		s[mask] = s[mask^1<<i] + matchsticks[i]
	}

	for a := 1; a < 1<<n; a++ {
		if s[a] != board {
			continue
		}
		x := (1<<n - 1) ^ a
		for b := x; b > 0; b = (b - 1) & x {
			if s[b] != board {
				continue
			}
			y := x ^ b
			for c := y; c > 0; c = (c - 1) & y {
				if s[c] != board {
					continue
				}
				return true
			}
		}
	}
	return false
}

// @lc code=end
