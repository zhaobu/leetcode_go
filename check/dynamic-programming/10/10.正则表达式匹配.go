package main

/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */

// @lc code=start

/*
解法1 dp
*/
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)

	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true

	/*
		1. i,j都为dp的下标,对应到字符串上都应-1
		2. i的取值范围是[0,m],j的取值范围是[1,n]
		3. p[j-1]不可能是*
	*/
	match := func(i, j int) bool {

		if i == 0 {
			return false
		}
		ch1, ch2 := s[i-1], p[j-1]
		return ch1 == ch2 || ch2 == '.'
	}

	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				/*
					如果p[i-1]是*
					1. 如果s[i-1]和p[j-2]匹配,也就是p是字符+*的方式. 则这个时候有2种匹配方式
					匹配0次: 把p的字符+*整个都丢掉, 看是否还匹配. 也就是 dp[i][j]= dp[i][j-2]
					匹配1至多次: 看丢掉s最后一个字符后,是否还能匹配. 也就是 dp[i][j]= dp[i-1][j]

					2. 如果s[i-1]和p[j-2]不匹配,则这个时候只能是p的字符+*整体匹配0次
					dp[i][j]= dp[i][j-2]
				*/
				if match(i, j-1) {
					dp[i][j] = dp[i-1][j] || (j >= 2 && dp[i][j-2])
				} else {
					dp[i][j] = j >= 2 && dp[i][j-2]
				}
			} else {
				/*
					如果p[i-1]不是*
					1. 如果s[i-1]和p[j-1]匹配,则dp[i][j] = dp[i-1][j-1]
					2. 如果s[i-1]和p[j-1]不匹配,则dp[i][j] = false
				*/
				dp[i][j] = match(i, j) && dp[i-1][j-1]
			}
		}
	}

	return dp[m][n]
}

/*
解法2 带记录的dfs 参考 https://mp.weixin.qq.com/s/rnaFK05IcFWvNN1ppNf2ug
*/
// @lc code=end
