package main

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=1200 lang=golang
 *
 * [1200] 最小绝对差
 */

// @lc code=start
func minimumAbsDifference(arr []int) [][]int {
	n := len(arr)
	if n < 2 {
		return nil
	}

	sort.Ints(arr)
	//找到最小绝对差
	minVal := arr[1] - arr[0]
	for i := 2; i < n; i++ {
		if arr[i]-arr[i-1] < minVal {
			minVal = arr[i] - arr[i-1]
		}
	}

	ret := [][]int{}

	for i := 1; i < n; i++ {
		if arr[i]-arr[i-1] == minVal {
			ret = append(ret, []int{arr[i-1], arr[i]})
		}
	}

	return ret
}

// @lc code=end
