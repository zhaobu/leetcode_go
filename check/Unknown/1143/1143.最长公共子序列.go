package main

/*
 * @lc app=leetcode.cn id=1143 lang=golang
 *
 * [1143] 最长公共子序列
 */

// @lc code=start

/*
解法1
递归
解法超时,只为展示该种思路
*/
func longestCommonSubsequence1(text1 string, text2 string) int {
	if len(text1) < 1 || len(text2) < 1 {
		return 0
	}

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var lcs func(s1, s2 []byte) int
	lcs = func(s1, s2 []byte) int {
		if len(s1) < 1 || len(s2) < 1 {
			return 0
		}
		l1, l2 := len(s1)-1, len(s2)-1
		if s1[l1] == s2[l2] {
			return lcs(s1[:l1], s2[:l2]) + 1
		}
		return max(lcs(s1[:l1], s2), lcs(s1, s2[:l2]))
	}

	return lcs([]byte(text1), []byte(text2))
}

/*
解法2
动态规划
状态定义:
dp[j][j] 表示 text1 的前 i 个元素和 text2 的前 j 个元素的最长公共子序列长度
初始状态:
当 i=0 或者 j=0 时,表示一个字符串和空串的最长公共子序列为 0, 也就是dp[i][j]=0
状态转移方程:
1. 当 text1[i] == text2[j] 时,dp[i][j]=dp[i-1][j-1]+1
比如 text1=abcd 和 text2=bfd时 dp[4][3]对应的最长公共子序列是 bd
dp[3][2]对应的最长公共子序列是 b, 加上最后一个公共字符 d 就是dp[4][3]

2. 当 text1[i] != text2[j] 时, dp[i][j] = max(dp[i-1][j],dp[i][j-1])
比如 text1=abcd 和 text2=agc 时 dp[4][3]对应的最长公共子序列是 ac
dp[3][2]对应的最长公共子序列是 b

*/
func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) < 1 || len(text2) < 1 {
		return 0
	}

	//dp[i][j] 表示text1[0,i]和text2[0,j]的最长公共子序列的长度
	dp := make([][]int, len(text1)+1)

	for i, _ := range dp {
		dp[i] = make([]int, len(text2)+1)
	}
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	//dp[i][0] 和dp[0][j] 都为0,不用初始化
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}

// @lc code=end
