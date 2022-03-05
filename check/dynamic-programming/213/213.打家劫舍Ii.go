package main

/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */
/*
解法1 动态规划

状态定义:
和第198题一样,dp[i]表示下标为[0,i]范围内的房屋能偷窃到的最高金额

basecase:
因为首位相连,所以第 0 间和第 len(nums)-1 间房屋只能二选1偷
所以如果偷第 0 间:
dp[0] = nums[0], dp[1]=max(nums[0],nums[1])
如果偷第 len(nums)-1 间:
dp[1] = nums[1], dp[2]=max(nums[1],nums[2])

转移方程
和之前一样
dp[i] = max(dp[i-1], dp[i-2]+nums[i])
如果偷第 i 间: dp[i] = dp[i-2]+nums[i], 也是就是第[0,i-2]间房最高金额+第 i 间房的金额
如果不偷第 i 间: dp[i] = dp[i-1] 也是就是第[0,i-1]间房最高金额
只不过是i的初始值不一样
*/
// @lc code=start
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

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	// dp[i]表示下标为[0,i]范围内的房屋能偷窃到的最高金额
	dp := make([]int, len(nums))

	start := []int{0, 1} //偷和不偷第0间房屋时,开始下标不一样

	maxValue := 0 //选择第0间房屋和第len-1间房屋偷, 者当中的较大值
	for _, k := range start {
		dp[k] = nums[k] //从下标为k的房屋开始偷
		dp[k+1] = max(nums[k], nums[k+1])
		end := k + len(nums) - 2 //不管开始偷窃的下标多少,结束的下标都是start+len(nums)-2
		for i := k + 2; i <= end; i++ {
			dp[i] = max(dp[i-1], dp[i-2]+nums[i])
		}
		maxValue = max(maxValue, dp[end])
	}

	return maxValue
}

/*
解法2 动态规划 空间复杂度优化
根据状态转移方程可以得知
dp[i] 只和 dp[i-1] 和dp[i-2] 相关,所以可以用3个变量替代dp数组
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

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	// dp[i]表示下标为[0,i]范围内的房屋能偷窃到的最高金额
	dp := 0

	start := []int{0, 1} //偷和不偷第0间房屋时,开始下标不一样

	maxValue := 0 //选择第0间房屋和第len-1间房屋偷, 者当中的较大值
	for _, k := range start {
		dp2 := nums[k] //从下标为k的房屋开始偷
		dp1 := max(nums[k], nums[k+1])
		end := k + len(nums) - 2 //不管开始偷窃的下标多少,结束的下标都是start+len(nums)-2
		for i := k + 2; i <= end; i++ {
			dp = max(dp1, dp2+nums[i])
			dp2 = dp1
			dp1 = dp
		}
		maxValue = max(maxValue, dp1)
	}

	return maxValue
}

// @lc code=end
