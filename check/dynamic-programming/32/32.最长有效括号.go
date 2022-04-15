package main

/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */

// @lc code=start

/*
解法1 动态规划
dp[i]表示以s[i]结尾时最长有效括号子串的长度
每个字符共有2种情况,每个字符和前一个字符共有4种情况
1: ((
2: )(
以(结尾的2种情况不用考虑,dp[i]=0
3: ()
直接dp[i]=dp[i-2]+2
4: ))
直接看dp[i-1]
如果dp[i-1]=0,则dp[i]=0
如果dp[i-1]>0,则要看s[i-dp[i-1]-1]是否为(
	如果不是说明不能和)配对,dp[i]=0
	如果是,则已经能得到的最长有效括号子串的长度为dp[i-1]+2,
	另外还要加上(之前的部分,还可能会有有效的子串,长度为dp[i-2-dp[i-1]]
	所以综合来看,就是dp[i]=dp[i-1]+2+dp[i-2-dp[i-1]]
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

	ret := 0
	for i := 1; i < m; i++ {
		dp[i] = 0
		if s[i] == ')' {
			if s[i-1] == '(' {
				dp[i] = 2     //至少已经有了一对()
				if i-2 >= 0 { //加上i-2前面的部分
					dp[i] += dp[i-2]
				}
			} else { //s[i-1] == ')'
				if dp[i-1] > 0 { //s[i-1]能配对才考虑s[i]
					left := i - dp[i-1] - 1
					if left >= 0 && s[left] == '(' {
						dp[i] = dp[i-1] + 2
						if left-1 >= 0 {
							dp[i] += dp[left-1]
						}
					}
				}
			}
			if dp[i] > ret {
				ret = dp[i]
			}
		}
	}

	return ret
}

// @lc code=end
