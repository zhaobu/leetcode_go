package main

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=149 lang=golang
 *
 * [149] 直线上最多的点数
 */

// @lc code=start

/*
解法1 暴力法
1. 当我们枚举到点 i 时，我们只需要考虑编号大于 i 的点到点 i 的斜率，
因为如果直线同时经过编号小于点 i 的点 j，那么当我们枚举到 j 时就已经考虑过该直线了
2. 在点的总数量小于等于 2 的情况下，我们总可以用一条直线将所有点串联，此时我们直接返回点的总数量即可
3. 当我们找到一条直线经过了图中超过半数的点时，我们即可以确定该直线即为经过最多点的直线；
4. 当我们枚举到点 i（假设编号从 0 开始）时，我们至多只能找到 n−i 个点共线。
假设此前找到的共线的点的数量的最大值为 k，如果有 k >=n-i，那么此时我们即可停止枚举，
因为不可能再找到更大的答案了。
*/
func maxPoints1(points [][]int) int {
	m := len(points)
	if m <= 2 { //点个数<3时,直接可得出答案
		return m
	}
	ret := 2 //点个数2时,最少有2个点在一条直线上
	for i := 0; i < m; i++ {
		p1 := points[i]
		for j := i + 1; j < m; j++ {
			p2 := points[j]
			count := 2
			for k := j + 1; k < m; k++ {
				p3 := points[k]
				s1 := (p2[1] - p1[1]) * (p3[0] - p1[0])
				s2 := (p2[0] - p1[0]) * (p3[1] - p1[1])
				if s1 == s2 {
					count++
				}
			}
			if count > 2 && count > ret {
				//优化3
				if count > m>>1 {
					return count
				}
				//优化4
				if ret >= m-i {
					return ret
				}
				ret = count
			}
		}
	}
	return ret
}

/*
解法2 暴力枚举法优化
*/
func maxPoints(points [][]int) int {
	m := len(points)
	if m <= 2 { //点个数<3时,直接可得出答案
		return m
	}
	//求最大公约数
	var gcd func(a, b int) int
	gcd = func(a, b int) int {
		if b == 0 {
			return a
		}
		return gcd(b, a%b)
	}
	ret := 0 //点个数2时,最少有2个点在一条直线上
	for i := 0; i < m; i++ {
		p1 := points[i]
		max := 0                   //与points[i]在同一条直线上的点的最大个数
		record := map[string]int{} //以points[i]为起点的斜率相同的直线的个数
		for j := i + 1; j < m; j++ {
			p2 := points[j]
			y := p2[1] - p1[1]
			x := p2[0] - p1[0]
			//求最大公约数
			t := gcd(x, y)
			key := fmt.Sprintf("%d%d", x/t, y/t)
			record[key] += 1
			count := record[key]
			if count > max {
				max = count
			}
		}
		//优化3
		max += 1 //加上points[i]自身
		if max > m>>1 {
			return max
		}
		//优化4
		if ret >= m-i {
			return ret
		}
		if max > ret {
			ret = max
		}
	}
	return ret
}

// @lc code=end
