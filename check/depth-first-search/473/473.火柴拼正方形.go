package main

import (
	"fmt"
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

	m := 1 << n
	s := make([]int, m)
	/*
		1. 要枚举个数为n的情况, 二进制的位数要为2^n
		2.预处理,s[mask]表示mask这种枚举状态下所有被选中的元素总和
		3. mask=0时表示所有的火柴都不选,所以mask至少>=1
	*/
	for mask := 1; mask < m; mask++ {
		//方法1:直接求和
		/*
			for i, v := range matchsticks {
				if mask&(1<<i) > 0 {
					s[mask] += v
				}
			}
		*/
		/*
			方法2: 动态规划,找到最低位的1,由前面转移而来
			1. mask^(1<<i): 表示消掉mask的最地位的1,等价于mask&(mask-1)
			2. 因为只消掉最低位的1,所以s[mask]和s[mask^(1<<i)]只相差选取第i位上的火柴
		*/
		i := 0
		for (mask & (1 << i)) == 0 {
			i++
		}
		// fmt.Printf("mask=%05b,pre=%05b,other=%05b\n", mask, mask^1<<i, mask&(mask-1))
		s[mask] = s[mask^(1<<i)] + matchsticks[i]
	}

	/*
		1. a, b, c, d 为四个边所选的元素
		2. 边a可以从 [1,m - 1] 中情况中任选一种
		3. 边b可以从 [1,m - 1] 中出掉a选的那种情况外选一种,也就是(m - 1) ^ a 的子集
		4. c 为 (m - 1)^a^b 的子集
		5. d 为(m - 1)^a^b^c
		即可保证 a | b | c | d = (1 << n) - 1 用完所有的元素
	*/

	aSet := m - 1 //边a的选取范围
	for a := 1; a < m; a++ {
		fmt.Printf("A1 aSet=%05b, a=%05b\n", aSet, a)
		if s[a] != board {
			continue
		}
		bSet := aSet ^ a
		for b := bSet; b > 0; b = (b - 1) & bSet {
			fmt.Printf("B1 aSet=%05b, a=%05b, bSet=%05b, b=%05b\n", aSet, a, bSet, b)
			if s[b] != board {
				continue
			}
			cSet := bSet ^ b
			for c := cSet; c > 0; c = (c - 1) & cSet {
				fmt.Printf("C1 aSet=%05b, a=%05b, bSet=%05b, b=%05b, cSet=%05b, c=%05b\n", aSet, a, bSet, b, cSet, c)
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
