package Solution

import (
	"sort"
)

func removeCoveredIntervals1(intervals [][]int) int {
	// 按照起点升序排列，起点相同时降序排列
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		} else {
			return intervals[i][1] > intervals[j][1]
		}
	})
	left, right := intervals[0][0], intervals[0][1]
	count := 0
	for i := 1; i < len(intervals); i++ {
		arr := intervals[i]
		//情况一，找到覆盖区间
		if left <= arr[0] && right >= arr[1] {
			count++
		}
		// 情况二，找到相交区间，合并
		if left <= arr[0] && right <= arr[1] {
			right = arr[1]
		}

		if right <= arr[0] {
			left = arr[0]
			right = arr[1]
		}
	}
	return len(intervals) - count
}

func removeCoveredIntervals2(intervals [][]int) int {
	if len(intervals) == 0 || len(intervals[0]) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		} else {
			return intervals[i][1] > intervals[j][1]
		}
	})
	result := 0
	max := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] == intervals[i-1][0] {
			result++
			continue
		}
		if max < intervals[i][1] {
			max = intervals[i][1]
		} else {
			result++
		}
	}

	return len(intervals) - result
}
