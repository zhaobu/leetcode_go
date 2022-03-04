package main

import "fmt"

/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */

// @lc code=start

/*
解法1 动态规划
状态定义: dp[i] 表示以下标i结尾的最长子序列的长度
basecase:
dp[0] = 1 表示以nums[0] 结尾的最长子序列就是nums[0]
状态转移方程:
在求dp[i]时,遍历 j ∈ [0,i)
1. 当 nums[j] > nums[i] 时,跳过
2. 当 nums[j] <= nums[i] 时说明nums[i]可以接在nums[j]后面
dp[i] = dp[j] + 1
*/
func lengthOfLIS1(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	//dp[i] 表示以下标i结尾的最长子序列的长度
	dp := make([]int, len(nums))
	dp[0] = 1
	maxLen := 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1 //默认初始化为 最长上升子序列只有nums[i]
		for j := 0; j < i; j++ {
			if nums[j] >= nums[i] {
				continue
			}
			if dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
		// fmt.Printf("dp[%d]=%+v\n", i, dp[i])
	}

	return maxLen
}

/*
解法2 贪心
如果要使上升子序列尽可能的长，则需要让序列上升得尽可能慢，
因此应该让每次在上升子序列最后加上的那个数尽可能的小

遍历nums数组,维护一个数组arrs,arr[i]表示遍历到nums[i]位置时
得到的所有长度为i+1的上升子序列当中,nums[i]是最小的

*/
func lengthOfLIS2(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	arrs := make([]int, 0, len(nums))
	for _, v := range nums {
		curIdx := -1
		// 查找curIdx的位置可以用二分查找
		for j, a := range arrs {
			if a >= v {
				curIdx = j
				arrs[j] = v
				break
			}
		}
		if curIdx == -1 {
			arrs = append(arrs, v)
		}
		fmt.Printf("arrs=%+v\n", arrs)
	}

	return len(arrs)
}

/*
解法2 贪心 + 二分查找
优化查找 curIdx 的位置

*/
func lengthOfLIS(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	//查找第一个大于等于给定值的元素的下标
	var binarySearchFirstGT = func(a []int, v int) int {
		if len(a) < 1 {
			return -1
		}
		low, high := 0, len(a)-1
		for low <= high {
			mid := low + (high-low)>>1
			if a[mid] >= v {
				if mid == 0 || a[mid-1] < v { //当a[mid]>=v时,判断mid左边的值是否<v
					return mid
				} else {
					high = mid - 1
				}
			} else {
				low = mid + 1
			}
		}
		return -1
	}

	arrs := make([]int, 0, len(nums))
	for _, v := range nums {
		//二分查找第一个>=v的下标mid
		curIdx := binarySearchFirstGT(arrs, v)
		if curIdx >= 0 && curIdx < len(arrs) {
			arrs[curIdx] = v
		} else {
			arrs = append(arrs, v)
		}
		fmt.Printf("arrs=%+v\n", arrs)
	}

	return len(arrs)
}

// @lc code=end
