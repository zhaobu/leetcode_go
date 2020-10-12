package Solution

/*
[leetcode 官方 方法一：动态规划)](https://leetcode-cn.com/problems/maximum-subarray/solution/zui-da-zi-xu-he-by-leetcode-solution/)
*/
func LengthOfLIS1(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	dp := make([]int, len(nums))
	maxAns := 0
	for k0, v0 := range nums {
		dp[k0] = 1
		for i := 0; i < k0; i++ {
			if nums[i] < v0 {
				dp[k0] = max(dp[k0], dp[i]+1)
			}
		}
		maxAns = max(maxAns, dp[k0])
	}
	return maxAns
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// [动态规划+二分查找](https://leetcode-cn.com/problems/longest-increasing-subsequence/solution/zui-chang-shang-sheng-zi-xu-lie-dong-tai-gui-hua-e/)

func LengthOfLIS2(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	dp := make([]int, 0, len(nums)) //用于保存最长上升子序列
	for _, v := range nums {
		// 如果 dp 中元素都比它小，将它插到最后
		if len(dp) == 0 || dp[len(dp)-1] < v {
			dp = append(dp, v)
			continue
		}
		//二分查找dp，用它覆盖掉比它大的元素中最小的那个
		start, end := 0, len(dp)-1
		for start < end {
			mid := (start + end) >> 1
			if dp[mid] < v {
				start = mid + 1
			} else {
				end = mid
			}
		}
		dp[start] = v
	}
	return len(dp)
}
