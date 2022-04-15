package main

/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */

// @lc code=start

/*
解法1 动态规划
*/
func longestValidParentheses(s string) int {
	/*
	   "(("
	   ")("
	   "()()"
	   "(())"
	*/
	m := len(s)
	if m < 2 {
		return 0
	}

	//dp[i]表示以s[i]结尾时最长有效括号子串的长度
	dp := make([]int, m)

	check := func(i int) int {
		if s[i-1] == '(' && s[i] == '(' {
			return 0
		} else if s[i-1] == ')' && s[i] == '(' {
			return 0
		} else if s[i-1] == '(' && s[i] == ')' {
			return 1
		} else {
			return 2
		}
	}

	ret := 0
	for i := 1; i < m; i++ {
		dp[i] = 0
		if check(i) == 1 {
			dp[i] = 2
			for j := i - 2; j >= 0 && dp[j] > 0; j = j - dp[j] {
				dp[i] += dp[j]
			}
		} else if check(i) == 2 {
			if dp[i-1] > 0 {
				dp[i] = dp[i-1]
				left := i - dp[i-1] - 1
				if left >= 0 && s[left] == '(' {
					dp[i] += 2
				} else {
					dp[i] = 0
					continue
				}
				for j := i - dp[i]; j >= 0 && dp[j] > 0; j = j - dp[j] {
					dp[i] += dp[j]
				}
			}
		}
		if dp[i] > ret {
			ret = dp[i]
		}
	}

	return ret
}

// @lc code=end
