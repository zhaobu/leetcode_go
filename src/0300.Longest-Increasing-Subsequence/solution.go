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
