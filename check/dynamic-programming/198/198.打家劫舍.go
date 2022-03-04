package main

/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start

/*
解法1 动态规划
状态定义: dp[i]表示前i+1间房能偷到的最高总金额

basecase:
dp[0] = nums[0] 只有一间房屋，则偷窃该房屋
dp[1] = max(nums[0], nums[1]) 只有两间房屋，选择其中金额较高的房屋进行偷窃

状态转移:
​dp[i]=max(dp[i−2]+nums[i],dp[i−1])
1. 偷窃第 i 间房屋，那么就不能偷窃第 i-1 间房屋，偷窃总金额为前 i-2 间房屋的最高总金额与第 i 间房屋的金额之和
2. 不偷窃第 i 间房屋，偷窃总金额为前 i-1 间房屋的最高总金额。

*/
func rob1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	/*
	   dp[i]表示前i+1间房能偷到的最高金额
	*/
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		a := nums[i] + dp[i-2] //偷第nums[i]间房屋
		b := dp[i-1]           //不偷第nums[i]间房屋
		dp[i] = max(a, b)
	}

	return dp[len(nums)-1]
}

/*
解法2 动态规划 优化空间复杂度
*/
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	/*
	   dp[i]表示前i+1间房能偷到的最高金额
	*/
	dp := 0
	dp2 := nums[0]               //偷到第i-2间房屋时最高总金额
	dp1 := max(nums[0], nums[1]) //偷到第i-1间房屋时最高总金额

	for i := 2; i < len(nums); i++ {
		dp = max(dp1, nums[i]+dp2)
		dp2 = dp1
		dp1 = dp
	}

	//不返回dp而返回dp1是为了兼容len(nums)=2时的情况
	return dp1
}

// @lc code=end
