package main

import "sort"

/*
 * @lc app=leetcode.cn id=436 lang=golang
 *
 * [436] 寻找右区间
 */

// @lc code=start

/*
解法1 二分查找
*/
func findRightInterval(intervals [][]int) []int {
	n := len(intervals)
	if n == 1 {
		return []int{-1}
	}

	for i := range intervals {
		intervals[i] = append(intervals[i], i)
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})

	ret := make([]int, len(intervals))

	search := func(end int) int {
		i, j := 0, len(intervals)
		for i <= j {
			mid := i + (j-i)>>1
			if intervals[mid][0] < end {
				i = mid + 1
			} else {
				if mid == 0 || intervals[mid-1][0] < end {
					return mid
				}
				j = mid - 1
			}
		}
		return -1
	}
	for i := range intervals {
		oldIdx, end := intervals[i][2], intervals[i][1]
		j := search(end)
		if j == -1 {
			ret[oldIdx] = -1
		} else {
			ret[oldIdx] = intervals[j][2]
		}
	}
	return ret
}

/*
解法2 双指针解法
*/

// @lc code=end
