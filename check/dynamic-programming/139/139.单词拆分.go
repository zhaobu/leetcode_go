package main /*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */

// @lc code=start

/*
解法1 自顶向下备忘录
*/
func wordBreak1(s string, wordDict []string) bool {
	m := len(s)
	dp := make([]int, m)

	/*
		返回s[i:]能否被拼接
	*/
	var dfs func(i int) bool
	dfs = func(i int) bool {
		if i == m {
			return true
		}
		if dp[i] != 0 {
			return dp[i] == 1
		}
		for _, v := range wordDict {
			n := len(v)
			if i+n <= m && s[i:i+n] == v {
				if dfs(i + n) {
					dp[i] = 1
					return true
				}
			}
		}

		dp[i] = -1
		return false
	}

	return dfs(0)
}

/*
解法2 动态规划
*/
func wordBreak(s string, wordDict []string) bool {
	m := len(s)

	dp := make([]int, m+1) //dp[i]表示s[:i]能否被拼接
	dp[0] = 1              //空字符串能直接拼接出

	//自底向上求解
	for i := 1; i <= m; i++ {
		dp[i] = -1
		for _, v := range wordDict {
			n := len(v)
			if i >= n && s[i-n:i] == v && dp[i-n] == 1 {
				dp[i] = 1
				break
			}
		}
	}
	return dp[m] == 1
}

// @lc code=end
