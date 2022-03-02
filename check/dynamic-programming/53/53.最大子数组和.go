/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */

// @lc code=start

/*
解法1
动态规划
优化空间复杂度
*/
func maxSubArray1(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	//dp[i] 表示以nums[i] 结尾的最大和连续子数组的和
	dp := nums[0]
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		sum := dp + nums[i]
		if sum > nums[i] {
			dp = sum
		} else {
			dp = nums[i]
		}
		if dp > maxSum {
			maxSum = dp
		}
	}

	return maxSum
}

/*
解法2
动态规划
优化判断条件
*/
func maxSubArray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	//dp[i] 表示以nums[i] 结尾的最大和连续子数组的和
	dp := nums[0]
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		if dp > 0 {
			dp = dp + nums[i] //可以比方法1少一部分计算加法的情况
		} else {
			dp = nums[i]
		}
		if dp > maxSum {
			maxSum = dp
		}
	}

	return maxSum
}

// @lc code=end

