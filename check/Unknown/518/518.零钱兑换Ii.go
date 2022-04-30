package main

/*
 * @lc app=leetcode.cn id=518 lang=golang
 *
 * [518] 零钱兑换 II
 */

// @lc code=start

/*
解法1 动态规划
*/
func change(amount int, coins []int) int {
	if amount == 0 {
		return 1
	}

	dp := make([]int, amount+1)
	dp[0] = 1

	for _, v := range coins {
		for i := v; i <= amount; i++ {
			if dp[i-v] > 0 {
				dp[i] += dp[i-v]
			}
		}
	}
	// fmt.Printf("dp=%+v\n", dp)
	return dp[amount]
}

// @lc code=end
