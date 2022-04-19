package main

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start

/*
解法1 按照结束元素升序排列
*/
func merge1(intervals [][]int) [][]int {
	n := len(intervals)
	if n < 2 {
		return intervals
	}

	//先把区间按照结束元素升序排列
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	// fmt.Printf("intervals=%+v\n", intervals)
	ret := [][]int{}
	cur := intervals[n-1] //当前已合并区间
	/*
		在处理每一个 intervals 时必须从后往前,
		如果不是存在类似[[2,3],[4,5],[6,7],[8,9],[1,10]]这种,本应该合并后只有[1,10],
		但求得的结果是[[2,3],[4,5],[6,7],[1,10]]

	*/
	for i := n - 2; i >= 0; i-- {
		v := intervals[i]
		if v[1] >= cur[0] {
			if v[0] < cur[0] {
				cur[0] = v[0]
			}
		} else {
			ret = append(ret, cur)
			cur = intervals[i]
		}
	}

	ret = append(ret, cur)
	return ret
}

/*
解法2 按照开始元素升序排列
*/
func merge(intervals [][]int) [][]int {
	n := len(intervals)
	if n < 2 {
		return intervals
	}

	//先把区间按照开始元素升序排列
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// fmt.Printf("intervals=%+v\n", intervals)
	ret := [][]int{}
	cur := intervals[0] //当前已合并区间

	/*
		原始数组:[[2,3],[4,5],[6,7],[8,9],[1,10]]
		排序后[[1 10] [2 3] [4 5] [6 7] [8 9]]
	*/
	for i := 1; i < n; i++ {
		v := intervals[i]
		if v[0] > cur[1] {
			ret = append(ret, cur)
			cur = intervals[i]
		} else if v[1] > cur[1] {
			cur[1] = v[1]
		}
	}

	ret = append(ret, cur)
	return ret
}

// @lc code=end
