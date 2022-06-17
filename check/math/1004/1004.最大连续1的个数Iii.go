package main

import "sort"

/*
 * @lc app=leetcode.cn id=1004 lang=golang
 *
 * [1004] 最大连续1的个数 III
 */

// @lc code=start
func longestOnes1(nums []int, k int) (ans int) {
	n := len(nums)
	P := make([]int, n+1)
	for i, v := range nums {
		P[i+1] = P[i] + 1 - v
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for right, v := range P {
		left := sort.SearchInts(P, v-k)
		ans = max(ans, right-left)
	}
	return
}

func longestOnes(nums []int, k int) int {
	//1变0, 0变1, 并且计算前缀和
	preSum := make([]int, len(nums)+1)
	for i := range nums {
		preSum[i+1] = preSum[i] + 1 - nums[i]
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	ret := 0
	//对于每一个right.计算最小的left使[left,right]之间的1的个数<=k
	// fmt.Printf("nums=%+v\n", nums)

	for right := 0; right < len(preSum); right++ {
		i, j := 0, right
		for i < j {
			mid := i + (j-i)>>1
			/*
				区间[mid,right)范围内1的个数
			*/
			cnt := preSum[right] - preSum[mid]
			// fmt.Printf("preSum=%+v,cnt=%d, mid=%d, right=%d\n", preSum, cnt, mid, right)
			if cnt > k {
				i = mid + 1
			} else {
				j = mid
			}
		}
		ret = max(ret, right-j)
	}

	return ret
}

// @lc code=end
