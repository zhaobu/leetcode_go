package main

import "fmt"

/*
 * @lc app=leetcode.cn id=516 lang=golang
 *
 * [516] 最长回文子序列
 */

// @lc code=start

/*
解法1 动态规划
dp[i][j] 表示s[i,j]字符串之间的最长回文子序列长度

basecase:
dp[i][i] = 1

转移方程:
如果s[i] == s[j]
dp[i][j] = dp[i+1][j-1] + 2
如果s[i] != s[j]
dp[i][j] = max(dp[i][j-1], dp[i+1][j])

注意i和j的遍历顺序由状态转移方程决定
*/
func longestPalindromeSubseq1(s string) int {
	n := len(s)
	if n < 2 {
		return n
	}
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// dp[i][j] s[i,j]之间的字符串的最长回文子序列长度
	dp := make([][]int, n)
	for i, _ := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}
	for i := n - 1; i >= 0; i-- { //由转移方程决定i要从大到小求解
		for j := i + 1; j < n; j++ { //由转移方程决定j要从小到大求解
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i+1][j])
			}
		}
		// fmt.Printf("dp[%d]=%+v\n", i, dp[i])
	}

	return dp[0][n-1]
}

/*
解法2 解法1的状态压缩
*/
func longestPalindromeSubseq(s string) int {
	n := len(s)
	if n < 2 {
		return n
	}
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// dp[i][j] s[i,j]之间的字符串的最长回文子序列长度
	dp := make([][]int, n)
	for i, _ := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}
	for i := n - 1; i >= 0; i-- { //由转移方程决定i要从大到小求解
		for j := i + 1; j < n; j++ { //由转移方程决定j要从小到大求解
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i+1][j])
			}
		}
		fmt.Printf("dp[%d]=%+v\n", i, dp[i])
	}

	return dp[0][n-1]
}

// @lc code=end
